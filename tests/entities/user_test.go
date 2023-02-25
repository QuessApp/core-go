package tests

import (
	"core/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	fakeUser := mocks.NewUser()

	fakeUser.Nick = "@ adsa32321@0-9"
	fakeUser.Format()
	assert.Equal(t, fakeUser.Nick, "adsa3232109")

	fakeUser.Nick = "_.test-user"
	fakeUser.Format()
	assert.Equal(t, fakeUser.Nick, "testuser")

	fakeUser.Nick = "-test.user@"
	fakeUser.Format()
	assert.Equal(t, fakeUser.Nick, "testuser")

	newUser := fakeUser.GetBasicInfos()

	assert.Empty(t, newUser.GetBasicInfos().Password)
	assert.Empty(t, newUser.GetBasicInfos().Email)
	assert.Empty(t, newUser.GetBasicInfos().SubscriptionID)

	newUser.IsShadowBanned = true

	assert.NotEmpty(t, newUser.IsShadowBanned)

	assert.Empty(t, newUser.PostsLimit)

	newUser.PostsLimit = 12
	assert.Equal(t, newUser.PostsLimit, 12)
}
