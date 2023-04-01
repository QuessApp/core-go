package emails

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/users"
	toolkitEntities "github.com/quessapp/toolkit/entities"
	"github.com/quessapp/toolkit/queue"

	"github.com/streadway/amqp"
)

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

	if err := queue.Publish(ch, q.Name, cfg.Crypto.Key, emailParsed); err != nil {
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

	if err := queue.Publish(ch, q.Name, cfg.Crypto.Key, emailParsed); err != nil {
		log.Fatalf("fail to send email to user %s \n", err)
		return err
	}

	return nil
}

// SendEmailPasswordChanged sends an email to the user with the new password.
// It takes in the configuration object cfg, the message queue channel ch, and the email queue q.
// It also takes in the userToSendEmail object which contains the email address of the user.
// It constructs an email object with the To field set to the email of the userToSendEmail.
// It encrypts the email using the encryption key in cfg.CipherKey and sends it to the email queue q using the message queue channel ch.
// It returns an error if there was a problem marshaling, encrypting, or sending the email.
func SendEmailPasswordChanged(cfg *configs.Conf, ch *amqp.Channel, q *amqp.Queue, userToSendEmail *users.User) error {
	email := toolkitEntities.Email{
		To:      userToSendEmail.Email,
		Subject: "Senha alterada",
		Body:    "Sua senha foi alterada com sucesso. Se você não solicitou essa alteração ou acha que é um engano, entre em contato com o suporte.",
	}

	emailParsed, err := json.Marshal(email)

	if err != nil {
		log.Fatalf("fail to marshal %s", err)
		return err
	}

	if err := queue.Publish(ch, q.Name, cfg.Crypto.Key, emailParsed); err != nil {
		log.Fatalf("fail to send email to user %s \n", err)
		return err
	}

	return nil
}

// SendEmailThanksForReporting sends an email to the user that reported a question.
// It takes in the configuration object cfg, the message queue channel ch, and the email queue q.
// It also takes in the userToSendEmail object which contains the email address of the user.
// It constructs an email object with the To field set to the email of the userToSendEmail.
// It encrypts the email using the encryption key in cfg.CipherKey and sends it to the email queue q using the message queue channel ch.
// It returns an error if there was a problem marshaling, encrypting, or sending the email.
func SendEmailThanksForReporting(cfg *configs.Conf, ch *amqp.Channel, q *amqp.Queue, userToSendEmail *users.User) {
	email := toolkitEntities.Email{
		To:      userToSendEmail.Email,
		Subject: "Denúncia enviada",
		Body:    "Agradecemos por nos ajudar a manter a comunidade segura. Sua denúncia será analisada e, se necessário, tomaremos as devidas providências.",
	}

	emailParsed, err := json.Marshal(email)

	if err != nil {
		log.Fatalf("fail to marshal %s", err)
		return
	}

	if err := queue.Publish(ch, q.Name, cfg.Crypto.Key, emailParsed); err != nil {
		log.Fatalf("fail to send email to user %s \n", err)
		return
	}
}
