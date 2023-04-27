package mocks

import (
	"fmt"
	"testing"

	"github.com/quessapp/core-go/internal/auth"
	"github.com/quessapp/core-go/pkg/tests"
	"github.com/stretchr/testify/assert"
)

// GetSignUpFormatDTOMock returns a slice of BatchTest for SignUpUserDTO.
func GetSignUpFormatDTOMock(t *testing.T, signUpData auth.SignUpUserDTO) []tests.BatchTest {
	signUpDataTest := []tests.BatchTest{
		{
			OnRun: func() {
				signUpData.Nick = "@ adsa32321@0-9"
				signUpData.Format()
				assert.Equal(t, "adsa3232109", signUpData.Nick)

				signUpData.Nick = "_.test-user"
				signUpData.Format()
				assert.Equal(t, "testuser", signUpData.Nick)

				signUpData.Nick = "-test.user@"
				signUpData.Format()
				assert.Equal(t, "testuser", signUpData.Nick)

				signUpData.Nick = "test.user"
				signUpData.Format()
				assert.Equal(t, "testuser", signUpData.Nick)

				signUpData.Nick = "FOOBAR-USER"
				signUpData.Format()
				assert.Equal(t, "foobaruser", signUpData.Nick)

				signUpData.Nick = "FOOBAR_USER"
				signUpData.Format()
				assert.Equal(t, "foobaruser", signUpData.Nick)

				signUpData.Nick = "FOOBAR.USER"
				signUpData.Format()
				assert.Equal(t, "foobaruser", signUpData.Nick)

				signUpData.Nick = "foOBar@a-uSrR"
				signUpData.Format()
				assert.Equal(t, "foobarausrr", signUpData.Nick)
			},
		},
	}

	return signUpDataTest
}

// GetSignUpFormatDTOMock returns a slice of BatchTest for SignUpUserDTO.
func GetSignUpValidateDTOMock(t *testing.T, signUpData auth.SignUpUserDTO) []tests.BatchTest {
	signUpDataTest := []tests.BatchTest{
		{
			OnRun: func() {
				signUpData.Nick = ""
				assert.ErrorContains(t, signUpData.Validate(), "nick_field_required")

				signUpData.Nick = "_.test-user"
				assert.NoError(t, signUpData.Validate())

				signUpData.Nick = tests.GenerateRandomString(51)
				assert.ErrorContains(t, signUpData.Validate(), "nick_field_length")

				signUpData.Nick = "foobar"
				assert.NoError(t, signUpData.Validate())
			},
		},
		{
			OnRun: func() {
				signUpData.Password = ""
				assert.ErrorContains(t, signUpData.Validate(), "password_field_required")

				signUpData.Password = tests.GenerateRandomString(300)
				assert.ErrorContains(t, signUpData.Validate(), "password_field_length")

				signUpData.Password = tests.GenerateRandomString(10)
				assert.NoError(t, signUpData.Validate())
			},
		},
		{
			OnRun: func() {
				signUpData.Email = ""
				assert.ErrorContains(t, signUpData.Validate(), "email_field_required")

				signUpData.Email = tests.GenerateRandomString(130)
				assert.ErrorContains(t, signUpData.Validate(), "email_format_invalid")

				e := fmt.Sprintf("%s@%s.com", tests.GenerateRandomString(130), tests.GenerateRandomString(130))
				signUpData.Email = e
				assert.ErrorContains(t, signUpData.Validate(), "email_field_length")

				signUpData.Email = "test-api@example.com"
				assert.NoError(t, signUpData.Validate())
			},
		},
		{
			OnRun: func() {
				signUpData.Locale = "foobar"
				assert.ErrorContains(t, signUpData.Validate(), "locale_field_invalid")

				signUpData.Locale = ""
				assert.ErrorContains(t, signUpData.Validate(), "locale_field_required")

				signUpData.Locale = "es-ES"
				assert.NoError(t, signUpData.Validate())

				signUpData.Locale = "pt-ES"
				assert.ErrorContains(t, signUpData.Validate(), "locale_field_invalid")

				signUpData.Locale = "en-ES"
				assert.ErrorContains(t, signUpData.Validate(), "locale_field_invalid")

				signUpData.Locale = "pt-US"
				assert.ErrorContains(t, signUpData.Validate(), "locale_field_invalid")

				signUpData.Locale = "pt-BR"
				assert.NoError(t, signUpData.Validate())

				signUpData.Locale = "en-US"
				assert.NoError(t, signUpData.Validate())
			},
		},
	}

	return signUpDataTest
}

// GetSignInValidateDTOMock returns a slice of BatchTest for SignInUserDTO.
func GetSignInValidateDTOMock(t *testing.T, signInData auth.SignInUserDTO) []tests.BatchTest {
	signInDataTest := []tests.BatchTest{
		{
			OnRun: func() {
				signInData.Nick = ""
				assert.ErrorContains(t, signInData.Validate(), "nick_field_required")

				signInData.Nick = tests.GenerateRandomString(51)
				assert.ErrorContains(t, signInData.Validate(), "nick_field_length")

				signInData.Nick = "_.test-user"
				assert.NoError(t, signInData.Validate())
			},
		},
		{
			OnRun: func() {
				signInData.Password = ""
				assert.ErrorContains(t, signInData.Validate(), "password_field_required")

				signInData.Password = tests.GenerateRandomString(300)
				assert.ErrorContains(t, signInData.Validate(), "password_field_length")

				signInData.Password = tests.GenerateRandomString(10)
				assert.NoError(t, signInData.Validate())
			},
		},
	}

	return signInDataTest
}

// GetFormatPasswordValidateDTOMock returns a slice of BatchTest for ForgotPasswordDTO.
func GetFormatPasswordValidateDTOMock(t *testing.T, forgotPasswordData auth.ForgotPasswordDTO) []tests.BatchTest {
	forgotPasswordDataTest := []tests.BatchTest{
		{
			OnRun: func() {
				forgotPasswordData.Email = ""
				assert.ErrorContains(t, forgotPasswordData.Validate(), "email_field_required")

				forgotPasswordData.Email = tests.GenerateRandomString(130)
				assert.ErrorContains(t, forgotPasswordData.Validate(), "email_format_invalid")

				e := fmt.Sprintf("%s@%s.com", tests.GenerateRandomString(130), tests.GenerateRandomString(130))
				forgotPasswordData.Email = e
				assert.ErrorContains(t, forgotPasswordData.Validate(), "email_field_length")

				forgotPasswordData.Email = "test-api@example.com"
				assert.NoError(t, forgotPasswordData.Validate())
			},
		},
	}

	return forgotPasswordDataTest
}

// GetResetPasswordValidateDTOMock returns a slice of BatchTest for ResetPasswordDTO.
func GetResetPasswordValidateDTOMock(t *testing.T, resetPasswordData auth.ResetPasswordDTO) []tests.BatchTest {
	forgotPasswordDataTest := []tests.BatchTest{
		{
			OnRun: func() {
				resetPasswordData.Password = ""
				assert.ErrorContains(t, resetPasswordData.Validate(), "password_field_required")

				resetPasswordData.Password = tests.GenerateRandomString(300)
				assert.ErrorContains(t, resetPasswordData.Validate(), "password_field_length")

				resetPasswordData.Password = tests.GenerateRandomString(10)
				assert.NoError(t, resetPasswordData.Validate())
			},
		},
		{
			OnRun: func() {
				resetPasswordData.Code = ""
				assert.ErrorContains(t, resetPasswordData.Validate(), "code_required")

				resetPasswordData.Code = tests.GenerateRandomString(20)
				assert.NoError(t, resetPasswordData.Validate())
			},
		},
	}

	return forgotPasswordDataTest
}
