package logger

//go:generate mockgen -destination=mocks/logger.go . Logger

import (
	"go.temporal.io/sdk/log"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Param struct {
	fx.In
}

var Module = fx.Options(
	fx.Provide(New),
)

type Logger interface {
	Info(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
	Warn(msg string, keyvals ...interface{})
	Debug(msg string, keyvals ...interface{})
	With(keyvals ...interface{}) log.Logger
}

func New(param Param) (Logger, error) {
	zapAdapter := ZapAdapter{}
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()
	zapAdapter.logger = logger

	return &zapAdapter, nil
}

type ZapAdapter struct {
	logger *zap.Logger
}

// Log method for Debug Level logging
func (z *ZapAdapter) Debug(msg string, keyvals ...interface{}) {
	z.logger.Sugar().Debugw(msg, keyvals...)

}

// Log method for Warn Level Logging
func (z *ZapAdapter) Warn(msg string, keyvals ...interface{}) {
	z.logger.Sugar().Warnw(msg, keyvals...)
}

// NewZapAdapter creates a new instance of ZapAdapter
func NewZapAdapter(logger *zap.Logger) *ZapAdapter {
	return &ZapAdapter{logger: logger}
}

// Log method for Info level logging
func (z *ZapAdapter) Info(msg string, keyvals ...interface{}) {
	z.logger.Sugar().Infow(msg, keyvals...)
}

// Log method for Error level logging
func (z *ZapAdapter) Error(msg string, keyvals ...interface{}) {
	z.logger.Sugar().Errorw(msg, keyvals...)
}

// With method to add context to the logger
func (z *ZapAdapter) With(keyvals ...interface{}) log.Logger {
	return &ZapAdapter{logger: z.logger.With(zap.Any("context", keyvals))}
}
