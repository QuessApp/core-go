package main

import (
	"core/internal/configs"
	"core/internal/database"
	"core/internal/queues"
	"core/internal/repositories"
	"core/internal/routes"
	"fmt"
	"log"

	"github.com/kuriozapp/toolkit/queue"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	URIConn := fmt.Sprintf("%s:%s", config.DBHost, config.DBPort)

	db, err := database.Connect(URIConn, config.DBName)

	if err != nil {
		panic(err)
	}

	conn, ch := queue.Connect(config.MessageQueueURI)
	defer conn.Close()

	authRepository := repositories.NewAuthRepository(db)
	usersRepository := repositories.NewUsersRepository(db)
	questionsRepository := repositories.NewQuestionsRepository(db)
	blocksRepository := repositories.NewBlocksRepository(db)

	AppCtx := &configs.AppCtx{
		App:                 fiber.New(),
		DB:                  db,
		Cfg:                 config,
		MessageQueueConn:    conn,
		MessageQueueCh:      ch,
		QuestionsRepository: questionsRepository,
		BlocksRepository:    blocksRepository,
		UsersRepository:     usersRepository,
		AuthRepository:      authRepository,
	}

	q, err := queues.DeclareSendEmailsQueue(AppCtx)

	if err != nil {
		log.Fatalln(err)
	}

	AppCtx.SendEmailsQueue = q

	routes.LoadRoutes(AppCtx)
}
