package repositories

import (
	"context"
	collections "core/internal/constants"
	internal "core/internal/entities"
	pkg "core/pkg/entities"

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
func (b *Blocks) BlockUser(userToBlock, blockedById string) error {
	userToBlockId, err := pkg.ParseID(userToBlock)

	if err != nil {
		return err
	}

	userIdThatIsBlocking, err := pkg.ParseID(blockedById)

	if err != nil {
		return err
	}

	coll := b.db.Collection(collections.BLOCKS)

	block := internal.BlockedUser{
		UserToBlock: userToBlockId,
		BlockedBy:   userIdThatIsBlocking,
	}

	_, err = coll.InsertOne(context.Background(), block)

	return err
}

// IsBlocked returns if user is blocked by someone.
func (b *Blocks) IsUserBlocked(userId pkg.ID) (bool, error) {
	coll := b.db.Collection(collections.BLOCKS)

	filter := bson.D{{Key: "userToBlock", Value: userId}}
	foundUser := internal.User{}

	coll.FindOne(context.Background(), filter).Decode(&foundUser)

	return foundUser.Nick != "", nil
}
