package main

import (
    "github.com/rabbitmq/amqp091-go"
)


func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

  	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
}