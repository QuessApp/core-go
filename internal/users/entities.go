package users

import (
	"regexp"
	"strings"
	"time"

	regexes "github.com/quessapp/toolkit/regexes"

	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// BlockedUser is a model for each blocked user in app.
type BlockedUser struct {
	ID          toolkitEntities.ID `json:"id" bson:"_id" `
	UserToBlock toolkitEntities.ID `json:"userToBlock" bson:"userToBlock"`
	BlockedBy   toolkitEntities.ID `json:"blockedBy" bson:"blockedBy"`
}

// UpdatePreferencesDTO is DTO for payload for update preferences handler.
type UpdatePreferencesDTO struct {
	EnanbleAPPPushNotifications bool `json:"enableAppPushNotifications" bson:"enableAppPushNotifications"`
	EnableAPPEmails             bool `json:"enableAppEmails" bson:"enableAppEmails"`
}

// User is a model for each user in app.
type User struct {
	ID        toolkitEntities.ID `json:"id" bson:"_id"`
	Nick      string             `json:"nick,omitempty"`
	Name      string             `json:"name,omitempty"`
	AvatarURL string             `json:"avatarUrl" bson:"avatarUrl"`

	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`

	// EnanbleAPPPushNotifications is a bool value to verify if user can push notifications (received questions, etc.)
	EnanbleAPPPushNotifications bool `json:"enableAppPushNotifications,omitempty" bson:"enableAppPushNotifications"`
	// EnableAPPEmails is a bool value to verify if user can receive emails (received questions, etc.)
	EnableAPPEmails bool `json:"enableAppEmails,omitempty" bson:"enableAppEmails"`
	IsShadowBanned  bool `json:"isShadowBanned,omitempty" bson:"isShadowBanned"`
	PostsLimit      int  `json:"postsLimit,omitempty" bson:"postsLimit"`
	// CustomerId of Stripe. Type must be String or nil.
	CustomerID *string `json:"customerId,omitempty" bson:"customerId"`
	IsPRO      bool    `json:"isPro,omitempty" bson:"isPro"`
	// SubscriptionID of Stripe. Type must be String or nil.
	SubscriptionID *string `json:"subscriptionId,omitempty" bson:"subscriptionId"`
	// ProExpiresAt of Stripe. Type must be Time.time or nil.
	ProExpiresAt *string `json:"proExpiresAt,omitempty" bson:"proExpiresAt"`
	// LastPublishAt is the last published post of user. Type must be Time.time or nil.
	LastPublishAt *time.Time `json:"lastPublishAt,omitempty" bson:"lastPublishAt"`
	// CreatedAt is the date that user is created. Type must be Time.time or nil.
	CreatedAt *time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	Locale    string     `json:"locale,omitempty" bson:"locale"`
}

// ResponseWithUser is a model to use with Response model.
// It can be returned like: { ..., message: null, data: { user: { ... } }}
type ResponseWithUser struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// PaginatedUsers is a model for paginated users in app.
type PaginatedUsers struct {
	Users      *[]User `json:"users"`
	TotalCount int64   `json:"totalCount"`
}

// Format formats user information. It removes special characters from nick, trim email, etc.
func (u *User) Format() {
	u.Nick = regexp.MustCompile(regexes.SPECIAL_CHARS).ReplaceAllString(u.Nick, "")
	u.Nick = strings.TrimSpace(strings.ToLower(u.Nick))
	u.Email = strings.TrimSpace(u.Email)
}

// GetBasicInfos gets user basic infos. It hide the sensible data like password, email, etc.
// It is a method for the "user" struct. It returns an pointer to user.
func (u User) GetBasicInfos() *User {
	return &User{
		ID:        u.ID,
		Name:      u.Name,
		Nick:      u.Nick,
		AvatarURL: u.AvatarURL,
	}
}
