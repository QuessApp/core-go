package dtos

import (
	"testing"

	"github.com/quessapp/core-go/internal/auth"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	data := &auth.SignUpUserDTO{}

	data.Nick = "@ adsa32321@0-9"
	data.Format()
	assert.Equal(t, data.Nick, "adsa3232109")

	data.Nick = "_.test-user"
	data.Format()
	assert.Equal(t, data.Nick, "testuser")

	data.Nick = "-test.user@"
	data.Format()
	assert.Equal(t, data.Nick, "testuser")
}
