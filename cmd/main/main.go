package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	"github.com/instill-ai/controller-model/config"
	"github.com/instill-ai/controller-model/internal/external"
	"github.com/instill-ai/controller-model/pkg/handler"
	"github.com/instill-ai/controller-model/pkg/logger"
	"github.com/instill-ai/controller-model/pkg/service"

	custom_otel "github.com/instill-ai/controller-model/pkg/logger/otel"
	controllerPB "github.com/instill-ai/protogen-go/model/controller/v1alpha"
)

var propagator propagation.TextMapPropagator

func grpcHandlerFunc(grpcServer *grpc.Server, gwHandler http.Handler) http.Handler {
	return h2c.NewHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			propagator = b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader))
			ctx := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
			r = r.WithContext(ctx)

			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
			} else {
				gwHandler.ServeHTTP(w, r)
			}
		}),
		&http2.Server{},
	)
}

func main() {
	if err := config.Init(); err != nil {
		log.Fatal(err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())

	if tp, err := custom_otel.SetupTracing(ctx, "controller"); err != nil {
		panic(err)
	} else {
		defer func() {
			err = tp.Shutdown(ctx)
		}()
	}

	if mp, err := custom_otel.SetupMetrics(ctx, "controller"); err != nil {
		panic(err)
	} else {
		defer func() {
			err = mp.Shutdown(ctx)
		}()
	}

	ctx, span := otel.Tracer("main-tracer").Start(ctx,
		"main",
	)
	defer cancel()

	logger, _ := logger.GetZapLogger(ctx)
	defer func() {
		// can't handle the error due to https://github.com/uber-go/zap/issues/880
		_ = logger.Sync()
	}()

	// verbosity 3 will avoid [transport] from emitting
	grpc_zap.ReplaceGrpcLoggerV2WithVerbosity(logger, 3)

	// Create tls based credential.
	var creds credentials.TransportCredentials
	var err error
	if config.Config.Server.HTTPS.Cert != "" && config.Config.Server.HTTPS.Key != "" {
		creds, err = credentials.NewServerTLSFromFile(config.Config.Server.HTTPS.Cert, config.Config.Server.HTTPS.Key)
		if err != nil {
			logger.Fatal(fmt.Sprintf("failed to create credentials: %v", err))
		}
	}

	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_zap.Option{
		grpc_zap.WithDecider(func(fullMethodName string, err error) bool {
			// will not log gRPC calls if it was a call to liveness or readiness and no error was raised
			if err == nil {
				// stop logging successful private function calls
				if match, _ := regexp.MatchString("model.controller.v1alpha.ControllerPrivateService/.*", fullMethodName); match {
					return false
				}
			}
			// by default everything will be logged
			return true
		}),
	}

	grpcServerOpts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(logger, opts...),
			grpc_recovery.StreamServerInterceptor(recoveryInterceptorOpt()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger, opts...),
			grpc_recovery.UnaryServerInterceptor(recoveryInterceptorOpt()),
		)),
	}

	if config.Config.Server.HTTPS.Cert != "" && config.Config.Server.HTTPS.Key != "" {
		grpcServerOpts = append(grpcServerOpts, grpc.Creds(creds))
	}

	grpcS := grpc.NewServer(grpcServerOpts...)
	reflection.Register(grpcS)

	modelPublicServiceClient, modelPublicServiceClientConn := external.InitModelPublicServiceClient(ctx)
	defer modelPublicServiceClientConn.Close()

	modelPrivateServiceClient, modelPrivateServiceClientConn := external.InitModelPrivateServiceClient(ctx)
	defer modelPrivateServiceClientConn.Close()

	mgmtPublicServiceClient, mgmtPublicServiceClientConn := external.InitMgmtPublicServiceClient(ctx)
	defer mgmtPublicServiceClientConn.Close()

	etcdClient := external.InitEtcdServiceClient(ctx)
	defer etcdClient.Close()

	redisClient := redis.NewClient(&config.Config.Cache.Redis.RedisOptions)
	defer redisClient.Close()

	service := service.NewService(
		modelPublicServiceClient,
		modelPrivateServiceClient,
		mgmtPublicServiceClient,
		*etcdClient,
		redisClient,
	)

	controllerPB.RegisterControllerPrivateServiceServer(
		grpcS, handler.NewPrivateHandler(
			service,
		),
	)

	serverMux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(httpResponseModifier),
		runtime.WithErrorHandler(errorHandler),
		runtime.WithIncomingHeaderMatcher(customMatcher),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
				UseEnumNumbers:  false,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)

	var dialOpts []grpc.DialOption
	if config.Config.Server.HTTPS.Cert != "" && config.Config.Server.HTTPS.Key != "" {
		dialOpts = []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	} else {
		dialOpts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}

	if err := controllerPB.RegisterControllerPrivateServiceHandlerFromEndpoint(ctx, serverMux, fmt.Sprintf(":%v", config.Config.Server.Port), dialOpts); err != nil {
		logger.Fatal(err.Error())
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Config.Server.Port),
		Handler: grpcHandlerFunc(grpcS, serverMux),
	}

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quitSig := make(chan os.Signal, 1)
	errSig := make(chan error)
	if config.Config.Server.HTTPS.Cert != "" && config.Config.Server.HTTPS.Key != "" {
		go func() {
			if err := httpServer.ListenAndServeTLS(config.Config.Server.HTTPS.Cert, config.Config.Server.HTTPS.Key); err != nil {
				errSig <- err
			}
		}()
	} else {
		go func() {
			if err := httpServer.ListenAndServe(); err != nil {
				errSig <- err
			}
		}()
	}
	span.End()
	logger.Info("gRPC server is running.")

	go func() {

		logger.Info("[controller] control loop started")
		var mainWG sync.WaitGroup
		for {
			for etcdClient.ActiveConnection().GetState() != connectivity.Ready {
				logger.Warn("[controller] etcd connection lost, waiting for state change...")
				etcdClient.ActiveConnection().WaitForStateChange(ctx, connectivity.TransientFailure)
				time.Sleep(50 * time.Millisecond)
			}

			// Backend services
			mainWG.Add(1)
			go func() {
				defer mainWG.Done()
				if err := service.ProbeBackend(context.WithTimeout(ctx, config.Config.Server.Timeout*time.Second)); err != nil {
					logger.Error(err.Error())
				}
			}()

			// Models
			mainWG.Add(1)
			go func() {
				defer mainWG.Done()
				if err := service.ProbeModels(context.WithTimeout(ctx, config.Config.Server.Timeout*time.Second)); err != nil {
					logger.Error(err.Error())
				}
			}()
			mainWG.Add(1)
			go func() {
				defer mainWG.Done()
				if err := service.MonitorModelCache(context.WithTimeout(ctx, config.Config.Server.Timeout*time.Second)); err != nil {
					logger.Error(err.Error())
				}
			}()

			time.Sleep(config.Config.Server.LoopInterval * time.Second)
			mainWG.Wait()
		}
	}()

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quitSig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errSig:
		logger.Error(fmt.Sprintf("Fatal error: %v\n", err))
	case <-quitSig:
		logger.Info("Shutting down server...")
		grpcS.GracefulStop()
	}

}
