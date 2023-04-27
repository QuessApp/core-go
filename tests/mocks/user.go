package mocks

import (
	"github.com/quessapp/core-go/internal/users"

	toolkitEntities "github.com/quessapp/toolkit/entities"

	"github.com/jaswdr/faker"
)

// NewUserMock generates a fake user with generated fake fields values.
func NewUserMock() *users.User {
	fake := faker.New()

	return &users.User{
		ID:              toolkitEntities.NewID(),
		Nick:            fake.Internet().User(),
		Name:            fake.Person().Name(),
		AvatarURL:       fake.Internet().URL(),
		Password:        fake.Internet().Password(),
		Email:           fake.Internet().CompanyEmail(),
		EnableAPPEmails: fake.Bool(),
		IsShadowBanned:  fake.Bool(),
		PostsLimit:      fake.RandomNumber(30),
		IsPRO:           fake.Bool(),
	}
}
