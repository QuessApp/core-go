package jwt

import (
	"core/internal/configs"
	"core/internal/entities"
	pkg "core/pkg/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const (
	THIRTY_DAYS_IN_HOURS = time.Hour * 720
	ONE_DAY_IN_HOURS     = time.Hour + 24
)

var (
	ACCESS_TOKEN_EXPIRES_IN  = time.Now().Add(ONE_DAY_IN_HOURS).Unix()
	REFRESH_TOKEN_EXPIRES_IN = time.Now().Add(THIRTY_DAYS_IN_HOURS).Unix()
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
	return CreateUserToken(u, ACCESS_TOKEN_EXPIRES_IN, cfg.JWTSecret)
}

// CreateRefreshToken creates an user JWT token with followed fields:
// id, name, email, exp. It returns string and error.
func CreateRefreshToken(u *entities.User, cfg *configs.Conf) (string, error) {
	return CreateUserToken(u, REFRESH_TOKEN_EXPIRES_IN, cfg.JWTSecret)
}

// DecodeUserToken decodes an user JWT token with followed fields:
// id, name and email.
func DecodeUserToken(c *fiber.Ctx) pkg.DecodeUserTokenResult {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	parsedId, _ := pkg.ParseID(claims["id"].(string))

	u := pkg.DecodeUserTokenResult{
		ID:    parsedId,
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
	}

	return u
}

// GetUserByToken decodes a token and get user info from token.
func GetUserByToken(c *fiber.Ctx) pkg.DecodeUserTokenResult {
	return DecodeUserToken(c)
}
