package logger

import (
	"OverheadTGBot/pkg"
	config "OverheadTGBot/pkg/config/entity"
	"go.uber.org/zap"
	"log"
)

var serviceLogger zapLogger

type zapLogger struct {
	config config.LoggerConfig
	logger *zap.SugaredLogger
}

func NewZapLogger(config config.LoggerConfig) Logger {
	var logger *zap.Logger
	var err error
	if pkg.IsDev() {
		logger, err = zap.NewDevelopment()
	}
	if pkg.IsProd() {
		logger, err = zap.NewProduction()
	}
	if !pkg.IsProd() && !pkg.IsDev() {
		logger = zap.NewExample()
	}
	if err != nil {
		log.Fatal(err)
	}
	/*defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(logger) // flushes buffer, if any*/
	serviceLogger := &zapLogger{logger: logger.Sugar(), config: config}

	return serviceLogger
}

func Get() Logger {
	return serviceLogger
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
