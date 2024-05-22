package main

import (
	"log"
	"rabbitmq/initializers"
	"rabbitmq/utilities"

	amqp "github.com/rabbitmq/amqp091-go"
)

func init() {
	initializers.RabbitmqInitializer()
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@172.27.128.1:15672")
	utilities.ErrorHandler(err)
	defer conn.Close()
	log.Println("final")
}
