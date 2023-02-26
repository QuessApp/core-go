package queues

import (
	"core/internal/configs"

	"github.com/streadway/amqp"
)

// DeclareSendEmailsQueue declares an queue in MQ to send emails through app.
func DeclareSendEmailsQueue(AppCtx *configs.AppCtx) (*amqp.Queue, error) {
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
