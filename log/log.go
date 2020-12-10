package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Logger struct {
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
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

func SetDefaultLogger(logger *Logger) {
	DefaultLogger = logger
}

func New(conf *Config, opts... zap.Option) *Logger {
	core := buildCore(conf)
	logger := zap.New(core, opts...)
	return &Logger{
		Logger: logger,
		Sugar: logger.Sugar(),
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
	logger.Logger = logger.Logger.WithOptions(opts...)
	logger.Sugar = logger.Logger.Sugar()
}

// Info ...
func (logger *Logger) Info(args ...interface{}) {
	logger.Sugar.Info(args...)
}

func (logger *Logger) Infof(template string, args ...interface{}) {
	logger.Sugar.Infof(template, args...)
}

func (logger *Logger) Infoh(msg string, fields ...Field) {
	logger.Logger.Info(msg, fields...)
}

func (logger *Logger) Infow(msg string, keysAndValues ...interface{}) {
	logger.Sugar.Infow(msg, keysAndValues...)
}

//Warn ...
func (logger *Logger) Warn(args ...interface{}) {
	logger.Sugar.Warn(args...)
}

func (logger *Logger) Warnf(template string, args ...interface{}) {
	logger.Sugar.Warnf(template, args...)
}

func (logger *Logger) Warnh(msg string, fields ...Field) {
	logger.Logger.Warn(msg, fields...)
}

func (logger *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	logger.Sugar.Warnw(msg, keysAndValues...)
}

//Error ...
func (logger *Logger) Error(args ...interface{}) {
	logger.Sugar.Error(args...)
}

func (logger *Logger) Errorf(template string, args ...interface{}) {
	logger.Sugar.Errorf(template, args...)
}

func (logger *Logger) Errorh(msg string, fields ...Field) {
	logger.Logger.Error(msg, fields...)
}

func (logger *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	logger.Sugar.Errorw(msg, keysAndValues...)
}

//Debug ...
func (logger *Logger) Debug(args ...interface{}) {
	logger.Sugar.Debug(args...)
}

func (logger *Logger) Debugf(template string, args ...interface{}) {
	logger.Sugar.Debugf(template, args...)
}

func (logger *Logger) Debugh(msg string, fields ...Field) {
	logger.Logger.Debug(msg, fields...)
}

func (logger *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	logger.Sugar.Debugw(msg, keysAndValues...)
}


//Fatal ...
func (logger *Logger) Fatal(args ...interface{}) {
	logger.Sugar.Fatal(args...)
}

func (logger *Logger) Fatalf(template string, args ...interface{}) {
	logger.Sugar.Fatalf(template, args...)
}

func (logger *Logger) Fatalh(msg string, fields ...Field) {
	logger.Logger.Fatal(msg, fields...)
}

func (logger *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	logger.Sugar.Fatalw(msg, keysAndValues...)
}

//Panic
func (logger *Logger) Panic(args ...interface{}) {
	logger.Sugar.Panic(args...)
}

func (logger *Logger) Panicf(template string, args ...interface{}) {
	logger.Sugar.Panicf(template, args...)
}

func (logger *Logger) Panich(msg string, fields ...Field) {
	logger.Logger.Panic(msg, fields...)
}

func (logger *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	logger.Sugar.Panicw(msg, keysAndValues...)
}

//Dpanic
func (logger *Logger) DPanic(args ...interface{}) {
	logger.Sugar.DPanic(args...)
}

func (logger *Logger) DPanicf(template string, args ...interface{}) {
	logger.Sugar.DPanicf(template, args...)
}

func (logger *Logger) DPanich(msg string, fields ...Field) {
	logger.Logger.DPanic(msg, fields...)
}

func (logger *Logger) DPanicw(msg string, keysAndValues ...interface{}) {
	logger.Sugar.DPanicw(msg, keysAndValues...)
}