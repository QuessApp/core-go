package main

import (
	"core/cmd/api"
	"core/configs"
	"core/internal/auth"
	"core/internal/blocks"
	"core/internal/database"
	"core/internal/questions"

	"core/internal/users"

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

	app := fiber.New()

	authRepository := auth.NewAuthRepository(db)
	usersRepository := users.NewRepository(db)
	questionsRepository := questions.NewRepository(db)
	blocksRepository := blocks.NewRepository(db)

	AppCtx := &configs.AppCtx{
		App:              app,
		DB:               db,
		Cfg:              config,
		MessageQueueConn: conn,
		MessageQueueCh:   ch,
	}

	q, err := questions.DeclareEmailsQueue(AppCtx)

	if err != nil {
		log.Fatalln(err)
	}

	AppCtx.SendEmailsQueue = q

	api.LoadRoutes(AppCtx, authRepository, usersRepository, blocksRepository, questionsRepository)
}
