package entities

import (
	"testing"

	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/pkg/tests"
	"github.com/stretchr/testify/assert"
)

// GetTestMapByAnonymousFieldsBatches returns a slice of BatchTest for TestMapByAnonymousFields method.
func GetTestMapByAnonymousFieldsBatches(t *testing.T, questionData questions.Question) []tests.BatchTest {
	return []tests.BatchTest{
		{
			OnRun: func() {
				questionData.IsAnonymous = true
				q := questionData.MapAnonymousFields()

				assert.True(t, q.IsAnonymous)
				assert.Nil(t, q.SentBy)
				assert.NotNil(t, q.Content)
				assert.NotNil(t, q.CreatedAt)
				assert.NotNil(t, q.Reply)
				assert.NotNil(t, q.IsReplied)
				assert.NotNil(t, q.RepliedAt)
				assert.NotNil(t, q.RepliesHistory)
			},
		},
		{
			OnRun: func() {
				questionData.IsAnonymous = false
				q := questionData.MapAnonymousFields()

				assert.False(t, q.IsAnonymous)
				assert.NotNil(t, q.SentBy)
			},
		},
	}
}
