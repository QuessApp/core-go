package mocks

import (
	internalEntities "core/internal/entities"
	"time"

	"github.com/jaswdr/faker"
	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// NewQuestion generates a fake question with generated fake fields values.
func NewQuestion() *internalEntities.Question {
	fake := faker.New()

	return &internalEntities.Question{
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
