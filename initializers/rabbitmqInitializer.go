package initializers

import (
	"rabbitmq/utilities"

	amqp "github.com/rabbitmq/amqp091-go"
)

func RabbitmqInitializer() {
	conn, err := amqp.Dial("amqp://user:pass@localhost:5672/")
	utilities.ErrorHandler(err)
	defer conn.Close()
}
