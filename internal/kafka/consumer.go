package kafka

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

type Consumer struct {
	consumerGroup sarama.ConsumerGroup
}

func NewConsumer(
	brokers []string,
	groupID string,
) *Consumer {

	config := sarama.NewConfig()

	config.Consumer.Return.Errors = true

	group, err := sarama.NewConsumerGroup(
		brokers,
		groupID,
		config,
	)

	if err != nil {
		log.Fatal(err)
	}

	return &Consumer{
		consumerGroup: group,
	}
}

type Handler struct{}

func (Handler) Setup(
	sarama.ConsumerGroupSession,
) error {
	return nil
}

func (Handler) Cleanup(
	sarama.ConsumerGroupSession,
) error {
	return nil
}

func (Handler) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {

	for message := range claim.Messages() {

		log.Printf(
			"Message received: %s\n",
			string(message.Value),
		)

		session.MarkMessage(message, "")
	}

	return nil
}

func (c *Consumer) Start() {

	handler := Handler{}

	for {

		err := c.consumerGroup.Consume(
			context.Background(),
			[]string{"match-events"},
			handler,
		)

		if err != nil {
			log.Println(err)
		}
	}
}