package repositories

import (
	"context"
	collections "core/internal/constants"
	"core/internal/dtos"
	internal "core/internal/entities"
	pkgEntities "core/pkg/entities"
	pkgErrors "core/pkg/errors"

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
		ID:          pkgEntities.NewID(),
		UserToBlock: payload.UserToBlock,
		BlockedBy:   payload.BlockedBy,
	}

	_, err := coll.InsertOne(context.Background(), block)

	return err
}

// IsBlocked returns if user is blocked by someone.
func (b *Blocks) IsUserBlocked(userId pkgEntities.ID) (bool, error) {
	coll := b.db.Collection(collections.BLOCKS)

	filter := bson.D{{Key: "userToBlock", Value: userId}}
	foundRegistry := internal.BlockedUser{}

	err := coll.FindOne(context.Background(), filter).Decode(&foundRegistry)

	if err.Error() == pkgErrors.MONGO_NO_DOCUMENTS {
		return false, nil
	}

	areValidIds := !pkgEntities.IsZeroID(foundRegistry.ID) && !pkgEntities.IsZeroID(foundRegistry.UserToBlock)

	return areValidIds, err
}
