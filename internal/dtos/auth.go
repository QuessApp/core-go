package dtos

import (
	pkg "core/pkg/entities"
	"core/pkg/errors"
	"core/pkg/regexes"
	"core/pkg/validations"
	"regexp"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// SignUpUserDTO is DTO for payload for signup handler.
type SignUpUserDTO struct {
	ID        pkg.ID
	Email     string
	Password  string
	Nick      string
	Name      string
	CreatedAt time.Time
}

// SignInUserDTO is DTO for payload for signin handler.
type SignInUserDTO struct {
	ID       pkg.ID
	Nick     string
	Password string
}

// Format formats DTO information. It removes special characters from nick, trim email, etc.
func (d *SignUpUserDTO) Format() {
	d.Nick = regexp.MustCompile(regexes.SPECIAL_CHARS).ReplaceAllString(d.Nick, "")
	d.Nick = strings.TrimSpace(strings.ToLower(d.Nick))
	d.Email = strings.TrimSpace(d.Email)
}

// Validate validates passed struct then returns a string.
func (d SignUpUserDTO) Validate() error {
	validationResult := validation.ValidateStruct(&d,
		validation.Field(&d.Nick, validation.Required.Error(errors.NICK_FIELD_REQUIRED), validation.Length(3, 50).Error(errors.NICK_FIELD_LENGTH)),
		validation.Field(&d.Password, validation.Required.Error(errors.PASSWORD_FIELD_REQUIRED), validation.Length(6, 200).Error(errors.PASSWORD_FIELD_LENGTH)),
		validation.Field(&d.Name, validation.Required.Error(errors.NAME_FIELD_REQUIRED), validation.Length(3, 50).Error(errors.NAME_FIELD_LENGTH)),
		validation.Field(&d.Email, validation.Required.Error(errors.EMAIL_FIELD_REQUIRED), validation.Length(5, 200).Error(errors.EMAIL_FIELD_LENGTH), is.Email.Error(errors.EMAIL_FORMAT_INVALID)),
	)

	return validations.GetValidationError(validationResult)
}
