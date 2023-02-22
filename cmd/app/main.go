package main

import (
	"core/internal/configs"
	"core/internal/database"
	"core/internal/repositories"
	"core/internal/routes"
	"fmt"
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

	routes.LoadRoutes(db, config, questionsRepository, authRepository, usersRepository, blocksRepository)
}
