package utils

import (
	"log"
	"sync"

	"go.uber.org/zap"
)

var logger *zap.Logger

var Log = getLogger()
var doOnce sync.Once

func getLogger() *zap.Logger {
	doOnce.Do(func() {
		l, err := zap.NewProduction()
		if err != nil {
			log.Panic("error getting logger:", zap.Error(err))
		}

		l.Info("Initialized logger..")
		logger = l
	})
	return logger
}
