package main

import (
	"log"
	"notes-transactions/src"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://user:password@192.168.49.2:31672/")

	if err != nil {
		log.Fatalf("Error al conectar a RabbitMQ: %v", err)
	}
	defer conn.Close()
	src.Server(conn)
}
