package external

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/instill-ai/controller-model/config"
	"github.com/instill-ai/controller-model/pkg/logger"

	mgmtPB "github.com/instill-ai/protogen-go/core/mgmt/v1beta"
	modelPB "github.com/instill-ai/protogen-go/model/model/v1alpha"
	etcdv3 "go.etcd.io/etcd/client/v3"
)

// InitModelPublicServiceClient initialises a ModelPublicServiceClient instance
func InitModelPublicServiceClient(ctx context.Context) (modelPB.ModelPublicServiceClient, *grpc.ClientConn) {
	logger, _ := logger.GetZapLogger(ctx)

	var clientDialOpts grpc.DialOption
	if config.Config.ModelBackend.HTTPS.Cert != "" && config.Config.ModelBackend.HTTPS.Key != "" {
		creds, err := credentials.NewServerTLSFromFile(config.Config.ModelBackend.HTTPS.Cert, config.Config.ModelBackend.HTTPS.Key)
		if err != nil {
			logger.Fatal(err.Error())
		}
		clientDialOpts = grpc.WithTransportCredentials(creds)
	} else {
		clientDialOpts = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	clientConn, err := grpc.Dial(fmt.Sprintf("%v:%v", config.Config.ModelBackend.Host, config.Config.ModelBackend.PublicPort), clientDialOpts)
	if err != nil {
		logger.Error(err.Error())
		return nil, nil
	}

	return modelPB.NewModelPublicServiceClient(clientConn), clientConn
}

// InitModelPrivateServiceClient initialises a ModelPrivateServiceClient instance
func InitModelPrivateServiceClient(ctx context.Context) (modelPB.ModelPrivateServiceClient, *grpc.ClientConn) {
	logger, _ := logger.GetZapLogger(ctx)

	var clientDialOpts grpc.DialOption
	if config.Config.ModelBackend.HTTPS.Cert != "" && config.Config.ModelBackend.HTTPS.Key != "" {
		creds, err := credentials.NewServerTLSFromFile(config.Config.ModelBackend.HTTPS.Cert, config.Config.ModelBackend.HTTPS.Key)
		if err != nil {
			logger.Fatal(err.Error())
		}
		clientDialOpts = grpc.WithTransportCredentials(creds)
	} else {
		clientDialOpts = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	clientConn, err := grpc.Dial(fmt.Sprintf("%v:%v", config.Config.ModelBackend.Host, config.Config.ModelBackend.PrivatePort), clientDialOpts)
	if err != nil {
		logger.Error(err.Error())
		return nil, nil
	}

	return modelPB.NewModelPrivateServiceClient(clientConn), clientConn
}

// InitMgmtPublicServiceClient initialises a MgmtPublicServiceClient instance
func InitMgmtPublicServiceClient(ctx context.Context) (mgmtPB.MgmtPublicServiceClient, *grpc.ClientConn) {
	logger, _ := logger.GetZapLogger(ctx)

	var clientDialOpts grpc.DialOption
	var creds credentials.TransportCredentials
	var err error
	if config.Config.MgmtBackend.HTTPS.Cert != "" && config.Config.MgmtBackend.HTTPS.Key != "" {
		creds, err = credentials.NewServerTLSFromFile(config.Config.MgmtBackend.HTTPS.Cert, config.Config.MgmtBackend.HTTPS.Key)
		if err != nil {
			logger.Fatal(err.Error())
		}
		clientDialOpts = grpc.WithTransportCredentials(creds)
	} else {
		clientDialOpts = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	clientConn, err := grpc.Dial(fmt.Sprintf("%v:%v", config.Config.MgmtBackend.Host, config.Config.MgmtBackend.PublicPort), clientDialOpts)
	if err != nil {
		logger.Error(err.Error())
		return nil, nil
	}

	return mgmtPB.NewMgmtPublicServiceClient(clientConn), clientConn
}

// InitEtcdServiceClient initialises a EtcdServiceClient instance
func InitEtcdServiceClient(ctx context.Context) *etcdv3.Client {
	logger, _ := logger.GetZapLogger(ctx)

	host := config.Config.Etcd.Host
	port := config.Config.Etcd.Port
	timeout := config.Config.Etcd.Timeout * time.Second
	// Create etcd client
	client, err := etcdv3.New(etcdv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%s", host, port)}, //TODO: multiple nodes
		DialTimeout: timeout,
	})
	if err != nil {
		logger.Fatal(err.Error())
	}

	return client
}
