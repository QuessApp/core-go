package usecases

import (
	"core/src/models"
	"core/src/repositories"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterUser reads payload from request body then try to create a new user in database.
func RegisterUser(c *fiber.Ctx, db *mongo.Database) (interface{}, error) {
	var payload models.User

	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		return nil, err
	}

	repository := repositories.NewUsersRepository(db)

	createdUser, err := repository.Create(payload)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
