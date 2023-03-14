package emails

import (
	"github.com/quessapp/core-go/configs"

	"github.com/streadway/amqp"
)

// DeclareQueue is a function that declares a new AMQP queue with the given name and attributes.
//
// The function takes an argument of type *configs.AppCtx, which is a struct containing the application context
// and configuration settings. The function uses the MessageQueueCh field of the AppCtx argument to declare a queue
// with the given queue name, durable flag, delete-when-unused flag, exclusive flag, no-wait flag, and arguments.
// If the queue is successfully declared, a pointer to the queue is returned along with nil error. Otherwise, an
// error is returned.
func DeclareQueue(AppCtx *configs.AppCtx) (*amqp.Queue, error) {
	q, err := AppCtx.MessageQueueCh.QueueDeclare(
		AppCtx.Cfg.SendEmailsQueueName, // name
		true,                           // durable
		false,                          // delete when unused
		false,                          // exclusive
		false,                          // no-wait
		nil,                            // arguments
	)

	return &q, err
}
