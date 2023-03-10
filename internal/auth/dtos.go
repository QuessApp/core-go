package auth

import (
	"regexp"
	"strings"
	"time"

	"github.com/quessapp/toolkit/regexes"
	"github.com/quessapp/toolkit/validations"

	toolkitEntities "github.com/quessapp/toolkit/entities"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// SignUpUserDTO is DTO for payload for signup handler.
type SignUpUserDTO struct {
	ID        toolkitEntities.ID
	Email     string
	Password  string
	Nick      string
	Name      string
	CreatedAt time.Time
	Locale    string
}

// SignInUserDTO is DTO for payload for signin handler.
type SignInUserDTO struct {
	ID       toolkitEntities.ID
	Nick     string
	Password string
}

// Format formats DTO information.
//
// It removes special characters from nick and trim email.
func (d *SignUpUserDTO) Format() {
	d.Nick = regexp.MustCompile(regexes.SPECIAL_CHARS).ReplaceAllString(d.Nick, "")
	d.Nick = strings.TrimSpace(strings.ToLower(d.Nick))
	d.Email = strings.TrimSpace(d.Email)
}

// Validate validates passed struct then returns a string.
//
// It validates if nick, password, name, email and locale are valid.
func (d SignUpUserDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Nick, validation.Required.Error(NICK_FIELD_REQUIRED), validation.Length(3, 50).Error(NICK_FIELD_LENGTH)),
		validation.Field(&d.Password, validation.Required.Error(PASSWORD_FIELD_REQUIRED), validation.Length(6, 200).Error(PASSWORD_FIELD_LENGTH)),
		validation.Field(&d.Name, validation.Required.Error(NAME_FIELD_REQUIRED), validation.Length(3, 50).Error(NAME_FIELD_LENGTH)),
		validation.Field(&d.Email, validation.Required.Error(EMAIL_FIELD_REQUIRED), validation.Length(5, 200).Error(EMAIL_FIELD_LENGTH), is.Email.Error(EMAIL_FORMAT_INVALID)),
		validation.Field(&d.Locale, validation.Required.Error(LOCALE_FIELD_REQUIRED), validation.Match(regexp.MustCompile(regexes.LOCALES)).Error(LOCALE_FIELD_INVALID)),
	)

	return validations.GetValidationError(validationResult)
}

// Validate validates passed struct then returns a string.
//
// It validates if nick and password are valid.
func (d SignInUserDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Nick, validation.Required.Error(NICK_FIELD_REQUIRED), validation.Length(3, 50).Error(NICK_FIELD_LENGTH)),
		validation.Field(&d.Password, validation.Required.Error(PASSWORD_FIELD_REQUIRED), validation.Length(6, 200).Error(PASSWORD_FIELD_LENGTH)),
	)

	return validations.GetValidationError(validationResult)
}
