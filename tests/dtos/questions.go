package dtos

import (
	"testing"

	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/pkg/tests"
	"github.com/stretchr/testify/assert"
)

// GetReplyQuestionValidateDTOBatches returns a slice of BatchTest for ReplyQuestionDTO. testing Validate method
func GetReplyQuestionValidateDTOBatches(t *testing.T, replyQuestionData questions.ReplyQuestionDTO) []tests.BatchTest {
	return []tests.BatchTest{
		{
			OnRun: func() {
				assert.ErrorContains(t, replyQuestionData.Validate(), "content_field_required")

				replyQuestionData.Content = tests.GenerateRandomString(300)
				assert.ErrorContains(t, replyQuestionData.Validate(), "content_field_length")

				replyQuestionData.Content = "foobar"
				assert.NoError(t, replyQuestionData.Validate())
			},
		},
	}
}

// GetEditReplyQuestionValidateDTOBatches returns a slice of BatchTest for EditQuestionReplyDTO testing Validate method.
func GetEditReplyQuestionValidateDTOBatches(t *testing.T, editReplyQuestionData questions.EditQuestionReplyDTO) []tests.BatchTest {
	return []tests.BatchTest{
		{
			OnRun: func() {
				assert.ErrorContains(t, editReplyQuestionData.Validate(), "content_field_required")
			},
		},
		{
			OnRun: func() {
				editReplyQuestionData.Content = "foobar"
				assert.NoError(t, editReplyQuestionData.Validate())

				editReplyQuestionData.Content = tests.GenerateRandomString(300)
				assert.ErrorContains(t, editReplyQuestionData.Validate(), "content_field_length")
			},
		},
	}
}

// GetCreateQuestionValidateDTOBatches returns a slice of BatchTest for CreateQuestionDTO testing Validate method.
func GetCreateQuestionValidateDTOBatches(t *testing.T, createQuestionData questions.CreateQuestionDTO) []tests.BatchTest {
	return []tests.BatchTest{
		{
			OnRun: func() {
				assert.ErrorContains(t, createQuestionData.Validate(), "content_field_required")
			},
		},
		{
			OnRun: func() {
				createQuestionData.Content = "foobar"
				assert.NoError(t, createQuestionData.Validate())

				createQuestionData.Content = tests.GenerateRandomString(300)
				assert.ErrorContains(t, createQuestionData.Validate(), "content_field_length")
			},
		},
	}
}
