package main

import (
	"context"

	"github.com/MichaelGenchev/uberZAP/pkg/logger"
)

type contextMessageKey interface{}

func main() {
	var key contextMessageKey = "message"
	ctx := context.WithValue(context.Background(), key, "My first message")

	loggerCtx := logger.NewLoggerCtx(ctx)

	loggerCtx.SimpleMessageWithContextField()
}


