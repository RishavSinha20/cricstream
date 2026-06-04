package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"

	"github.com/RishavSinha20/cricstream/internal/models"
	"github.com/RishavSinha20/cricstream/internal/worker"
)

type Consumer struct {
	consumerGroup sarama.ConsumerGroup
	pool          *worker.Pool
}

func NewConsumer(
	brokers []string,
	groupID string,
	pool *worker.Pool,
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
		pool:          pool,
	}
}

type Handler struct {
	pool *worker.Pool
}

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

func (h Handler) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {

	for message := range claim.Messages() {

		var event models.MatchEvent

		err := json.Unmarshal(
			message.Value,
			&event,
		)

		if err != nil {
			log.Println(err)
			continue
		}

		h.pool.JobQueue <- event

		session.MarkMessage(message, "")
	}

	return nil
}

func (c *Consumer) Start() {

	handler := Handler{
		pool : c.pool,
	}

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
