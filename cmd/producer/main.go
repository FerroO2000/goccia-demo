package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/FerroO2000/goccia-demo/internal"
	"github.com/FerroO2000/goccia/ingress"
)

const connectorSize = 1024

type PingEventMessage = ingress.EBPFMessage[internal.PingEvent]

func main() {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	// Get the network interface
	ifname := "eth0"
	if len(os.Args) > 1 {
		ifname = os.Args[1]
	}

	log.Printf("Using network interface: %s", ifname)

	internal.InitTelemetry(ctx, "producer-service")
	defer internal.CloseTelemetry()

	// // Setup connectors
	// ebpfToHandler := connector.NewRingBuffer[*PingEventMessage](connectorSize)
	// handlerToKafka := connector.NewRingBuffer[*egress.KafkaMessage](connectorSize)

	// // Setup stages

	// // 1. eBPF Ingress Stage
	// ebpfConfig := ingress.NewEBPFConfig(
	// 	loadBpf,
	// 	func(objs *bpfObjects) (link.Link, error) {
	// 		iface, err := net.InterfaceByName(ifname)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return link.AttachXDP(link.XDPOptions{
	// 			Program:   objs.PingMonitor,
	// 			Interface: iface.Index,
	// 		})
	// 	},
	// 	func(objs *bpfObjects) *ebpf.Map {
	// 		return objs.PingEvents
	// 	},
	// )

	// ebpfStage := ingress.NewEBPFStage(ebpfToHandler, ebpfConfig)

	// // 2. Custom Processor Stage
	// pingHandlerConfig := processor.NewCustomConfig(goccia.StageRunningModeSingle)
	// pingHandlerConfig.Name = "custom_ping_handler"
	// pingHandlerStage := processor.NewCustomStage(newPingHandler(), ebpfToHandler, handlerToKafka, pingHandlerConfig)

	// // 3. Kafka Egress Stage
	// kafkaConfig := egress.NewKafkaConfig(goccia.StageRunningModeSingle)
	// kafkaStage := egress.NewKafkaStage(handlerToKafka, kafkaConfig)

	// // Setup pipeline
	// pipeline := goccia.NewPipeline()

	// pipeline.AddStage(ebpfStage)
	// pipeline.AddStage(pingHandlerStage)
	// pipeline.AddStage(kafkaStage)

	// if err := pipeline.Init(ctx); err != nil {
	// 	panic(err)
	// }

	// // Run pipeline
	// go pipeline.Run(ctx)
	// defer pipeline.Close()

	<-ctx.Done()
}
