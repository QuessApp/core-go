package main

import (
	"core/internal/configs"
	"core/internal/database"
	"core/internal/repositories"
	"core/internal/routes"
	"fmt"

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

	authRepository := repositories.NewAuthRepository(db)
	usersRepository := repositories.NewUsersRepository(db)
	questionsRepository := repositories.NewQuestionsRepository(db)
	blocksRepository := repositories.NewBlocksRepository(db)

	AppCtx := &routes.AppCtx{
		App:                 fiber.New(),
		DB:                  db,
		Cfg:                 config,
		QuestionsRepository: questionsRepository,
		BlocksRepository:    blocksRepository,
		UsersRepository:     usersRepository,
		AuthRepository:      authRepository,
	}

	routes.LoadRoutes(AppCtx)
}
