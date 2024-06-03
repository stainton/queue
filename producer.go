package queue

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer interface {
	Publish(string, interface{}) error
}

// publish publishes data to the queue, data will be discarded if no queue is ready
func (q *Queue) Publish(routingKey string, data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return q.channel.PublishWithContext(ctx,
		exchange,
		routingKey,
		true, // undeliverable if no queue bound
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
}
