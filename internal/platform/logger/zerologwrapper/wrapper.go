package zerologwrapper

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type zeroLogWrapper struct {
	logger *zerolog.Logger
}

//NewZeroLogWrapper returns ZeroLogWrapper instance
func NewZeroLogWrapper(logger zerolog.Logger) *zeroLogWrapper {
	return &zeroLogWrapper{logger: &logger}
}

// Debug log msg if log level is set to DEBUG
func (z *zeroLogWrapper) Debug(args ...interface{}) {
	//TODO: add multi type values in debug
	z.logger.Debug().Msg(fmt.Sprint(args...))
}

// Debugf log msg if log level is set to DEBUG with format specified
func (z *zeroLogWrapper) Debugf(format string, args ...interface{}) {
	//TODO: add multi type values in debugf
	z.logger.Debug().Msgf(format, args...)
}

// Info log msg if log level is set to INFO
func (z *zeroLogWrapper) Info(args ...interface{}) {
	//TODO: add multi type values
	z.logger.Info().Msg(fmt.Sprint(args...))
}

// Infof log msg if log level is set to INFO with format specified
func (z *zeroLogWrapper) Infof(format string, args ...interface{}) {
	//TODO: add multi type values
	z.logger.Info().Msgf(format, args...)
}

// Warn log msg if log level is set to WARN
func (z *zeroLogWrapper) Warn(args ...interface{}) {
	//TODO: add multi type values
	z.logger.Warn().Msg(fmt.Sprint(args...))
}

// Warnf log msg if log level is set to WARN with format specified
func (z *zeroLogWrapper) Warnf(format string, args ...interface{}) {
	//TODO: add multi type values
	z.logger.Warn().Msgf(format, args...)
}

// Error log msg if log level is set to ERROR
func (z *zeroLogWrapper) Error(args ...interface{}) {
	//TODO: add multi type values
	z.logger.Error().Msg(fmt.Sprint(args...))
}

// Errorf log msg if log level is set to ERROR with format specified
func (z *zeroLogWrapper) Errorf(format string, args ...interface{}) {
	//TODO: add multi type values
	z.logger.Error().Msgf(format, args...)
}

// Fatal log msg if log level is set to FATAL
func (z *zeroLogWrapper) Fatal(args ...interface{}) {
	//TODO: add multi type values
	z.logger.Fatal().Msg(fmt.Sprint(args...))
}

// Fatalf log msg if log level is set to FATAL with format specified
func (z *zeroLogWrapper) Fatalf(format string, args ...interface{}) {
	//TODO: add multi type values
	z.logger.Fatal().Msgf(format, args...)
}

// loglevel maps the string level to zerolog.Level
func logLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.DebugLevel
	}
}

// GetZerologDefaultLogger returns zerolog.ConsoleWriter
// writer: used to write the logs
// level: loglevel
// skipCallFrameCount: use to correctly print caller's lineno. and filename
func GetZerologDefaultLogger(writer io.Writer, level string, skipCallerFrameCount int) zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: writer, TimeFormat: time.RFC3339}
	return zerolog.New(output).With().CallerWithSkipFrameCount(skipCallerFrameCount).Timestamp().Logger().Level(logLevel(level))
}
