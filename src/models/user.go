package models

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is a model for each user in app.
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Nick      string             `json:"nick,omitempty"`
	Name      string             `json:"name,omitempty"`
	AvatarUrl string             `json:"avatarUrl,omitempty"`

	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`

	EnableAppEmails bool `json:"enableAppEmails,omitempty"`
	IsShadowBanned  bool `json:"isShadowBanned,omitempty"`

	PostsLimit     int    `json:"postsLimit,omitempty"`
	CustomerId     string `json:"customerId,omitempty"`
	IsPro          bool   `json:"isPro,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty"`

	ProExpiresAt  time.Time `json:"proExpiresAt,omitempty"`
	LastPublishAt time.Time `json:"lastPublishAt,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
}

func (u User) NormalizeNick() string {
	return strings.ToLower(strings.TrimSpace(u.Nick))
}

// BlockedUser is a model for each blocked user in app.
type BlockedUser struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserToBlock User               `json:"userToBlock"`
	BlockedBy   User               `json:"blockedBy"`
}
