package service

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/instill-ai/controller-model/config"
	"github.com/instill-ai/controller-model/pkg/logger"
)

func (s *service) MonitorModelCache(ctx context.Context, cancel context.CancelFunc) error {
	defer cancel()

	logger, _ := logger.GetZapLogger(ctx)

	var wg sync.WaitGroup

	iter := s.redisClient.Scan(ctx, 0, "instill-ai/*", 0).Iterator()

	for iter.Next(ctx) {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()

			idleTime, err := s.redisClient.ObjectIdleTime(ctx, key).Result()
			if err != nil {
				logger.Error(err.Error())
				return
			}

			retentionPeriod, err := time.ParseDuration(config.Config.Cache.Model.RetentionPeriod)
			if err != nil {
				logger.Error(err.Error())
				return
			}

			pathSplit := strings.Split(key, ":")
			modelPath := config.Config.Cache.Model.CacheDir + "/" + fmt.Sprintf("%s_%s", pathSplit[0], pathSplit[1])

			if idleTime >= retentionPeriod {
				err := os.RemoveAll(modelPath)
				if err != nil {
					logger.Error(err.Error())
					return
				}
				_, err = s.redisClient.Del(ctx, key).Result()
				if err != nil {
					logger.Error(err.Error())
					return
				}
			}

		}(iter.Val())
	}

	wg.Wait()

	return nil
}
