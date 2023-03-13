package auth

import (
	"regexp"
	"strings"
	"time"

	"github.com/quessapp/toolkit/regexes"
	"github.com/quessapp/toolkit/validations"

	"core/pkg/errors"

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
// It removes special characters from nick and trim email.
func (d *SignUpUserDTO) Format() {
	d.Nick = regexp.MustCompile(regexes.SPECIAL_CHARS).ReplaceAllString(d.Nick, "")
	d.Nick = strings.TrimSpace(strings.ToLower(d.Nick))
	d.Email = strings.TrimSpace(d.Email)
}

// Validate is a method of SignUpUserDTO that validates the fields of the struct.
// The method uses the validation package to validate the Nick, Password, Name, Email and Locale fields.
// The Nick, Password, Name and Email fields are required and must have a length between 3 and 50 characters for the Nick, Name and Password fields
// and between 5 and 200 characters for the Email field.
// The Email field must also match a valid email format using the is.Email method.
// The Locale field is required and must match the regular expression defined in the regexes package for valid locales.
// The method then returns the validation error, if any, using the validations.GetValidationError method.
// If there are no validation errors, the method returns nil.
func (d SignUpUserDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Nick, validation.Required.Error(errors.NICK_FIELD_REQUIRED), validation.Length(3, 50).Error(errors.NICK_FIELD_LENGTH)),
		validation.Field(&d.Password, validation.Required.Error(errors.PASSWORD_FIELD_REQUIRED), validation.Length(6, 200).Error(errors.PASSWORD_FIELD_LENGTH)),
		validation.Field(&d.Name, validation.Required.Error(errors.NAME_FIELD_REQUIRED), validation.Length(3, 50).Error(errors.NAME_FIELD_LENGTH)),
		validation.Field(&d.Email, validation.Required.Error(errors.EMAIL_FIELD_REQUIRED), validation.Length(5, 200).Error(errors.EMAIL_FIELD_LENGTH), is.Email.Error(errors.EMAIL_FORMAT_INVALID)),
		validation.Field(&d.Locale, validation.Required.Error(errors.LOCALE_FIELD_REQUIRED), validation.Match(regexp.MustCompile(regexes.LOCALES)).Error(errors.LOCALE_FIELD_INVALID)),
	)

	return validations.GetValidationError(validationResult)
}

// Validate is a method of SignInUserDTO that validates the fields of the struct.
// The method uses the validation package to validate the Nick and Password fields.
// Both fields are required and must have a length between 3 and 50 characters for the Nick field and between 6 and 200 characters for the Password field.
// The method then returns the validation error, if any, using the validations.GetValidationError method.
// If there are no validation errors, the method returns nil.
func (d SignInUserDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Nick, validation.Required.Error(errors.NICK_FIELD_REQUIRED), validation.Length(3, 50).Error(errors.NICK_FIELD_LENGTH)),
		validation.Field(&d.Password, validation.Required.Error(errors.PASSWORD_FIELD_REQUIRED), validation.Length(6, 200).Error(errors.PASSWORD_FIELD_LENGTH)),
	)

	return validations.GetValidationError(validationResult)
}
