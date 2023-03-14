package tests

import (
	"testing"

	"github.com/quessapp/core-go/tests/mocks"

	"github.com/stretchr/testify/assert"
)

func TestMapByAnonymousFields(t *testing.T) {
	fakeQuestion := mocks.NewQuestion()
	fakeQuestion.IsAnonymous = true

	mappedFields := fakeQuestion.MapAnonymousFields()

	assert.Nil(t, mappedFields.SentBy)
	assert.True(t, mappedFields.IsAnonymous)

	fakeQuestion.IsAnonymous = false

	assert.NotNil(t, fakeQuestion.MapAnonymousFields().SentBy)
	assert.False(t, fakeQuestion.MapAnonymousFields().IsAnonymous)
}
