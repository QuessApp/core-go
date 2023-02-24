package entities

import (
	"regexp"
	"strings"

	"core/pkg/errors"
	"core/pkg/validations"

	pkg "core/pkg/entities"
	"core/pkg/regexes"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// BlockedUser is a model for each blocked user in app.
type BlockedUser struct {
	ID          pkg.ID `json:"id" bson:"_id" `
	UserToBlock pkg.ID `json:"userToBlock" bson:"userToBlock"`
	BlockedBy   pkg.ID `json:"blockedBy" bson:"blockedBy"`
}

// User is a model for each user in app.
type User struct {
	ID        pkg.ID `json:"id" bson:"_id"`
	Nick      string `json:"nick,omitempty"`
	Name      string `json:"name,omitempty"`
	AvatarURL string `json:"avatarUrl" bson:"avatarUrl"`

	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`

	// EnableAppEmails is a bool value to verify if user can receive emails (received questions, etc.)
	EnableAppEmails bool `json:"enableAppEmails,omitempty" bson:"enableAppEmails"`
	IsShadowBanned  bool `json:"isShadowBanned,omitempty" bson:"isShadowBanned"`
	PostsLimit      int  `json:"postsLimit,omitempty" bson:"postsLimit"`
	// CustomerId of Stripe. Type must be String or nil.
	CustomerID any  `json:"customerId,omitempty" bson:"customerId"`
	IsPRO      bool `json:"isPro,omitempty" bson:"isPro"`
	// SubscriptionID of Stripe. Type must be String or nil.
	SubscriptionID any `json:"subscriptionId,omitempty" bson:"subscriptionId"`
	// ProExpiresAt of Stripe. Type must be Time.time or nil.
	ProExpiresAt any `json:"proExpiresAt,omitempty" bson:"proExpiresAt"`
	// LastPublishAt is the last published post of user. Type must be Time.time or nil.
	LastPublishAt any `json:"lastPublishAt,omitempty" bson:"lastPublishAt"`
	// CreatedAt is the date that user is created. Type must be Time.time or nil.
	CreatedAt any `json:"createdAt,omitempty" bson:"createdAt"`
}

// Format formats user information. It removes special characters from nick, trim email, etc.
func (u *User) Format() {
	u.Nick = regexp.MustCompile(regexes.SPECIAL_CHARS).ReplaceAllString(u.Nick, "")
	u.Nick = strings.TrimSpace(strings.ToLower(u.Nick))
	u.Email = strings.TrimSpace(u.Email)
}

// Validate validates passed struct then returns a string.
func (u User) Validate() error {
	validationResult := validation.ValidateStruct(&u,
		validation.Field(&u.Nick, validation.Required.Error(errors.NICK_FIELD_REQUIRED), validation.Length(3, 50).Error(errors.NICK_FIELD_LENGTH)),
		validation.Field(&u.Password, validation.Required.Error(errors.PASSWORD_FIELD_REQUIRED), validation.Length(6, 200).Error(errors.PASSWORD_FIELD_LENGTH)),
		validation.Field(&u.Name, validation.Required.Error(errors.NAME_FIELD_REQUIRED), validation.Length(3, 50).Error(errors.NAME_FIELD_LENGTH)),
		validation.Field(&u.Email, validation.Required.Error(errors.EMAIL_FIELD_REQUIRED), validation.Length(5, 200).Error(errors.EMAIL_FIELD_LENGTH), is.Email.Error(errors.EMAIL_FORMAT_INVALID)),
	)

	return validations.GetValidationError(validationResult)
}

// GetBasicInfos gets basic data of an user like id, name, nick, etc.
func (u User) GetBasicInfos() *User {
	return &User{
		ID:        u.ID,
		Name:      u.Name,
		Nick:      u.Nick,
		AvatarURL: u.AvatarURL,
	}
}
