package jwt

import (
	"core/internal/configs"
	"core/internal/entities"
	pkg "core/pkg/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// CreateUserToken creates an user JWT token with followed fields:
// id, name, email, exp. It returns string and error.
func CreateUserToken(u *entities.User, expiresIn int64, secret string) (string, error) {
	claims := jwt.MapClaims{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
		"exp":   expiresIn,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

// CreateAccessToken creates an user JWT token with followed fields:
// id, name, email, exp. It returns string and error.
func CreateAccessToken(u *entities.User, cfg *configs.Conf) (string, error) {
	expiresIn := time.Now().Add(time.Hour + 24).Unix()

	return CreateUserToken(u, expiresIn, cfg.JWTSecret)
}

// CreateRefreshToken creates an user JWT token with followed fields:
// id, name, email, exp. It returns string and error.
func CreateRefreshToken(u *entities.User, cfg *configs.Conf) (string, error) {
	expiresIn := time.Now().Add(time.Hour + 720).Unix()

	return CreateUserToken(u, expiresIn, cfg.JWTSecret)
}

// DecodeUserToken decodes an user JWT token with followed fields:
// id, name and email.
func DecodeUserToken(c *fiber.Ctx) entities.User {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	claimedUser := claims["user"].(map[string]interface{})

	parsedId, _ := pkg.ParseID(claimedUser["id"].(string))

	u := entities.User{
		ID:    parsedId,
		Name:  claimedUser["name"].(string),
		Email: claimedUser["email"].(string),
	}

	return u
}

// GetUserByToken decodes a token and get user info from token.
func GetUserByToken(c *fiber.Ctx) entities.User {
	return DecodeUserToken(c)
}
