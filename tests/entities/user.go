package entities

import (
	"testing"

	"github.com/quessapp/core-go/internal/users"
	"github.com/quessapp/core-go/pkg/tests"
	"github.com/stretchr/testify/assert"
)

// GetFormatUserDataBatches returns a slice of BatchTest for User testing Format method.
func GetFormatUserDataBatches(t *testing.T, userData *users.User) []tests.BatchTest {
	return []tests.BatchTest{
		{
			OnRun: func() {
				userData.Nick = "@ adsa32321@0-9"
				userData.Format()
				assert.Equal(t, "adsa3232109", userData.Nick)

				userData.Nick = "_.test-user"
				userData.Format()
				assert.Equal(t, "testuser", userData.Nick)

				userData.Nick = "-test.user@"
				userData.Format()
				assert.Equal(t, "testuser", userData.Nick)

				userData.Nick = "test.user"
				userData.Format()
				assert.Equal(t, "testuser", userData.Nick)

				userData.Nick = "FOOBAR-USER"
				userData.Format()
				assert.Equal(t, "foobaruser", userData.Nick)

				userData.Nick = "FOOBAR_USER"
				userData.Format()
				assert.Equal(t, "foobaruser", userData.Nick)

				userData.Nick = "FOOBAR.USER"
				userData.Format()
				assert.Equal(t, "foobaruser", userData.Nick)

				userData.Nick = "foOBar@a-uSrR"
				userData.Format()
				assert.Equal(t, "foobarausrr", userData.Nick)
			},
		},
		{
			OnRun: func() {
				userData.Email = " caio@api.com"
				userData.Format()
				assert.Equal(t, "caio@api.com", userData.Email)

				userData.Email = "caio@api.com "
				userData.Format()
				assert.Equal(t, "caio@api.com", userData.Email)

				userData.Email = " caio@api.com "
				userData.Format()
				assert.Equal(t, "caio@api.com", userData.Email)
			},
		},
	}
}

// GetBasicUserDataBatches returns a slice of BatchTest for User testing GetBasicInfos method.
func GetBasicUserDataBatches(t *testing.T, userData *users.User) []tests.BatchTest {
	return []tests.BatchTest{
		{
			OnRun: func() {
				userData.Locale = "en-US"
				assert.NotNil(t, userData.GetBasicInfos().ID)
				assert.NotNil(t, userData.GetBasicInfos().Name)
				assert.NotNil(t, userData.GetBasicInfos().Nick)
				assert.NotNil(t, userData.GetBasicInfos().AvatarURL)
			},
		},
	}
}
