package main

import (
	"context"
	"encoding/json"

	"github.com/FerroO2000/goccia-demo/internal"
	"github.com/FerroO2000/goccia/egress"
	"github.com/FerroO2000/goccia/ingress"
	"github.com/FerroO2000/goccia/processor"
)

type questDBHandler struct {
	processor.CustomHandlerBase
}

func newQuestDBHandler() *questDBHandler {
	return &questDBHandler{}
}

func (h *questDBHandler) Handle(_ context.Context, kafkaMsg *ingress.KafkaMessage, questDBMsg *egress.QuestDBMessage) error {

	pingEvent := &internal.PingEvent{}

	// Decode Kafka's message value
	if err := json.Unmarshal(kafkaMsg.Value, &pingEvent); err != nil {
		return err
	}

	row := egress.NewQuestDBRow("ping_events")
	row.AddColumn(egress.NewQuestDBStringColumn("src_ip", pingEvent.GetSrcIP().String()))
	row.AddColumn(egress.NewQuestDBStringColumn("dst_ip", pingEvent.GetDstIP().String()))
	row.AddColumn(egress.NewQuestDBIntColumn("id", int64(pingEvent.ID)))
	row.AddColumn(egress.NewQuestDBIntColumn("sequence", int64(pingEvent.Seq)))

	questDBMsg.AddRow(row)

	return nil
}
