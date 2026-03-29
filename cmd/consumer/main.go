package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/FerroO2000/goccia-demo/internal"
)

const connectorSize = 1024

func main() {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	internal.Init(ctx, "consumer-service")

	<-ctx.Done()
}
