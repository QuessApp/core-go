package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is a model for each user in app.
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Nick      string             `json:"nick"`
	Name      string             `json:"name"`
	AvatarUrl string             `json:"avatarUrl"`

	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`

	EnableAppEmails bool `json:"enableAppEmails,omitempty"`
	IsShadowBanned  bool `json:"isShadowBanned,omitempty"`

	PostsLimit     int    `json:"postsLimit"`
	CustomerId     string `json:"customerId,omitempty"`
	IsPro          bool   `json:"isPro,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty"`

	ProExpiresAt  time.Time `json:"proExpiresAt,omitempty"`
	LastPublishAt time.Time `json:"lastPublishAt"`
	CreatedAt     time.Time `json:"createdAt"`
}

// BlockedUser is a model for each blocked user in app.
type BlockedUser struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserToBlock User               `json:"userToBlock"`
	BlockedBy   User               `json:"blockedBy"`
}
