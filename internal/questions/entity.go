package questions

import (
	"time"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// ReplyHistory is a model for each reply in app.
type ReplyHistory struct {
	ID        toolkitEntities.ID `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	Content   string             `json:"content"`
}

// Questions is a model for each question in app.
type Question struct {
	ID      toolkitEntities.ID `json:"id" bson:"_id,omitempty"`
	Content string             `json:"content"`

	// SendTo represents the user that will receive the question. Type must bet Entities.ID, nill ou Entities.User
	SendTo any `json:"sendTo,omitempty" bson:"sendTo"`
	// SentBy represents the user who sent the question. Type must bet Entities.ID, nill ou Entities.User
	SentBy any `json:"sentBy,omitempty" bson:"sentBy"`
	// Reply is replied data content. Type must be Entities.Reply or nil.
	Reply any `json:"reply,omitempty"`

	IsAnonymous        bool `json:"isAnonymous" bson:"IsAnonymous"`
	IsHiddenByReceiver bool `json:"isHiddenByReceiver,omitempty" bson:"isHiddenByReceiver"`
	IsReplied          bool `json:"isReplied,omitempty" bson:"isReplied"`

	// RepliesHistory is the historic of how many times an user updated the question reply.
	RepliesHistory []ReplyHistory `json:"repliedHistory,omitempty" bson:"repliedHistory"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	RepliedAt time.Time `json:"repliedAt" bson:"repliedAt,omitempty"`
}

// PaginatedQuestions is a model for paginated questions in app.
type PaginatedQuestions struct {
	Questions  *[]Question `json:"questions"`
	TotalCount int64       `json:"totalCount"`
}

// MapAnonymousFields maps question in order to hide who sent the question, if the questions is anonymous.Otherwise, just returns the whole data.
func (q Question) MapAnonymousFields() *Question {
	if q.IsAnonymous {
		return &Question{
			ID:          q.ID,
			Content:     q.Content,
			IsAnonymous: q.IsAnonymous,
			CreatedAt:   q.CreatedAt,
			Reply:       q.Reply,
			IsReplied:   q.IsReplied,
			RepliedAt:   q.RepliedAt,
		}
	}
	return &q
}
