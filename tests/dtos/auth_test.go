package dtos

import (
	"core/internal/dtos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	data := &dtos.SignUpUserDTO{}

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
