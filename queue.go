package queue

import (
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/pflag"
)

type Options interface {
	Validate() error
}

type Queue struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

var url string
var exchange string

func AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&url, "rabbit", "", "the url for connecting to RabbitMQ")
	fs.StringVar(&exchange, "exchange", "", "the exchange for consumer and producer")
}

func NewQueue() (*Queue, error) {
	if url == "" {
		return nil, errors.New("url is required for connecting to RabbitMQ")
	}
	// connect to RabbitMQ
	connection, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := connection.Channel()
	if err != nil {
		connection.Close()
		return nil, err
	}

	// declare exchange
	if err = ch.ExchangeDeclare(
		exchange,
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		ch.Close()
		connection.Close()
		return nil, err
	}

	return &Queue{
		connection: connection,
		channel:    ch,
	}, nil
}
