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

	internal.InitTelemetry(ctx, "consumer-service")
	defer internal.CloseTelemetry()

	// // Setup connectors
	// kafkaToTee := connector.NewRingBuffer[*ingress.KafkaMessage](connectorSize)
	// teeToFileHandler := connector.NewRingBuffer[*ingress.KafkaMessage](connectorSize)
	// handlerToFile := connector.NewRingBuffer[*FileMessage](connectorSize)
	// teeToQuestDBHandler := connector.NewRingBuffer[*ingress.KafkaMessage](connectorSize)
	// handlerToQuestDB := connector.NewRingBuffer[*egress.QuestDBMessage](connectorSize)

	// // Setup stages
	// // 1. Kafka Ingress Stage
	// kafkaConfig := ingress.NewKafkaConfig("ping_events")
	// kafkaStage := ingress.NewKafkaStage(kafkaToTee, kafkaConfig)

	// // 2. Tee Processor Stage
	// teeStage := processor.NewTeeStage(kafkaToTee, teeToFileHandler, teeToQuestDBHandler)

	// // 3.1.1. Custom Processor Stage (File)
	// fileHandlerConfig := processor.NewCustomConfig(goccia.StageRunningModeSingle)
	// fileHandlerConfig.Name = "custom_file_handler"
	// fileHandlerStage := processor.NewCustomStage(newFileHandler(), teeToFileHandler, handlerToFile, fileHandlerConfig)

	// // 3.1.2. File Egress Stage
	// fileEgressConfig := egress.NewFileConfig("ping_events.txt")
	// fileEgressStage := egress.NewFileStage(handlerToFile, fileEgressConfig)

	// // 3.2.1 Custom Processor Stage (QuestDB)
	// questDbHandlerConfig := processor.NewCustomConfig(goccia.StageRunningModeSingle)
	// questDbHandlerConfig.Name = "custom_quest_db_handler"
	// questDbHandlerStage := processor.NewCustomStage(newQuestDBHandler(), teeToQuestDBHandler, handlerToQuestDB, questDbHandlerConfig)

	// // 3.2.2. QuestDB Egress Stage
	// questDBConfig := egress.NewQuestDBConfig(goccia.StageRunningModeSingle)
	// questDBStage := egress.NewQuestDBStage(handlerToQuestDB, questDBConfig)

	// // Setup pipeline
	// pipeline := goccia.NewPipeline()

	// pipeline.AddStage(kafkaStage)
	// pipeline.AddStage(teeStage)
	// pipeline.AddStage(fileHandlerStage)
	// pipeline.AddStage(fileEgressStage)
	// pipeline.AddStage(questDbHandlerStage)
	// pipeline.AddStage(questDBStage)

	// if err := pipeline.Init(ctx); err != nil {
	// 	panic(err)
	// }

	// // Run pipeline
	// go pipeline.Run(ctx)
	// defer pipeline.Close()

	<-ctx.Done()
}
