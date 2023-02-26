package services

import (
	"core/internal/configs"
	"core/internal/dtos"
	internalEntities "core/internal/entities"
	"encoding/json"
	"fmt"
	"log"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"

	"github.com/kuriozapp/toolkit/crypto"
	"github.com/streadway/amqp"
)

// SendNewQuestionReceivedEmail sends an email to MQ.
func SendNewQuestionReceivedEmail(cfg *configs.Conf, ch *amqp.Channel, q *amqp.Queue, payload *dtos.CreateQuestionDTO, userToSendQuestion *internalEntities.User, userThatIsSendingQuestion *internalEntities.User) {
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
