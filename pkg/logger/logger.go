package logger

import (
	"context"

	"go.uber.org/zap"
)

type LoggerCtx struct {
	l *zap.SugaredLogger
	ctx context.Context
}

func NewLoggerCtx(ctx context.Context) *LoggerCtx {
	logger, _ := zap.NewProduction()
	
	sugar := logger.Sugar()
	return &LoggerCtx{ctx: ctx, l: sugar}
}

func (lctx *LoggerCtx) SimpleMessageWithContextField() {

	message := lctx.ctx.Value("message")

	lctx.l.Infow("here is the new message", "message", message,)
}