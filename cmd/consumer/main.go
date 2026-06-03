package main

import "github.com/RishavSinha20/cricstream/internal/kafka"

func main() {

	consumer := kafka.NewConsumer(
		[]string{"localhost:9092"},
		"analytics-group",
	)

	consumer.Start()
}
