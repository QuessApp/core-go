package questions

import (
	"core/configs"
	"core/internal/users"
	"encoding/json"
	"fmt"
	"log"

	"github.com/kuriozapp/toolkit/crypto"
	toolkitEntities "github.com/kuriozapp/toolkit/entities"
	"github.com/streadway/amqp"
)

// DeclareEmailsQueue declares an queue in MQ to send emails through app.
func DeclareEmailsQueue(AppCtx *configs.AppCtx) (*amqp.Queue, error) {
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

// SendEmailNewQuestionReceived sends an email to MQ.
func SendEmailNewQuestionReceived(cfg *configs.Conf, ch *amqp.Channel, q *amqp.Queue, payload *CreateQuestionDTO, userToSendQuestion *users.User, userThatIsSendingQuestion *users.User) {
	var subject string

	if payload.IsAnonymous {
		subject = "Você recebeu uma pergunta anônima"
	} else {
		subject = fmt.Sprintf("Você recebeu uma pergunta de @%s", userThatIsSendingQuestion.Nick)
	}

	email := toolkitEntities.Email{
		To:      userToSendQuestion.Email,
		Subject: subject,
		Body:    fmt.Sprintf(`"%v" - %v`, payload.Content, userThatIsSendingQuestion.Name),
	}

	emailParsed, err := json.Marshal(email)

	if err != nil {
		log.Fatalf("fail to marshal %s", err)
	}

	encryptedMsg, err := crypto.Encrypt(string(emailParsed), cfg.CipherKey)

	if err != nil {
		log.Fatalf("fail to to encrypt email email to user %s \n", err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(encryptedMsg),
		})

	if err != nil {
		log.Fatalf("fail to send email to user %s \n", err)
	}
}
