package main

import (
	"context"

	"github.com/MichaelGenchev/uberZAP/pkg/logger"
)

type contextMessageKey interface{}

func main() {
	logger := logger.NewLogger()
	ctxBackground := context.Background()
	context := context.WithValue(ctxBackground, "id", "1234")

	logger.WithContext(context).Info(context, "Hello")
}


