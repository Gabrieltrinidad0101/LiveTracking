package rabbitMQ

import (
	"github.com/streadway/amqp"
)

type RabbitConsumer struct {
	channel *amqp.Channel
	queue   string
	handler func(delivery amqp.Delivery)
}

func NewRabbitMQConsumer(conn *amqp.Connection, queue string, exchange string, routingKey string, handler func(delivery amqp.Delivery)) (*RabbitConsumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if err := ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil); err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	if err := ch.QueueBind(q.Name, routingKey, exchange, false, nil); err != nil {
		return nil, err
	}

	return &RabbitConsumer{
		channel: ch,
		queue:   q.Name,
		handler: handler,
	}, nil
}

func (c *RabbitConsumer) StartConsuming() error {
	msgs, err := c.channel.Consume(
		c.queue,
		"",
		true,  // auto-ack
		false, // not exclusive
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			c.handler(d)
		}
	}()

	return nil
}

func (c *RabbitConsumer) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
}
