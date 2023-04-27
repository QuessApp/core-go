package mocks

import (
	"testing"

	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/pkg/tests"
	"github.com/stretchr/testify/assert"
)

// GetReplyQuestionValidateDTOMock returns a slice of BatchTest for ReplyQuestionDTO.
func GetReplyQuestionValidateDTOMock(t *testing.T, replyQuestionData questions.ReplyQuestionDTO) []tests.BatchTest {
	replyQuestionDataTest := []tests.BatchTest{
		{
			OnRun: func() {
				assert.EqualError(t, replyQuestionData.Validate(), "content_field_required.")
			},
		},
		{
			OnRun: func() {
				replyQuestionData.Content = "foobar"
				assert.NoError(t, replyQuestionData.Validate())
			},
		},
		{
			OnRun: func() {
				replyQuestionData.Content = tests.GenerateRandomString(300)
				assert.EqualError(t, replyQuestionData.Validate(), "content_field_length.")
			},
		},
	}

	return replyQuestionDataTest
}

// GetEditReplyQuestionValidateDTOMock returns a slice of BatchTest for EditQuestionReplyDTO.
func GetEditReplyQuestionValidateDTOMock(t *testing.T, editReplyQuestionData questions.EditQuestionReplyDTO) []tests.BatchTest {
	editReplyQuestionDataTest := []tests.BatchTest{
		{
			OnRun: func() {
				assert.EqualError(t, editReplyQuestionData.Validate(), "content_field_required.")
			},
		},
		{
			OnRun: func() {
				editReplyQuestionData.Content = "foobar"
				assert.NoError(t, editReplyQuestionData.Validate())
			},
		},
		{
			OnRun: func() {
				editReplyQuestionData.Content = tests.GenerateRandomString(300)
				assert.EqualError(t, editReplyQuestionData.Validate(), "content_field_length.")
			},
		},
	}

	return editReplyQuestionDataTest
}
