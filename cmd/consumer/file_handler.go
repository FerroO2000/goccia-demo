package main

import (
	"context"

	"github.com/FerroO2000/goccia/ingress"
	"github.com/FerroO2000/goccia/processor"
)

type FileMessage struct {
	Row string
}

func (fm *FileMessage) Destroy() {}

func (fm *FileMessage) GetBytes() []byte {
	return []byte(fm.Row)
}

type fileHandler struct {
	processor.CustomHandlerBase
}

func newFileHandler() *fileHandler {
	return &fileHandler{}
}

func (h *fileHandler) Handle(_ context.Context, kafkaMsg *ingress.KafkaMessage, fileMsg *FileMessage) error {

	// pingEvent := &internal.PingEvent{}

	// // Decode Kafka's message value
	// if err := json.Unmarshal(kafkaMsg.Value, &pingEvent); err != nil {
	// 	return err
	// }

	// // Build the file row
	// fileMsg.Row = fmt.Sprintf("src_ip: %s, dst_ip: %s, id: %d, seq: %d\n",
	// 	pingEvent.GetSrcIP(), pingEvent.GetDstIP(), pingEvent.ID, pingEvent.Seq)

	return nil
}
