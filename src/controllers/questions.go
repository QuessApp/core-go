package controllers

import (
	"core/src/database"
	"core/src/models"
	"core/src/repositories"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateQuestion(c *fiber.Ctx) error {
	var payload models.Question

	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		panic(err)
	}

	db, _ := database.Connect()

	repository := repositories.NewQuestionsRepository(db)
	createdQuestion, err := repository.Create(payload)

	if err != nil {
		return fiber.NewError(403, fmt.Sprint(err))
	}

	fmt.Println(createdQuestion)

	return errors.New("foo")
}

func FindQuestionById() {}

func FindQuestions() {}

func DeleteQuestion() {}
