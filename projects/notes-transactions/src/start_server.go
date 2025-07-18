package src

import (
	"log"
	"notes-transactions/src/notes-transactions/infranstructure/rabbitMQ"

	"github.com/labstack/echo"
	"github.com/streadway/amqp"
)

func handler(delivery amqp.Delivery) {
	print("sadasdasd")
}

func Server(conn *amqp.Connection) {
	e := echo.New()
	consumer, err := rabbitMQ.NewRabbitMQConsumer(conn, "note-transactions:create", "note", "note:create", handler)
	if err != nil {
		log.Fatalf("Error al crear el publisher: %v", err)
	}
	consumer.StartConsuming()

	e.Start("0.0.0.0:9000")
}
