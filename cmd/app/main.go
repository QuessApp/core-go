package main

import (
	"core/configs"
	"core/internal/auth"
	"core/internal/blocks"
	"core/internal/database"
	"core/internal/emails"
	"core/internal/questions"
	"core/internal/router"
	"core/internal/users"

	"fmt"
	"log"

	"github.com/quessapp/toolkit/queue"
	"github.com/quessapp/toolkit/s3"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	DatabaseURIConn := fmt.Sprintf("%s:%s", config.DBHost, config.DBPort)

	db, err := database.Connect(DatabaseURIConn, config.DBName)

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

	S3Client, err := s3.Configure(&config.S3Region, &s3.S3Credentials{
		AccessKey: config.S3AccessKey,
		Secret:    config.S3Secret,
		Token:     config.S3Token,
	})

	if err != nil {
		log.Fatalln(err)
	}

	AppCtx := &configs.AppCtx{
		App:            app,
		DB:             db,
		Cfg:            config,
		MessageQueueCh: ch,
		S3Client:       S3Client,
	}

	q, err := emails.DeclareQueue(AppCtx)

	if err != nil {
		log.Fatalln(err)
	}

	AppCtx.SendEmailsQueue = q

	router.Setup(AppCtx, authRepository, usersRepository, blocksRepository, questionsRepository)
}
