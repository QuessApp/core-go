package auth

import (
	"time"

	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// Token is a struct that represents an authentication token.
// It can be a refresh token, an access token, or a code.
type Token struct {
	ID toolkitEntities.ID `json:"id" bson:"_id"`
	// Bearer or Code.
	// Bearer is for access and refresh tokens, and code may be for email verification.
	Type string `json:"type" bson:"type"`

	ExpiresAt time.Time           `json:"expiresAt" bson:"expiresAt"`
	CreatedAt time.Time           `json:"createdAt" bson:"createdAt"`
	CreatedBy *toolkitEntities.ID `json:"createdBy" bson:"createdBy"`

	AccessToken  string `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	// It can be a code because it can be used for email verification like reset password.
	Code string `json:"code,omitempty" bson:"code,omitempty"`
}
