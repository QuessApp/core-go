package blocks

import toolkitEntities "github.com/kuriozapp/toolkit/entities"

// BlockedUser is a model for each blocked user in app.
type BlockedUser struct {
	ID          toolkitEntities.ID `json:"id" bson:"_id" `
	UserToBlock toolkitEntities.ID `json:"userToBlock" bson:"userToBlock"`
	BlockedBy   toolkitEntities.ID `json:"blockedBy" bson:"blockedBy"`
}
