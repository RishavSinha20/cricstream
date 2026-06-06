package main

import (
	"github.com/RishavSinha20/cricstream/internal/kafka"
	"github.com/RishavSinha20/cricstream/internal/redisstore"
	"github.com/RishavSinha20/cricstream/internal/worker"
)

func main() {

	redisstore.Init()

	pool := worker.NewPool(5)

	consumer := kafka.NewConsumer(
		[]string{"localhost:9092"},
		"analytics-group",
		pool,
	)

	consumer.Start()
}