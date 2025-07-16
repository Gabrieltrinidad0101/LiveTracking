package rabbitmq

import (
	"github.com/streadway/amqp"
)

type rabbitMQPublisher struct {
	channel  *amqp.Channel
	exchange string
}

func NewRabbitMQ(conn *amqp.Connection, exchange string, type_ string) (error, *rabbitMQPublisher) {
	channel, err := conn.Channel()
	if err != nil {
		return err, nil
	}
	err = channel.ExchangeDeclare(exchange, type_, true, false, false, false, nil)
	if err != nil {
		return err, nil
	}
	return nil, &rabbitMQPublisher{
		channel,
		exchange,
	}
}

func (p *rabbitMQPublisher) Publish(routingKey string, body []byte) error {
	return p.channel.Publish(
		p.exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (p *rabbitMQPublisher) Close() {
	p.channel.Close()
}
