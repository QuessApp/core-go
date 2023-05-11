package emails

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/users"
	"github.com/quessapp/core-go/pkg/i18n"
	toolkitEntities "github.com/quessapp/toolkit/entities"
	"github.com/quessapp/toolkit/queue"
)

// SendEmailNewQuestionReceived sends an email notification to the user that receives a new question.
// The email contains the content of the question, the sender's name, and whether the question is anonymous or not.
// The email is encrypted and sent using an AMQP channel and queue.
func SendEmailNewQuestionReceived(handlerCtx *configs.HandlersCtx, content string, isAnonymous bool, userToSendQuestion *users.User, userThatIsSendingQuestion *users.User) {
	var subject string

	if isAnonymous {
		subject = i18n.Translate(handlerCtx, "emails_new_question_subject")
	} else {
		translatedSubject := i18n.Translate(handlerCtx, "emails_new_question_subject")
		subject = fmt.Sprintf(translatedSubject, userThatIsSendingQuestion.Nick)
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

	if err := queue.Publish(handlerCtx.MessageQueueCh, handlerCtx.EmailsQueue.Name, handlerCtx.Cfg.Crypto.Key, emailParsed); err != nil {
		log.Fatalf("fail to send email to user %s \n", err)
	}
}

// SendEmailForgotPassword sends an email notification to the user that wants to reset password.
// The email contains the code that the user will use to reset the password.
// The email is encrypted and sent using an AMQP channel and queue.
func SendEmailForgotPassword(handlerCtx *configs.HandlersCtx, code string, userToSendEmail *users.User) error {
	translatedBody := i18n.Translate(handlerCtx, "emails_forgot_password_body")
	url := fmt.Sprintf("%s?code=%s", handlerCtx.Cfg.App.FrontendURL, code)

	email := toolkitEntities.Email{
		To:      userToSendEmail.Email,
		Subject: i18n.Translate(handlerCtx, "emails_forgot_password_subject"),
		Body:    fmt.Sprintf(translatedBody, url),
	}

	emailParsed, err := json.Marshal(email)

	if err != nil {
		log.Fatalf("fail to marshal %s", err)
		return err
	}

	if err := queue.Publish(handlerCtx.MessageQueueCh, handlerCtx.EmailsQueue.Name, handlerCtx.Cfg.Crypto.Key, emailParsed); err != nil {
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
func SendEmailPasswordChanged(handlerCtx *configs.HandlersCtx, userToSendEmail *users.User) error {
	email := toolkitEntities.Email{
		To:      userToSendEmail.Email,
		Subject: i18n.Translate(handlerCtx, "emails_password_changed_subject"),
		Body:    i18n.Translate(handlerCtx, "emails_password_changed_body"),
	}

	emailParsed, err := json.Marshal(email)

	if err != nil {
		log.Fatalf("fail to marshal %s", err)
		return err
	}

	if err := queue.Publish(handlerCtx.MessageQueueCh, handlerCtx.EmailsQueue.Name, handlerCtx.Cfg.Crypto.Key, emailParsed); err != nil {
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
func SendEmailThanksForReporting(handlerCtx *configs.HandlersCtx, userToSendEmail *users.User) {
	email := toolkitEntities.Email{
		To:      userToSendEmail.Email,
		Subject: i18n.Translate(handlerCtx, "emails_report_sent_subject"),
		Body:    i18n.Translate(handlerCtx, "emails_report_sent_body"),
	}

	emailParsed, err := json.Marshal(email)

	if err != nil {
		log.Fatalf("fail to marshal %s", err)
		return
	}

	if err := queue.Publish(handlerCtx.MessageQueueCh, handlerCtx.EmailsQueue.Name, handlerCtx.Cfg.Crypto.Key, emailParsed); err != nil {
		log.Fatalf("fail to send email to user %s \n", err)
		return
	}
}
