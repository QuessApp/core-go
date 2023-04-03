package trustedips

import (
	"fmt"
	"log"
	"strings"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/toolkit/queue"
	"github.com/streadway/amqp"
)

// SendIPToQueue sends an IP address to an AMQP queue using the provided configuration, channel, queue, and destination queue name.
// If the IP address is local (either "127.0.0.1" or "0.0.0.0"), the function returns without sending anything to the queue.
// The IP address is normalized by removing dots (".") and concatenating it with the destination queue name to create the message.
// The message is then published to the specified AMQP queue using the Publish function from the "queue" package.
// If there is an error during the publishing process, the function logs a fatal error message with the details of the error.
func SendIPToQueue(cfg *configs.Conf, ch *amqp.Channel, q *amqp.Queue, ip, sendToEmail string) {
	isLocalHost := ip == "127.0.0.1" || ip == "0.0.0.0"

	if isLocalHost {
		return
	}

	normalizedIP := strings.ReplaceAll(ip, ".", "")
	message := fmt.Sprintf("%s-%s", sendToEmail, normalizedIP)

	if err := queue.Publish(ch, q.Name, cfg.Crypto.Key, []byte(message)); err != nil {
		log.Fatalf("fail to send email to user %s \n", err)
	}
}
