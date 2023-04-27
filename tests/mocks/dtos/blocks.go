package mocks

import (
	"testing"

	"github.com/quessapp/core-go/internal/blocks"
	"github.com/quessapp/core-go/pkg/tests"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetBlockUserValidateDTOMock returns a slice of BatchTest for BlockUserDTO.
func GetBlockUserValidateDTOMock(t *testing.T, blockUserData blocks.BlockUserDTO) []tests.BatchTest {
	blockUserDataTest := []tests.BatchTest{
		{
			OnRun: func() {
				assert.NoError(t, blockUserData.Validate())
			},
		},
		{
			OnRun: func() {
				blockUserData.UserToBlock = primitive.ObjectID{}
				assert.NoError(t, blockUserData.Validate())

				blockUserData.UserToBlock = primitive.NewObjectID()
				assert.NoError(t, blockUserData.Validate())
			},
		},
	}

	return blockUserDataTest
}