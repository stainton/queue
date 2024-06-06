package queue

import (
	"context"
	"time"
)

func (q *Queue) RegisterConsumer(ctx context.Context, consumer func([]byte) error, queueName, routingKey, consumerName string) error {
	if _, err := q.channel.QueueDeclare(queueName, true, false, false, false, nil); err != nil {
		return err
	}
	if err := q.channel.QueueBind(queueName, routingKey, exchange, false, nil); err != nil {
		return err
	}
	messageChannel, err := q.channel.ConsumeWithContext(ctx, queueName, consumerName, false, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for {
			ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			select {
			case <-ctxTimeout.Done():
				cancel()
				return
			case message, ok := <-messageChannel:
				cancel()
				if !ok {
					return
				}
				if err = consumer(message.Body); err != nil {
					err = message.Nack(false, true)
					if err != nil {
						return
					}
				} else {
					message.Ack(false)
				}
			case <-ctx.Done():
				cancel()
				return
			}
		}
	}()
	return nil
}
