package logger

import (
	"sync"
)

var (
	defaultLogger Logger
	singleton     sync.Once
)

//Logger provide interface for logging
type Logger interface {
	//Debug log msg if log level set to DEBUG
	Debug(args ...interface{})
	//Debug log msg with format specifier if log level set to DEBUG
	Debugf(format string, args ...interface{})
	//Info log msg with INFO level
	Info(args ...interface{})
	//Infof log msg with format specifier if log level set to INFO
	Infof(format string, args ...interface{})
	//Warn log msg with WARN level
	Warn(args ...interface{})
	//Warnf log msg with format specifier if log level set to WARN level
	Warnf(format string, args ...interface{})
	//Error log msg if log level set to ERROR
	Error(args ...interface{})
	//Errorf log msg with format specifier if log level set to ERROR
	Errorf(format string, args ...interface{})
	//Fatal log msg if log level set to FATAL
	Fatal(args ...interface{})
	//Fatal log msg with format specifier if log level set to FATAL
	Fatalf(format string, args ...interface{})
}

//InitLogger initializes the Logger interface with singleton pattern
func InitLogger(logger Logger) {
	singleton.Do(func() {
		defaultLogger = logger
	})
}

// Debug log msg if log level is set to DEBUG
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Debugf log msg if log level is set to DEBUG with format specified
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Info log msg if log level is set to INFO
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Infof log msg if log level is set to INFO with format specified
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warn log msg if log level is set to WARN
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Warnf log msg if log level is set to WARN with format specified
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Error log msg if log level is set to ERROR
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Errorf log msg if log level is set to ERROR with format specified
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Fatal log msg if log level is set to FATAL
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Fatalf log msg if log level is set to FATAL with format specified
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}
