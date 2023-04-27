package batches

import (
	"testing"

	"github.com/quessapp/core-go/internal/auth"
	"github.com/quessapp/core-go/internal/blocks"
	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/internal/reports"
	"github.com/quessapp/core-go/internal/users"
	"github.com/quessapp/core-go/pkg/tests"
	DTOsBatches "github.com/quessapp/core-go/tests/batches/dtos"
)

var mockedSignUpPayload = auth.SignUpUserDTO{
	Email:    "api@example.com",
	Password: "test123",
	Nick:     "foobar",
	Name:     "example",
	Locale:   "en-US",
}

func TestFormat(t *testing.T) {
	signUpFormatDTOBatches := DTOsBatches.GetSignUpFormatDTOBatches(t, mockedSignUpPayload)
	tests.RunBatchTests(signUpFormatDTOBatches)
}

func TestValidate(t *testing.T) {
	signInValidateDTOBatches := DTOsBatches.GetSignInValidateDTOBatches(t, auth.SignInUserDTO{
		Nick:     "foobar",
		Password: "test123",
		TrustIP:  true,
	})
	tests.RunBatchTests(signInValidateDTOBatches)

	signUpValidateDTOBatches := DTOsBatches.GetSignUpValidateDTOBatches(t, mockedSignUpPayload)
	tests.RunBatchTests(signUpValidateDTOBatches)

	forgotPasswordValidateDTOBatches := DTOsBatches.GetFormatPasswordValidateDTOBatches(t, auth.ForgotPasswordDTO{
		Email: "text-api@example.com",
	})
	tests.RunBatchTests(forgotPasswordValidateDTOBatches)

	resetPasswordValidateDTOBatches := DTOsBatches.GetResetPasswordValidateDTOBatches(t, auth.ResetPasswordDTO{
		Code:                 "1bc23",
		Password:             "123test",
		LogoutFromAllDevices: true,
	})
	tests.RunBatchTests(resetPasswordValidateDTOBatches)

	blockUserValidateDTO := DTOsBatches.GetBlockUserValidateDTOBatches(t, blocks.BlockUserDTO{})
	tests.RunBatchTests(blockUserValidateDTO)

	replyQuestionValidateDTOBatches := DTOsBatches.GetReplyQuestionValidateDTOBatches(t, questions.ReplyQuestionDTO{})
	tests.RunBatchTests(replyQuestionValidateDTOBatches)

	editReplyQuestionValidateDTOBatches := DTOsBatches.GetEditReplyQuestionValidateDTOBatches(t, questions.EditQuestionReplyDTO{})
	tests.RunBatchTests(editReplyQuestionValidateDTOBatches)

	createReplyQuestionValidateDTOBatches := DTOsBatches.GetCreateQuestionValidateDTOBatches(t, questions.CreateQuestionDTO{})
	tests.RunBatchTests(createReplyQuestionValidateDTOBatches)

	createReportValidateDTOBatches := DTOsBatches.GetCreateReportValidateDTOBatches(t, reports.CreateReportDTO{
		Reason: "spam",
		Type:   "question",
	})
	tests.RunBatchTests(createReportValidateDTOBatches)

	updateUserProfileValidateDTOBatches := DTOsBatches.GetUpdateUserProfileValidateDTOBatches(t, users.UpdateProfileDTO{
		Email:  "api@example.com",
		Nick:   "foobar",
		Name:   "example",
		Locale: "en-US",
	})
	tests.RunBatchTests(updateUserProfileValidateDTOBatches)
}
