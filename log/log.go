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
}

type Field = zap.Field

var DefaultLogger = New(DefaultConfig())

func DefaultConfig() *Config {
	return &Config{
		MaxSize:    500,
		MaxBackups: 10,
		MaxAge:     1,
		Compress:   false,
	}
}

func New(conf *Config) *Logger{
	core := zapcore.NewTee(
		zapcore.NewCore(getEncoder(), getLogWriter(conf, "debug"), DebugLevel),
		zapcore.NewCore(getEncoder(), getLogWriter(conf, "info"), InfoLevel),
		zapcore.NewCore(getEncoder(), getLogWriter(conf, "warn"), WarnLevel),
		zapcore.NewCore(getEncoder(), getLogWriter(conf, "error"), ErrorLevel),
		zapcore.NewCore(getEncoder(), getLogWriter(conf, "dpanic"), DpanicLevel),
		zapcore.NewCore(getEncoder(), getLogWriter(conf, "panic"), PanicLevel),
		zapcore.NewCore(getEncoder(), getLogWriter(conf, "fatal"), FatalLevel),
		)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(PanicLevel))
	return &Logger{
		logger: logger,
		sugar: logger.Sugar(),
	}
}

func getEncoder() zapcore.Encoder {
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
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(conf *Config, levelName string) zapcore.WriteSyncer {
	var name = conf.Path
	index := strings.Index(conf.FileName, ".")
	if index <= 0 {
		name = conf.FileName + "_" + levelName + ".log"
	}else {
		name = conf.FileName[:index] + "_" + levelName + conf.FileName[index:]
	}
	if conf.FileName == "" {
		name = levelName + ".log"
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   name,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}



// Info ...
func (logger *Logger) Info(args ...interface{}) {
	logger.sugar.Info(args...)
}

func (logger *Logger) Infof(template string, args ...interface{}) {
	logger.sugar.Infof(template, args...)
}

func (logger *Logger) Infow(msg string, fields ...Field) {
	logger.logger.Info(msg, fields...)
}

//Warn ...
func (logger *Logger) Warn(args ...interface{}) {
	logger.sugar.Warn(args...)
}

func (logger *Logger) Warnf(template string, args ...interface{}) {
	logger.sugar.Warnf(template, args...)
}

func (logger *Logger) Warnw(msg string, fields ...Field) {
	logger.logger.Warn(msg, fields...)
}

//Error ...
func (logger *Logger) Error(args ...interface{}) {
	logger.sugar.Error(args...)
}

func (logger *Logger) Errorf(template string, args ...interface{}) {
	logger.sugar.Errorf(template, args...)
}

func (logger *Logger) Errorw(msg string, fields ...Field) {
	logger.logger.Error(msg, fields...)
}

//Debug ...
func (logger *Logger) Debug(args ...interface{}) {
	logger.sugar.Debug(args...)
}

func (logger *Logger) Debugf(template string, args ...interface{}) {
	logger.sugar.Debugf(template, args...)
}

func (logger *Logger) Debugw(msg string, fields ...Field) {
	logger.logger.Debug(msg, fields...)
}

//Fatal ...
func (logger *Logger) Fatal(args ...interface{}) {
	logger.sugar.Fatal(args...)
}

func (logger *Logger) Fatalf(template string, args ...interface{}) {
	logger.sugar.Fatalf(template, args...)
}

func (logger *Logger) Fatalw(msg string, fields ...Field) {
	logger.logger.Fatal(msg, fields...)
}

//Panic
func (logger *Logger) Panic(args ...interface{}) {
	logger.sugar.Panic(args...)
}

func (logger *Logger) Panicf(template string, args ...interface{}) {
	logger.sugar.Panicf(template, args...)
}

func (logger *Logger) Panicw(msg string, fields ...Field) {
	logger.logger.Panic(msg, fields...)
}

//Dpanic
func (logger *Logger) DPanic(args ...interface{}) {
	logger.sugar.DPanic(args...)
}

func (logger *Logger) DPanicf(template string, args ...interface{}) {
	logger.sugar.DPanicf(template, args...)
}

func (logger *Logger) DPanicw(msg string, fields ...Field) {
	logger.logger.DPanic(msg, fields...)
}