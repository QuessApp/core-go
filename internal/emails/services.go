package emails

import (
	"core/configs"
	"core/internal/users"
	"encoding/json"
	"fmt"
	"log"

	"github.com/quessapp/toolkit/crypto"
	toolkitEntities "github.com/quessapp/toolkit/entities"
	"github.com/streadway/amqp"
)

// SendEmailNewQuestionReceived sends an email to MQ.
func SendEmailNewQuestionReceived(cfg *configs.Conf, ch *amqp.Channel, q *amqp.Queue, content string, isAnonymous bool, userToSendQuestion *users.User, userThatIsSendingQuestion *users.User) {
	var subject string

	if isAnonymous {
		subject = "Você recebeu uma pergunta anônima"
	} else {
		subject = fmt.Sprintf("Você recebeu uma pergunta de @%s", userThatIsSendingQuestion.Nick)
	}

	email := toolkitEntities.Email{
		To:      userToSendQuestion.Email,
		Subject: subject,
		Body:    fmt.Sprintf(`"%v" - %v`, content, userThatIsSendingQuestion.Name),
	}

	emailParsed, err := json.Marshal(email)

	if err != nil {
		log.Fatalf("fail to marshal %s", err)
	}

	encryptedMsg, err := crypto.Encrypt(string(emailParsed), cfg.CipherKey)

	if err != nil {
		log.Fatalf("fail to encrypt email email to user %s \n", err)
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(encryptedMsg),
		})

	if err != nil {
		log.Fatalf("fail to send email to user %s \n", err)
	}
}
