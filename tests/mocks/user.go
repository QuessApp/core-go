package mocks

import (
	internalEntities "core/internal/entities"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"

	"github.com/jaswdr/faker"
)

// NewUser generates a fake user with generated fake fields values.
func NewUser() *internalEntities.User {
	fake := faker.New()

	return &internalEntities.User{
		ID:              toolkitEntities.NewID(),
		Nick:            fake.Internet().User(),
		Name:            fake.Person().Name(),
		AvatarURL:       fake.Internet().URL(),
		Password:        fake.Internet().Password(),
		Email:           fake.Internet().CompanyEmail(),
		EnableAppEmails: fake.Bool(),
		IsShadowBanned:  fake.Bool(),
		PostsLimit:      fake.RandomNumber(30),
		IsPRO:           fake.Bool(),
	}
}
