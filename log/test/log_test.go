package test

import (
	"github.com/typeck/frame/log"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestSugaredLog(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "www.baidu.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "www.baidu.com")
}

func TestLog(t *testing.T) {
	log.Infow("tttt", log.Field{Key: "key", Type: zapcore.Int64Type, Integer: 65})
	log.Error("ggjggk")
	log.Debug("ffjfj")
	log.Warn("fff")
}