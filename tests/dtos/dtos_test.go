package dtos

import (
	"testing"

	"github.com/quessapp/core-go/internal/auth"
	"github.com/quessapp/core-go/internal/blocks"
	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/internal/reports"
	"github.com/quessapp/core-go/internal/users"
	"github.com/quessapp/core-go/pkg/tests"
)

var mockedSignUpPayload = auth.SignUpUserDTO{
	Email:    "api@example.com",
	Password: "test123",
	Nick:     "foobar",
	Name:     "example",
	Locale:   "en-US",
}

func TestFormat(t *testing.T) {
	signUpFormatDTOBatches := GetSignUpFormatDTOBatches(t, mockedSignUpPayload)
	tests.RunBatchTests(signUpFormatDTOBatches)
}

func TestValidate(t *testing.T) {
	signInValidateDTOBatches := GetSignInValidateDTOBatches(t, auth.SignInUserDTO{
		Nick:     "foobar",
		Password: "test123",
		TrustIP:  true,
	})
	tests.RunBatchTests(signInValidateDTOBatches)

	signUpValidateDTOBatches := GetSignUpValidateDTOBatches(t, mockedSignUpPayload)
	tests.RunBatchTests(signUpValidateDTOBatches)

	forgotPasswordValidateDTOBatches := GetFormatPasswordValidateDTOBatches(t, auth.ForgotPasswordDTO{
		Email: "text-api@example.com",
	})
	tests.RunBatchTests(forgotPasswordValidateDTOBatches)

	resetPasswordValidateDTOBatches := GetResetPasswordValidateDTOBatches(t, auth.ResetPasswordDTO{
		Code:                 "1bc23",
		Password:             "123test",
		LogoutFromAllDevices: true,
	})
	tests.RunBatchTests(resetPasswordValidateDTOBatches)

	blockUserValidateDTO := GetBlockUserValidateDTOBatches(t, blocks.BlockUserDTO{})
	tests.RunBatchTests(blockUserValidateDTO)

	replyQuestionValidateDTOBatches := GetReplyQuestionValidateDTOBatches(t, questions.ReplyQuestionDTO{})
	tests.RunBatchTests(replyQuestionValidateDTOBatches)

	editReplyQuestionValidateDTOBatches := GetEditReplyQuestionValidateDTOBatches(t, questions.EditQuestionReplyDTO{})
	tests.RunBatchTests(editReplyQuestionValidateDTOBatches)

	createReplyQuestionValidateDTOBatches := GetCreateQuestionValidateDTOBatches(t, questions.CreateQuestionDTO{})
	tests.RunBatchTests(createReplyQuestionValidateDTOBatches)

	createReportValidateDTOBatches := GetCreateReportValidateDTOBatches(t, reports.CreateReportDTO{
		Reason: "spam",
		Type:   "question",
	})
	tests.RunBatchTests(createReportValidateDTOBatches)

	updateUserProfileValidateDTOBatches := GetUpdateUserProfileValidateDTOBatches(t, users.UpdateProfileDTO{
		Email:  "api@example.com",
		Nick:   "foobar",
		Name:   "example",
		Locale: "en-US",
	})
	tests.RunBatchTests(updateUserProfileValidateDTOBatches)
}
