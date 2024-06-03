package queue

import (
	"time"
)

type Consumer interface {
	Name() string
	Queue() string
	RoutingKey() string
	Consumer([]byte) error
}

func (q *Queue) RegisterConsumer(consumer Consumer, stopChan <-chan struct{}, errChan chan<- error) error {
	_, err := q.channel.QueueDeclare(consumer.Queue(), true, false, false, false, nil)
	if err != nil {
		return err
	}
	err = q.channel.QueueBind(consumer.Queue(), consumer.RoutingKey(), exchange, false, nil)
	if err != nil {
		return err
	}
	messageChannel, err := q.channel.Consume(consumer.Queue(), consumer.Name(), false, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case message := <-messageChannel:
				if e := consumer.Consumer(message.Body); e != nil {
					e = message.Nack(false, true)
					errChan <- e
				} else {
					message.Ack(false)
				}
			case <-stopChan:
				return
			default:
				time.Sleep(5 * time.Second)
			}
		}
	}()
	return nil
}
