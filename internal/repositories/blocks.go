package repositories

import (
	"context"
	collections "core/internal/constants"
	"core/internal/dtos"
	internalEntities "core/internal/entities"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Blocks represents blocks repository.
type Blocks struct {
	db *mongo.Database
}

// NewBlocksRepository returns blocks repository.
func NewBlocksRepository(db *mongo.Database) *Blocks {
	return &Blocks{db}
}

// BlockUser blocks an user.
func (b *Blocks) BlockUser(payload *dtos.BlockUserDTO) error {
	coll := b.db.Collection(collections.BLOCKS)

	block := dtos.BlockUserDTO{
		ID:          toolkitEntities.NewID(),
		UserToBlock: payload.UserToBlock,
		BlockedBy:   payload.BlockedBy,
	}

	_, err := coll.InsertOne(context.Background(), block)

	return err
}

// UnblockUser removes block from database.
func (b *Blocks) UnblockUser(blockId toolkitEntities.ID) error {
	coll := b.db.Collection(collections.BLOCKS)

	filter := bson.D{{Key: "userToBlock", Value: blockId}}

	_, err := coll.DeleteOne(context.Background(), filter)

	return err
}

// IsBlocked returns if user is blocked by someone.
func (b *Blocks) IsUserBlocked(userId toolkitEntities.ID) bool {
	coll := b.db.Collection(collections.BLOCKS)

	filter := bson.D{{Key: "userToBlock", Value: userId}}
	foundRegistry := internalEntities.BlockedUser{}

	coll.FindOne(context.Background(), filter).Decode(&foundRegistry)

	return !toolkitEntities.IsZeroID(foundRegistry.ID)
}
