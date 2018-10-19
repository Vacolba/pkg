package log

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logLevels   = []string{"debug", "info", "warn", "error"}
	logEncoders = []string{"json", "text"}
)

// NewZapLogger creates and returns a zap Logger
func NewZapLogger(level, format string) (*zap.Logger, error) {
	var logLevel zapcore.Level
	var encoder zapcore.Encoder
	var encoderCfg zapcore.EncoderConfig

	if strings.ToLower(level) == "debug" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	switch strings.ToLower(level) {
	case "debug":
		logLevel = zap.DebugLevel
	case "", "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	default:
		return nil, fmt.Errorf("log level is not one of the supported values (%s): %s", strings.Join(logLevels, ", "), level)
	}

	switch strings.ToLower(format) {
	case "", "text":
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	default:
		return nil, fmt.Errorf("log format is not one of the supported values (%s): %s", strings.Join(logEncoders, ", "), format)
	}

	atom := zap.NewAtomicLevel()

	logger := zap.New(zapcore.NewCore(
		encoder,
		zapcore.Lock(os.Stdout),
		atom,
	))
	defer logger.Sync() // flushes buffer, if any

	atom.SetLevel(logLevel)
	return logger, nil
}

// NewZapSugaredLogger creates and returns a zap SugaredLogger
func NewZapSugaredLogger(level, format string) (*zap.SugaredLogger, error) {
	logger, err := NewZapLogger(level, format)
	return logger.Sugar(), err
}

// NewGooseLogger creates and returs a logger for Goose
func NewGooseLogger(level, format string) (*GooseLogger, error) {
	zLogger, err := NewZapLogger(level, format)
	if err != nil {
		return nil, err
	}
	logger := &GooseLogger{
		log:    zLogger.Sugar(),
		fatal:  (*zap.SugaredLogger).Fatal,
		fatalf: (*zap.SugaredLogger).Fatalf,
		print:  (*zap.SugaredLogger).Info,
		printf: (*zap.SugaredLogger).Infof,
	}
	return logger, nil
}

// GooseLogger adapts zap's GooseLogger to be compatible with goose.Logger.
type GooseLogger struct {
	log    *zap.SugaredLogger
	fatal  func(*zap.SugaredLogger, ...interface{})
	fatalf func(*zap.SugaredLogger, string, ...interface{})
	print  func(*zap.SugaredLogger, ...interface{})
	printf func(*zap.SugaredLogger, string, ...interface{})
}

// Fatal implements goose.Logger.
func (l *GooseLogger) Fatal(args ...interface{}) {
	l.fatal(l.log, args...)
}

// Fatalf implements goose.Logger.
func (l *GooseLogger) Fatalf(format string, args ...interface{}) {
	l.fatalf(l.log, format, args...)
}

// Fatalln implements goose.Logger.
func (l *GooseLogger) Fatalln(args ...interface{}) {
	l.fatal(l.log, args...)
}

// Print implements goose.Logger.
func (l *GooseLogger) Print(args ...interface{}) {
	l.print(l.log, args...)
}

// Printf implements goose.Logger.
func (l *GooseLogger) Printf(format string, args ...interface{}) {
	l.printf(l.log, format, args...)
}

// Println implements goose.Logger.
func (l *GooseLogger) Println(args ...interface{}) {
	l.print(l.log, args...)
}
