package batches

import (
	"testing"
	"time"

	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/pkg/tests"
	entitiesBatches "github.com/quessapp/core-go/tests/batches/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMapAnonymousFields(t *testing.T) {
	mapAnonymousFieldsBatches := entitiesBatches.GetTestMapByAnonymousFieldsBatches(t, questions.Question{
		Content:            "test",
		Reply:              primitive.NewObjectID(),
		IsAnonymous:        false,
		IsHiddenByReceiver: true,
		SentBy:             primitive.NewObjectID(),
		IsReplied:          true,
		RepliesHistory:     []questions.ReplyHistory{},
		CreatedAt:          time.Now(),
		RepliedAt:          &time.Time{},
	})
	tests.RunBatchTests(mapAnonymousFieldsBatches)
}
