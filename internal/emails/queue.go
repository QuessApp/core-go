package emails

import (
	"github.com/streadway/amqp"
)

// DeclareQueue declares an AMQP queue on the provided channel with the specified queueName.
// It takes a pointer to an AMQP channel object (ch) and a string (queueName) representing the name of the queue.
// It uses the ch.QueueDeclare function to declare the queue on the channel, specifying the queueName and other properties such as durability and deletion policy.
// The function returns a pointer to the declared queue object and an error object if an error occurs.
func DeclareQueue(ch *amqp.Channel, queueName string) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	return &q, err
}
