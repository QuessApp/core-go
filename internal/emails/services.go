package emails

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/users"
	"github.com/quessapp/toolkit/crypto"
	toolkitEntities "github.com/quessapp/toolkit/entities"

	"github.com/streadway/amqp"
)

// TODO: use AWS SQS instead of rabbitmq.

// SendEmailNewQuestionReceived sends an email notification to the user that receives a new question.
// The email contains the content of the question, the sender's name, and whether the question is anonymous or not.
// The email is encrypted and sent using an AMQP channel and queue.
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

// SendEmailForgotPassword sends an email notification to the user that wants to reset password.
// The email contains the code that the user will use to reset the password.
// The email is encrypted and sent using an AMQP channel and queue.
func SendEmailForgotPassword(cfg *configs.Conf, ch *amqp.Channel, q *amqp.Queue, code string, userToSendEmail *users.User) error {
	email := toolkitEntities.Email{
		To:      userToSendEmail.Email,
		Subject: "Recuperação de senha",
		Body:    fmt.Sprintf(`Você solicitou a recuperação de senha. Seu código é: %v`, code),
	}

	emailParsed, err := json.Marshal(email)

	if err != nil {
		log.Fatalf("fail to marshal %s", err)
		return err
	}

	encryptedMsg, err := crypto.Encrypt(string(emailParsed), cfg.CipherKey)

	if err != nil {
		log.Fatalf("fail to encrypt email email to user %s \n", err)
		return err
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
		return err
	}

	return nil
}
