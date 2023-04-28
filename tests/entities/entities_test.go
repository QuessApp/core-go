package entities

import (
	"testing"

	"github.com/quessapp/core-go/internal/users"
	"github.com/quessapp/core-go/pkg/tests"
	"github.com/quessapp/core-go/tests/mocks"
)

func TestMapAnonymousFields(t *testing.T) {
	mapAnonymousFieldsBatches := GetTestMapByAnonymousFieldsBatches(t, *mocks.NewQuestionMock())
	tests.RunBatchTests(mapAnonymousFieldsBatches)
}

func TestFormat(t *testing.T) {
	formatUserDataBatches := GetFormatUserDataBatches(t, &users.User{
		Nick:  "test",
		Email: "",
	})
	tests.RunBatchTests(formatUserDataBatches)
}

func TestGetBasicInfos(t *testing.T) {
	getBasicUserDataBatches := GetBasicUserDataBatches(t, mocks.NewUserMock())
	tests.RunBatchTests(getBasicUserDataBatches)
}
