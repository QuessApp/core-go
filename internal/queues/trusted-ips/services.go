package trustedips

import (
	"encoding/json"
	"log"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/toolkit/queue"
	"github.com/streadway/amqp"
)

type Message struct {
	SendToEmail string
	IP          string
	Locale      string
}

// SendIPToQueue sends the IP address information to a message queue.
// If the IP address is a local host (127.0.0.1 or 0.0.0.0), the function returns without sending the message to the queue.
//
// The function creates a Message struct with the sendToEmail, IP, and Locale fields.
// It then marshals the Message struct into JSON format.
// If an error occurs during marshaling, a fatal log is printed.
//
// Finally, the function publishes the marshaled message to the specified queue using the queue.Publish function,
// encrypting the message using the crypto key from the configuration object.
// If an error occurs during publishing, a fatal log is printed.
func SendIPToQueue(cfg *configs.Conf, ch *amqp.Channel, q *amqp.Queue, locale, ip, sendToEmail string) {
	isLocalHost := ip == "127.0.0.1" || ip == "0.0.0.0"

	if isLocalHost {
		return
	}

	msg := Message{
		SendToEmail: sendToEmail,
		IP:          ip,
		Locale:      locale,
	}

	m, err := json.Marshal(msg)

	if err != nil {
		log.Fatalf("fail to marshal message %s \n", err)
	}

	if err := queue.Publish(ch, q.Name, cfg.Crypto.Key, m); err != nil {
		log.Fatalf("fail to send email to user %s \n", err)
	}
}
