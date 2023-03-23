package auth

import (
	"time"

	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// Token is a struct that represents an authentication token.
// It can be a refresh token, an access token, or a code.
type Token struct {
	ID toolkitEntities.ID `json:"id" bson:"_id"`
	// Refresh or Code
	Type string `json:"type" bson:"type"`

	ExpiresAt time.Time           `json:"expiresAt" bson:"expiresAt"`
	CreatedAt time.Time           `json:"createdAt" bson:"createdAt"`
	CreatedBy *toolkitEntities.ID `json:"createdBy" bson:"createdBy"`

	AccessToken  string `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	Code         string `json:"code,omitempty" bson:"code,omitempty"`
}
