package logger

import (
	"go.uber.org/zap"
	"songsLibrary/config"
)

type Logger interface {
	InitLogger()
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
}

type ApiLogger struct {
	suggarLogger *zap.SugaredLogger
	cfg          *config.Config
}

func NewApiLogger(cfg *config.Config) *ApiLogger {
	return &ApiLogger{
		cfg: cfg,
	}
}

func (l *ApiLogger) InitLogger() {
	configL := zap.NewProductionConfig()
	configL.OutputPaths = []string{"app.log", "stdout"}

	switch l.cfg.Logger.LogLevel {
	case "debug":
		configL.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		configL.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		configL.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		configL.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		configL.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	logger, err := configL.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	l.suggarLogger = logger.Sugar()
}

func (l *ApiLogger) Info(args ...interface{}) {
	l.suggarLogger.Info(args...)
}

func (l *ApiLogger) Infof(template string, args ...interface{}) {
	l.suggarLogger.Infof(template, args...)
}

func (l *ApiLogger) Error(args ...interface{}) {
	l.suggarLogger.Error(args...)
}

func (l *ApiLogger) Errorf(template string, args ...interface{}) {
	l.suggarLogger.Errorf(template, args...)
}

func (l *ApiLogger) Fatal(args ...interface{}) {
	l.suggarLogger.Fatal(args...)
}

func (l *ApiLogger) Fatalf(template string, args ...interface{}) {
	l.suggarLogger.Fatalf(template, args...)
}

func (l *ApiLogger) Warn(args ...interface{}) {
	l.suggarLogger.Warn(args...)
}
func (l *ApiLogger) Warnf(template string, args ...interface{}) {
	l.suggarLogger.Warnf(template, args...)
}

func (l *ApiLogger) Debug(args ...interface{}) {
	l.suggarLogger.Debug(args...)
}

func (l *ApiLogger) Debugf(template string, args ...interface{}) {
	l.suggarLogger.Debugf(template, args...)
}
