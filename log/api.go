package log

// Info ...
func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}

// Debug ...
func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}

// Warn ...
func Warn(args ...interface{}) {
	DefaultLogger.Warn(args...)
}

// Error ...
func Error(args ...interface{}) {
	DefaultLogger.Error(args...)
}

// Panic ...
func Panic(args ...interface{}) {
	DefaultLogger.Panic(args...)
}

// DPanic ...
func DPanic(args ...interface{}) {
	DefaultLogger.DPanic(args...)
}

// Fatal ...
func Fatal(args ...interface{}) {
	DefaultLogger.Fatal(args...)
}

// Debugw ...
func Debugh(msg string, fields ...Field) {
	DefaultLogger.Debugh(msg, fields...)
}

// Infow ...
func Infoh(msg string, fields ...Field) {
	DefaultLogger.Infoh(msg, fields...)
}

// Warnw ...
func Warnh(msg string, fields ...Field) {
	DefaultLogger.Warnh(msg, fields...)
}

// Errorw ...
func Errorh(msg string, fields ...Field) {
	DefaultLogger.Errorh(msg, fields...)
}

// Panicw ...
func Panich(msg string, fields ...Field) {
	DefaultLogger.Panich(msg, fields...)
}

// DPanicw ...
func DPanich(msg string, fields ...Field) {
	DefaultLogger.DPanich(msg, fields...)
}

// Fatalw ...
func Fatalh(msg string, fields ...Field) {
	DefaultLogger.Fatalh(msg, fields...)
}

// Debugf ...
func Debugf(msg string, args ...interface{}) {
	DefaultLogger.Debugf(msg, args...)
}

// Infof ...
func Infof(msg string, args ...interface{}) {
	DefaultLogger.Infof(msg, args...)
}

// Warnf ...
func Warnf(msg string, args ...interface{}) {
	DefaultLogger.Warnf(msg, args...)
}

// Errorf ...
func Errorf(msg string, args ...interface{}) {
	DefaultLogger.Errorf(msg, args...)
}

// Panicf ...
func Panicf(msg string, args ...interface{}) {
	DefaultLogger.Panicf(msg, args...)
}

// DPanicf ...
func DPanicf(msg string, args ...interface{}) {
	DefaultLogger.DPanicf(msg, args...)
}

// Fatalf ...
func Fatalf(msg string, args ...interface{}) {
	DefaultLogger.Fatalf(msg, args...)
}

// Debugw ...
func Debugw(msg string, args ...interface{}) {
	DefaultLogger.Debugw(msg, args...)
}

// Infow ...
func Infow(msg string, args ...interface{}) {
	DefaultLogger.Infow(msg, args...)
}

// Warnf ...
func Warnw(msg string, args ...interface{}) {
	DefaultLogger.Warnw(msg, args...)
}

// Errorf ...
func Errorw(msg string, args ...interface{}) {
	DefaultLogger.Errorw(msg, args...)
}

// Panicf ...
func Panicw(msg string, args ...interface{}) {
	DefaultLogger.Panicw(msg, args...)
}

// DPanicf ...
func DPanicw(msg string, args ...interface{}) {
	DefaultLogger.DPanicw(msg, args...)
}

// Fatalf ...
func Fatalw(msg string, args ...interface{}) {
	DefaultLogger.Fatalw(msg, args...)
}

