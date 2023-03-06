package mocks

import (
	"core/internal/questions"
	"time"

	"github.com/jaswdr/faker"
	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// NewQuestion generates a fake question with generated fake fields values.
func NewQuestion() *questions.Question {
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
