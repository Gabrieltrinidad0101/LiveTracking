package rabbitmq

import (
	"github.com/streadway/amqp"
)

type JsonToBytes func(p interface{}) ([]byte, error)

type rabbitMQPublisher struct {
	channel     *amqp.Channel
	exchange    string
	jsonToBytes JsonToBytes
}

func NewRabbitMQ(conn *amqp.Connection, exchange string, type_ string, jsonToBytes JsonToBytes) (*rabbitMQPublisher, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	err = channel.ExchangeDeclare(exchange, type_, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return &rabbitMQPublisher{
		channel,
		exchange,
		jsonToBytes,
	}, nil
}

func (p *rabbitMQPublisher) Publish(routingKey string, body interface{}) error {
	bytes, err := p.jsonToBytes(body)
	if err != nil {
		return err
	}
	return p.channel.Publish(
		p.exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		},
	)
}

func (p *rabbitMQPublisher) Close() {
	p.channel.Close()
}
