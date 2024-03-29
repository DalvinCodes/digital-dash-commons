package utils

import (
	"go.uber.org/zap"
)

type LoggerInterface interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	With(args ...interface{}) *Logger
	Sync() error
}

type Logger struct {
	*zap.SugaredLogger
}

// Config is the configuration for the logger.
// Refer to https://pkg.go.dev/go.uber.org/zap#Config for more information on the fields.
// If you pass missing fields, the logger will not function properly. All fields are required.
// The default configuration is:
//
//	{ Level: "info", Development: false, Encoding: "json", OutputPaths: ["stdout"], ErrorOutputPaths: ["stderr"] }
type Config struct {
	Level            string
	Development      bool
	Encoding         string
	OutputPaths      []string
	ErrorOutputPaths []string
}

func defaultConfig() Config {
	return Config{
		Level:            "info",
		Development:      false,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func NewLogger(configs ...Config) *Logger {
	var cfg Config

	if len(configs) > 0 {
		cfg = configs[0]
	} else {
		cfg = defaultConfig()
	}

	zapLevel := zap.InfoLevel
	err := zapLevel.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		zapLevel = zap.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	if cfg.Development {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapLevel),
		Development:      cfg.Development,
		Encoding:         cfg.Encoding,
		EncoderConfig:    encoderConfig,
		OutputPaths:      cfg.OutputPaths,
		ErrorOutputPaths: cfg.ErrorOutputPaths,
	}

	baseLogger, err := config.Build()
	if err != nil {
		panic(err)
	}

	sugar := baseLogger.Sugar()

	return &Logger{sugar}
}

func (l *Logger) Debug(args ...interface{}) {
	l.SugaredLogger.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.SugaredLogger.Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.SugaredLogger.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.SugaredLogger.Error(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.SugaredLogger.Panic(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.SugaredLogger.Fatal(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.SugaredLogger.Debugf(template, args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.SugaredLogger.Infof(template, args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.SugaredLogger.Warnf(template, args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.SugaredLogger.Errorf(template, args...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.SugaredLogger.Panicf(template, args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.SugaredLogger.Fatalf(template, args...)
}

func (l *Logger) With(args ...interface{}) *Logger {
	return &Logger{l.SugaredLogger.With(args...)}
}

func (l *Logger) Sync() error {
	return l.SugaredLogger.Sync()
}
