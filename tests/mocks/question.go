package mocks

import (
	"time"

	"github.com/quessapp/core-go/internal/questions"

	"github.com/jaswdr/faker"
	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// NewQuestionMock generates a fake question with generated fake fields values.
func NewQuestionMock() *questions.Question {
	fake := faker.New()

	return &questions.Question{
		ID:                 toolkitEntities.NewID(),
		Content:            fake.Lorem().Text(100),
		SendTo:             toolkitEntities.NewID(),
		SentBy:             toolkitEntities.NewID(),
		Reply:              fake.Bool(),
		IsAnonymous:        fake.Bool(),
		IsHiddenByReceiver: fake.Bool(),
		IsReplied:          fake.Bool(),
		CreatedAt:          fake.Time().Time(time.Now()),
	}
}
