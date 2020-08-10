package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Logger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

type Config struct {
	Path 		string		`toml:"path"`
	FileName 	string		`toml:"file_name"`
	MaxSize 	int			`toml:"max_size"`
	MaxBackups 	int			`toml:"max_backups"`
	MaxAge 		int			`toml:"max_age"`
	Compress	bool		`toml:"compress"`
	//json or console
	Encoding 	string		`toml:"encoding"`
}

type Field = zap.Field

var DefaultLogger *Logger

func DefaultConfig() *Config {
	return &Config{
		MaxSize:    500,
		MaxBackups: 10,
		MaxAge:     1,
		Compress:   false,
		Encoding: "json",
	}
}

func New(conf *Config, opts... zap.Option) *Logger {
	core := buildCore(conf)
	logger := zap.New(core, opts...)
	return &Logger{
		logger: logger,
		sugar: logger.Sugar(),
	}
}

func getEncoder(conf *Config) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	if conf.Encoding == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(conf *Config, levelName string) zapcore.WriteSyncer {
	var name string
	index := strings.Index(conf.FileName, ".")
	if index <= 0 {
		if conf.FileName != "" {
			name = conf.FileName + "_"
		}
		name = name + levelName + ".log"
	}else {
		name = conf.FileName[:index] + "_" + levelName + conf.FileName[index:]
	}
	name = conf.Path + name
	lumberJackLogger := &lumberjack.Logger{
		Filename:   name,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func buildCore(conf *Config) zapcore.Core{
	core := zapcore.NewTee(
		zapcore.NewCore(getEncoder(conf), getLogWriter(conf, "debug"), DebugLevel),
		zapcore.NewCore(getEncoder(conf), getLogWriter(conf, "info"), InfoLevel),
		zapcore.NewCore(getEncoder(conf), getLogWriter(conf, "warn"), WarnLevel),
		zapcore.NewCore(getEncoder(conf), getLogWriter(conf, "error"), ErrorLevel),
		zapcore.NewCore(getEncoder(conf), getLogWriter(conf, "dpanic"), DpanicLevel),
		zapcore.NewCore(getEncoder(conf), getLogWriter(conf, "panic"), PanicLevel),
		zapcore.NewCore(getEncoder(conf), getLogWriter(conf, "fatal"), FatalLevel),
	)
	return core
}

func (logger *Logger) WithOptions(opts... zap.Option) {
	logger.logger = logger.logger.WithOptions(opts...)
	logger.sugar = logger.logger.Sugar()
}

// Info ...
func (logger *Logger) Info(args ...interface{}) {
	logger.sugar.Info(args...)
}

func (logger *Logger) Infof(template string, args ...interface{}) {
	logger.sugar.Infof(template, args...)
}

func (logger *Logger) Infoh(msg string, fields ...Field) {
	logger.logger.Info(msg, fields...)
}

func (logger *Logger) Infow(msg string, keysAndValues ...interface{}) {
	logger.sugar.Infow(msg, keysAndValues...)
}

//Warn ...
func (logger *Logger) Warn(args ...interface{}) {
	logger.sugar.Warn(args...)
}

func (logger *Logger) Warnf(template string, args ...interface{}) {
	logger.sugar.Warnf(template, args...)
}

func (logger *Logger) Warnh(msg string, fields ...Field) {
	logger.logger.Warn(msg, fields...)
}

func (logger *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	logger.sugar.Warnw(msg, keysAndValues...)
}

//Error ...
func (logger *Logger) Error(args ...interface{}) {
	logger.sugar.Error(args...)
}

func (logger *Logger) Errorf(template string, args ...interface{}) {
	logger.sugar.Errorf(template, args...)
}

func (logger *Logger) Errorh(msg string, fields ...Field) {
	logger.logger.Error(msg, fields...)
}

func (logger *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	logger.sugar.Errorw(msg, keysAndValues...)
}

//Debug ...
func (logger *Logger) Debug(args ...interface{}) {
	logger.sugar.Debug(args...)
}

func (logger *Logger) Debugf(template string, args ...interface{}) {
	logger.sugar.Debugf(template, args...)
}

func (logger *Logger) Debugh(msg string, fields ...Field) {
	logger.logger.Debug(msg, fields...)
}

func (logger *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	logger.sugar.Debugw(msg, keysAndValues...)
}


//Fatal ...
func (logger *Logger) Fatal(args ...interface{}) {
	logger.sugar.Fatal(args...)
}

func (logger *Logger) Fatalf(template string, args ...interface{}) {
	logger.sugar.Fatalf(template, args...)
}

func (logger *Logger) Fatalh(msg string, fields ...Field) {
	logger.logger.Fatal(msg, fields...)
}

func (logger *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	logger.sugar.Fatalw(msg, keysAndValues...)
}

//Panic
func (logger *Logger) Panic(args ...interface{}) {
	logger.sugar.Panic(args...)
}

func (logger *Logger) Panicf(template string, args ...interface{}) {
	logger.sugar.Panicf(template, args...)
}

func (logger *Logger) Panich(msg string, fields ...Field) {
	logger.logger.Panic(msg, fields...)
}

func (logger *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	logger.sugar.Panicw(msg, keysAndValues...)
}

//Dpanic
func (logger *Logger) DPanic(args ...interface{}) {
	logger.sugar.DPanic(args...)
}

func (logger *Logger) DPanicf(template string, args ...interface{}) {
	logger.sugar.DPanicf(template, args...)
}

func (logger *Logger) DPanich(msg string, fields ...Field) {
	logger.logger.DPanic(msg, fields...)
}

func (logger *Logger) DPanicw(msg string, keysAndValues ...interface{}) {
	logger.sugar.DPanicw(msg, keysAndValues...)
}