package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/FerroO2000/goccia-demo/internal"
)

const connectorSize = 1024

func main() {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	// Get the network interface
	ifname := "eth0"
	if len(os.Args) > 1 {
		ifname = os.Args[1]
	}

	log.Printf("Using network interface: %s", ifname)

	internal.Init(ctx, "producer-service")

	<-ctx.Done()
}
