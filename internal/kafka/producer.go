package kafka

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/RishavSinha20/cricstream/internal/models"
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer(brokers []string) *Producer {

	config := sarama.NewConfig()

	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(
		brokers,
		config,
	)

	if err != nil {
		log.Fatal(err)
	}

	return &Producer{
		producer: producer,
	}
}

func (p *Producer) PublishMatchEvent(
	event models.MatchEvent,
) {

	data, err := json.Marshal(event)

	if err != nil {
		log.Println(err)
		return
	}

	message := &sarama.ProducerMessage{
		Topic: "match-events",

		Key: sarama.StringEncoder(event.MatchID),

		Value: sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(message)

	if err != nil {
		log.Println(err)
	}
}