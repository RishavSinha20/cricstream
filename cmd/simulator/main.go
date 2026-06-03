package main

import (
	"time"

	"github.com/RishavSinha20/cricstream/internal/kafka"
	"github.com/RishavSinha20/cricstream/internal/match"
	"github.com/RishavSinha20/cricstream/internal/models"
	"github.com/RishavSinha20/cricstream/internal/utils"
)

func main() {

	utils.InitRandom()

	producer := kafka.NewProducer(
		[]string{"localhost:9092"},
	)

	events := make(chan models.MatchEvent, 100)

	go match.StartMatch(
		"ipl_final_2026",
		events,
	)

	go match.StartMatch(
		"india_vs_australia",
		events,
	)

	for {

		event := <-events

		producer.PublishMatchEvent(event)

		time.Sleep(500 * time.Millisecond)
	}
}
