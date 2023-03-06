package dtos

import (
	"core/internal/questions"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	data := &questions.CreateQuestionDTO{}

	data.Content = ""
	assert.EqualError(t, data.Validate(), "campo de conteúdo é obrigatório.")

	data.Content = "ab"
	assert.EqualError(t, data.Validate(), "campo de conteúdo deve conter entre 3 a 250 caracteres.")

	data.Content = "abc"
	assert.Nil(t, data.Validate())

	data.Content = strings.Join(make([]string, 300), "test")
	assert.EqualError(t, data.Validate(), "campo de conteúdo deve conter entre 3 a 250 caracteres.")
}
