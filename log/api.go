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
func Debugw(msg string, fields ...Field) {
	DefaultLogger.Debugw(msg, fields...)
}

// Infow ...
func Infow(msg string, fields ...Field) {
	DefaultLogger.Infow(msg, fields...)
}

// Warnw ...
func Warnw(msg string, fields ...Field) {
	DefaultLogger.Warnw(msg, fields...)
}

// Errorw ...
func Errorw(msg string, fields ...Field) {
	DefaultLogger.Errorw(msg, fields...)
}

// Panicw ...
func Panicw(msg string, fields ...Field) {
	DefaultLogger.Panicw(msg, fields...)
}

// DPanicw ...
func DPanicw(msg string, fields ...Field) {
	DefaultLogger.DPanicw(msg, fields...)
}

// Fatalw ...
func Fatalw(msg string, fields ...Field) {
	DefaultLogger.Fatalw(msg, fields...)
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

