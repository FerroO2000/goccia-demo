package main

import (
	"context"

	"github.com/FerroO2000/goccia/egress"
	"github.com/FerroO2000/goccia/processor"
)

type pingHandler struct {
	processor.CustomHandlerBase
}

func newPingHandler() *pingHandler {
	return &pingHandler{}
}

func (h *pingHandler) Handle(_ context.Context, ebpfMsg *PingEventMessage, kafkaMsg *egress.KafkaMessage) error {

	// pingEvent := ebpfMsg.Data
	// srcIP := pingEvent.GetSrcIP()
	// dstIP := pingEvent.GetDstIP()

	// h.Telemetry.LogInfo("received ping event", "src_ip", srcIP, "dst_ip", dstIP)

	// // Build the Kafka event
	// kafkaMsg.Topic = "ping_events"
	// kafkaMsg.Key = fmt.Appendf(nil, "%s_%s_%d", srcIP, dstIP, pingEvent.Seq)

	// // Encode the event as JSON
	// value, err := json.Marshal(&pingEvent)
	// if err != nil {
	// 	return err
	// }
	// kafkaMsg.Value = value

	return nil
}
