package dtos

import (
	"testing"

	"github.com/quessapp/core-go/internal/auth"
	"github.com/quessapp/core-go/internal/blocks"
	"github.com/quessapp/core-go/pkg/tests"
	mocksDTOs "github.com/quessapp/core-go/tests/mocks/dtos"
)

var mockedSignUpPayload = auth.SignUpUserDTO{
	Email:    "api@example.com",
	Password: "test123",
	Nick:     "foobar",
	Name:     "example",
	Locale:   "en-US",
}

func TestFormat(t *testing.T) {
	signUpFormatDTOMock := mocksDTOs.GetSignUpFormatDTOMock(t, mockedSignUpPayload)
	tests.RunBatchTests(signUpFormatDTOMock)
}

func TestValidate(t *testing.T) {
	signInValidateDTOMock := mocksDTOs.GetSignInValidateDTOMock(t, auth.SignInUserDTO{
		Nick:     "foobar",
		Password: "test123",
		TrustIP:  true,
	})
	tests.RunBatchTests(signInValidateDTOMock)

	signUpValidateDTOMock := mocksDTOs.GetSignUpValidateDTOMock(t, mockedSignUpPayload)
	tests.RunBatchTests(signUpValidateDTOMock)

	forgotPasswordValidateDTOMock := mocksDTOs.GetFormatPasswordValidateDTOMock(t, auth.ForgotPasswordDTO{
		Email: "text-api@example.com",
	})
	tests.RunBatchTests(forgotPasswordValidateDTOMock)

	resetPasswordValidateDTOMock := mocksDTOs.GetResetPasswordValidateDTOMock(t, auth.ResetPasswordDTO{
		Code:                 "1bc23",
		Password:             "123test",
		LogoutFromAllDevices: true,
	})
	tests.RunBatchTests(resetPasswordValidateDTOMock)

	blockUserValidateDTO := mocksDTOs.GetBlockUserValidateDTOMock(t, blocks.BlockUserDTO{})
	tests.RunBatchTests(blockUserValidateDTO)
}
