package users

import (
	"core/pkg/errors"
	"regexp"

	"github.com/quessapp/toolkit/regexes"

	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/quessapp/toolkit/validations"

	validation "github.com/go-ozzo/ozzo-validation"
)

// UpdateProfileDTO is DTO for payload for update user profile handler.
type UpdateProfileDTO struct {
	Nick   string `json:"nick,omitempty" bson:"nick"`
	Name   string `json:"name,omitempty" bson:"name"`
	Locale string `json:"locale,omitempty" bson:"locale"`
	Email  string `json:"email,omitempty" bson:"email"`
}

// UpdatePreferencesDTO is DTO for payload for update preferences handler.
type UpdatePreferencesDTO struct {
	EnanbleAPPPushNotifications bool `json:"enableAppPushNotifications" bson:"enableAppPushNotifications"`
	EnableAPPEmails             bool `json:"enableAppEmails" bson:"enableAppEmails"`
}

// Validate validates passed struct then returns a string.
//
// It validates if nick, password, name, email and locale are valid.
func (d UpdateProfileDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Nick, validation.Required.Error(errors.NICK_FIELD_REQUIRED), validation.Length(3, 50).Error(errors.NICK_FIELD_LENGTH)),
		validation.Field(&d.Name, validation.Required.Error(errors.NAME_FIELD_REQUIRED), validation.Length(3, 50).Error(errors.NAME_FIELD_LENGTH)),
		validation.Field(&d.Email, validation.Required.Error(errors.EMAIL_FIELD_REQUIRED), validation.Length(5, 200).Error(errors.EMAIL_FIELD_LENGTH), is.Email.Error(errors.EMAIL_FORMAT_INVALID)),
		validation.Field(&d.Locale, validation.Required.Error(errors.LOCALE_FIELD_REQUIRED), validation.Match(regexp.MustCompile(regexes.LOCALES)).Error(errors.LOCALE_FIELD_INVALID)),
	)

	return validations.GetValidationError(validationResult)
}
