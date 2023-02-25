package entities

import (
	"core/internal/configs"
	"core/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

// AppCtx is a global model for app. It defines the router, db, config, repositories, etc.
// Use AppCtx to avoid long function params.
type AppCtx struct {
	App                 *fiber.App
	DB                  *mongo.Database
	Cfg                 *configs.Conf
	MessageQueueConn    *amqp.Connection
	MessageQueueCh      *amqp.Channel
	QuestionsRepository *repositories.Questions
	BlocksRepository    *repositories.Blocks
	UsersRepository     *repositories.Users
	AuthRepository      *repositories.Auth
}
