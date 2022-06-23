package logger

import (
	"OverheadTGBot/pkg"
	"OverheadTGBot/pkg/config/model"
	"go.uber.org/zap"
	"log"
)

type zapLogger struct {
	config model.LoggerConfig
	logger *zap.SugaredLogger
}

func NewZapLogger(config model.LoggerConfig) Logger {
	var logger *zap.Logger
	var err error
	if pkg.EnvironmentIsDev() {
		logger, err = zap.NewDevelopment()
	}
	if pkg.EnvironmentIsProd() {
		logger, err = zap.NewProduction()
	}
	if !pkg.EnvironmentIsProd() && !pkg.EnvironmentIsDev() {
		logger = zap.NewExample()
	}
	if err != nil {
		log.Fatal(err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(logger) // flushes buffer, if any
	return &zapLogger{logger: logger.Sugar(), config: config}
}

func (z zapLogger) Debug(msg ...interface{}) {
	z.logger.Debug(msg)
}

func (z zapLogger) Debugf(format string, args ...interface{}) {
	z.logger.Debugf(format, args...)
}

func (z zapLogger) Info(msg ...interface{}) {
	z.logger.Info(msg...)
}

func (z zapLogger) Infof(format string, args ...interface{}) {
	z.logger.Infof(format, args...)
}

func (z zapLogger) Warn(msg ...interface{}) {
	z.logger.Warn(msg...)
}

func (z zapLogger) Warnf(format string, args ...interface{}) {
	z.logger.Warnf(format, args...)
}

func (z zapLogger) Error(msg ...interface{}) {
	z.logger.Error(msg...)
}

func (z zapLogger) Errorf(format string, args ...interface{}) {
	z.logger.Errorf(format, args...)
}