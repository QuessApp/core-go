package blocks

import (
	"context"

	collections "github.com/quessapp/toolkit/constants"
	toolkitEntities "github.com/quessapp/toolkit/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// BlocksRepository represents blocks repository.
type BlocksRepository struct {
	db *mongo.Database
}

// NewRepository returns blocks repository.
func NewRepository(db *mongo.Database) *BlocksRepository {
	return &BlocksRepository{db}
}

// BlockUser is a method of BlocksRepository that adds a new block for the given user ID.
// It takes in a BlockUserDTO and returns an error if there is one.
func (b *BlocksRepository) BlockUser(payload *BlockUserDTO) error {
	coll := b.db.Collection(collections.BLOCKS)

	block := BlockUserDTO{
		ID:          toolkitEntities.NewID(),
		UserToBlock: payload.UserToBlock,
		BlockedBy:   payload.BlockedBy,
	}

	_, err := coll.InsertOne(context.Background(), block)

	return err
}

// UnblockUser is a method of BlocksRepository that removes a block for the given user ID.
// It takes in a blockID of type toolkitEntities.ID and returns an error if there is one.
func (b *BlocksRepository) UnblockUser(blockID toolkitEntities.ID) error {
	coll := b.db.Collection(collections.BLOCKS)

	filter := bson.D{{Key: "userToBlock", Value: blockID}}

	_, err := coll.DeleteOne(context.Background(), filter)

	return err
}

// IsUserBlocked is a method of BlocksRepository that checks if a user is blocked given their ID.
// It takes in a userID of type toolkitEntities.ID and returns a boolean value.
func (b *BlocksRepository) IsUserBlocked(userID toolkitEntities.ID) bool {
	coll := b.db.Collection(collections.BLOCKS)

	filter := bson.D{{Key: "userToBlock", Value: userID}}
	foundRegistry := BlockedUser{}

	coll.FindOne(context.Background(), filter).Decode(&foundRegistry)

	return !toolkitEntities.IsZeroID(foundRegistry.ID)
}
