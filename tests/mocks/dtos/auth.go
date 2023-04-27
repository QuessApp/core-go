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
			},
		},
		{
			OnRun: func() {
				signUpData.Nick = "_.test-user"
				signUpData.Format()
				assert.Equal(t, "testuser", signUpData.Nick)
			},
		},
		{
			OnRun: func() {
				signUpData.Nick = "-test.user@"
				signUpData.Format()
				assert.Equal(t, "testuser", signUpData.Nick)
			},
		},
		{
			OnRun: func() {
				signUpData.Nick = "test.user"
				signUpData.Format()
				assert.Equal(t, "testuser", signUpData.Nick)
			},
		},
		{
			OnRun: func() {
				signUpData.Nick = "FOOBAR-USER"
				signUpData.Format()
				assert.Equal(t, "foobaruser", signUpData.Nick)
			},
		},
		{
			OnRun: func() {
				signUpData.Nick = "FOOBAR_USER"
				signUpData.Format()
				assert.Equal(t, "foobaruser", signUpData.Nick)
			},
		},
		{
			OnRun: func() {
				signUpData.Nick = "FOOBAR.USER"
				signUpData.Format()
				assert.Equal(t, "foobaruser", signUpData.Nick)
			},
		},
		{
			OnRun: func() {
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
				assert.EqualError(t, signUpData.Validate(), "nick_field_required.")
			},
		},
		{
			OnRun: func() {
				signUpData.Nick = "_.test-user"
				assert.NoError(t, signUpData.Validate())
			},
		},
		{
			OnRun: func() {
				signUpData.Nick = tests.GenerateRandomString(51)
				assert.EqualError(t, signUpData.Validate(), "nick_field_length.")
			},
		},
		{
			OnRun: func() {
				signUpData.Nick = "foobar"
				signUpData.Password = ""
				assert.EqualError(t, signUpData.Validate(), "password_field_required.")
			},
		},
		{
			OnRun: func() {
				signUpData.Password = tests.GenerateRandomString(10)
				assert.NoError(t, signUpData.Validate())
			},
		},
		{
			OnRun: func() {
				signUpData.Password = tests.GenerateRandomString(300)
				assert.EqualError(t, signUpData.Validate(), "password_field_length.")
			},
		},
		{
			OnRun: func() {
				signUpData.Email = ""
				assert.EqualError(t, signUpData.Validate(), "email_field_required")
			},
		},
		{
			OnRun: func() {
				signUpData.Email = tests.GenerateRandomString(130)
				assert.EqualError(t, signUpData.Validate(), "email_format_invalid")
			},
		},
		{
			OnRun: func() {
				e := fmt.Sprintf("%s@%s.com", tests.GenerateRandomString(130), tests.GenerateRandomString(130))
				signUpData.Email = e
				assert.EqualError(t, signUpData.Validate(), "email_field_length")
			},
		},
		{
			OnRun: func() {
				signUpData.Email = "test-api@example.com"
				signUpData.Password = tests.GenerateRandomString(130)
				assert.NoError(t, signUpData.Validate())
			},
		},
		{
			OnRun: func() {
				signUpData.Locale = ""
				assert.EqualError(t, signUpData.Validate(), "locale_field_required.")
			},
		},
		{
			OnRun: func() {
				signUpData.Locale = "foobar"
				assert.EqualError(t, signUpData.Validate(), "locale_field_invalid.")
			},
		},
		{
			OnRun: func() {
				signUpData.Locale = "es-ES"
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
				assert.EqualError(t, signInData.Validate(), "nick_field_required.")
			},
		},
		{
			OnRun: func() {
				signInData.Nick = "_.test-user"
				assert.NoError(t, signInData.Validate())
			},
		},
		{
			OnRun: func() {
				signInData.Nick = tests.GenerateRandomString(51)
				assert.EqualError(t, signInData.Validate(), "nick_field_length.")
			},
		},
		{
			OnRun: func() {
				signInData.Nick = "foobar"
				signInData.Password = ""
				assert.EqualError(t, signInData.Validate(), "password_field_required.")
			},
		},
		{
			OnRun: func() {
				signInData.Password = tests.GenerateRandomString(10)
				assert.NoError(t, signInData.Validate())
			},
		},
		{
			OnRun: func() {
				signInData.Password = tests.GenerateRandomString(300)
				assert.EqualError(t, signInData.Validate(), "password_field_length.")
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
				assert.EqualError(t, forgotPasswordData.Validate(), "email_field_required.")
			},
		},
		{
			OnRun: func() {
				forgotPasswordData.Email = tests.GenerateRandomString(130)
				assert.EqualError(t, forgotPasswordData.Validate(), "email_format_invalid.")
			},
		},
		{
			OnRun: func() {
				e := fmt.Sprintf("%s@%s.com", tests.GenerateRandomString(130), tests.GenerateRandomString(130))
				forgotPasswordData.Email = e
				assert.EqualError(t, forgotPasswordData.Validate(), "email_field_length.")
			},
		},
		{
			OnRun: func() {
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
				assert.EqualError(t, resetPasswordData.Validate(), "password_field_required.")
			},
		},
		{
			OnRun: func() {
				resetPasswordData.Password = tests.GenerateRandomString(10)
				assert.NoError(t, resetPasswordData.Validate())
			},
		},
		{
			OnRun: func() {
				resetPasswordData.Password = tests.GenerateRandomString(300)
				assert.EqualError(t, resetPasswordData.Validate(), "password_field_length.")
			},
		},
		{
			OnRun: func() {
				resetPasswordData.Code = ""
				assert.EqualError(t, resetPasswordData.Validate(), "code_required")
			},
		},
	}

	return forgotPasswordDataTest
}
