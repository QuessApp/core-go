package usecases

import (
	helpers "core/src/helpers/requests"
	"core/src/models"
	"core/src/repositories"
	validations "core/src/validations/auth"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterUser reads payload from request body then try to register a new user in database.
func RegisterUser(c *fiber.Ctx, db *mongo.Database) (interface{}, models.RequestError) {
	var payload models.User

	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		return nil, helpers.NewRequestError(err, 500)
	}

	usersRepository := repositories.NewUsersRepository(db)
	isEmailAlreadyInUse, emailInUseErr := validations.IsEmailInUse(usersRepository, payload.Email)

	if isEmailAlreadyInUse {
		return nil, emailInUseErr
	}

	isNickAlreadyInUse, nickInUseErr := validations.IsNickInUse(usersRepository, payload.Nick)

	if isNickAlreadyInUse {
		return nil, nickInUseErr
	}

	authRepository := repositories.NewAuthRepository(db)

	res, err := authRepository.RegisterUser(payload)

	if err != nil {
		return nil, helpers.NewRequestError(err, 500)
	}

	return res, models.RequestError{}
}
