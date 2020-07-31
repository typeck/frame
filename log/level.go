package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levelMap = map[string]zap.LevelEnablerFunc{
	"debug":debugLevel,
	"info":infoLevel,
	"warn":warnLevel,
	"error":errorLevel,
	"dpanic":dpanicLevel,
	"panic":panicLevel,
	"fatal":fatalLevel,
}
var (
	debugLevel =  zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})
	infoLevel =  zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	warnLevel =  zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})
	errorLevel =  zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	})
	dpanicLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DPanicLevel
	})
	panicLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.PanicLevel
	})
	fatalLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.FatalLevel
	})
)
