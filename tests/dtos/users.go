package dtos

import (
	"fmt"
	"testing"

	"github.com/quessapp/core-go/internal/users"
	"github.com/quessapp/core-go/pkg/tests"
	"github.com/stretchr/testify/assert"
)

// GetUpdateUserProfileValidateDTOBatches returns a slice of BatchTest for UpdateProfileDTO.
func GetUpdateUserProfileValidateDTOBatches(t *testing.T, updateProfileData users.UpdateProfileDTO) []tests.BatchTest {
	updateUserProfileDataTest := []tests.BatchTest{
		{
			OnRun: func() {
				updateProfileData.Nick = ""
				assert.ErrorContains(t, updateProfileData.Validate(), "nick_field_required")

				updateProfileData.Nick = tests.GenerateRandomString(51)
				assert.ErrorContains(t, updateProfileData.Validate(), "nick_field_length")

				updateProfileData.Nick = "_.test-user"
				assert.NoError(t, updateProfileData.Validate())
			},
		},
		{
			OnRun: func() {
				updateProfileData.Name = ""
				assert.ErrorContains(t, updateProfileData.Validate(), "name_field_required")

				updateProfileData.Name = "ab"
				assert.ErrorContains(t, updateProfileData.Validate(), "name_field_length")

				updateProfileData.Name = tests.GenerateRandomString(51)
				assert.ErrorContains(t, updateProfileData.Validate(), "name_field_length")

				updateProfileData.Name = tests.GenerateRandomString(20)
				assert.NoError(t, updateProfileData.Validate())
			},
		},
		{
			OnRun: func() {
				updateProfileData.Email = ""
				assert.ErrorContains(t, updateProfileData.Validate(), "email_field_required")

				updateProfileData.Email = tests.GenerateRandomString(130)
				assert.ErrorContains(t, updateProfileData.Validate(), "email_format_invalid")

				e := fmt.Sprintf("%s@%s.com", tests.GenerateRandomString(130), tests.GenerateRandomString(130))
				updateProfileData.Email = e
				assert.ErrorContains(t, updateProfileData.Validate(), "email_field_length")

				updateProfileData.Email = "test-api@example.com"
				assert.NoError(t, updateProfileData.Validate())
			},
		},
		{
			OnRun: func() {
				updateProfileData.Locale = "foobar"
				assert.ErrorContains(t, updateProfileData.Validate(), "locale_field_invalid")

				updateProfileData.Locale = ""
				assert.ErrorContains(t, updateProfileData.Validate(), "locale_field_required")

				updateProfileData.Locale = "es-ES"
				assert.NoError(t, updateProfileData.Validate())

				updateProfileData.Locale = "pt-ES"
				assert.ErrorContains(t, updateProfileData.Validate(), "locale_field_invalid")

				updateProfileData.Locale = "en-ES"
				assert.ErrorContains(t, updateProfileData.Validate(), "locale_field_invalid")

				updateProfileData.Locale = "pt-US"
				assert.ErrorContains(t, updateProfileData.Validate(), "locale_field_invalid")

				updateProfileData.Locale = "pt-BR"
				assert.NoError(t, updateProfileData.Validate())

				updateProfileData.Locale = "en-US"
				assert.NoError(t, updateProfileData.Validate())
			},
		},
	}

	return updateUserProfileDataTest
}
