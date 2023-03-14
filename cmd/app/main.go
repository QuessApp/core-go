package main

import (
	"github.com/quessapp/core-go/configs"

	"github.com/quessapp/core-go/internal/auth"
	"github.com/quessapp/core-go/internal/blocks"
	"github.com/quessapp/core-go/internal/database"
	"github.com/quessapp/core-go/internal/emails"
	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/internal/reports"
	"github.com/quessapp/core-go/internal/router"
	"github.com/quessapp/core-go/internal/users"

	"fmt"
	"log"

	"github.com/quessapp/toolkit/queue"
	"github.com/quessapp/toolkit/s3"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := configs.LoadConfig(".")

	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	DatabaseURIConn := fmt.Sprintf("%s:%s", config.DBHost, config.DBPort)

	db, err := database.Connect(DatabaseURIConn, config.DBName)

	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	conn, ch := queue.Connect(config.MessageQueueURI)
	defer conn.Close()

	app := fiber.New()

	authRepository := auth.NewAuthRepository(db)
	usersRepository := users.NewRepository(db)
	questionsRepository := questions.NewRepository(db)
	blocksRepository := blocks.NewRepository(db)
	reportsRepository := reports.NewRepository(db)

	S3Client, err := s3.Configure(&config.S3Region, &s3.S3Credentials{
		AccessKey: config.S3AccessKey,
		Secret:    config.S3Secret,
		Token:     config.S3Token,
	})

	if err != nil {
		log.Fatalf("failed to configure S3 client: %s", err)
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
		log.Fatalf("failed to declare emails queue: %s", err)
	}

	AppCtx.SendEmailsQueue = q

	router.Setup(AppCtx, authRepository, usersRepository, blocksRepository, questionsRepository, reportsRepository)
}
