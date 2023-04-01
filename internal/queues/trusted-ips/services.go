package trustedips

import (
	"fmt"
	"log"
	"strings"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/toolkit/queue"
	"github.com/streadway/amqp"
)

// SendIPToQueue send IP to queue
func SendIPToQueue(cfg *configs.Conf, ch *amqp.Channel, q *amqp.Queue, ip, sendTo string) {
	isLocalHost := ip == "127.0.0.1" || ip == "0.0.0.0"

	if isLocalHost {
		return
	}

	normalizedIP := strings.ReplaceAll(ip, ".", "")
	message := fmt.Sprintf("%s-%s", sendTo, normalizedIP)

	if err := queue.Publish(ch, q.Name, cfg.Crypto.Key, []byte(message)); err != nil {
		log.Fatalf("fail to send email to user %s \n", err)
	}
}
