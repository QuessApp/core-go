package validations

import (
	"core/internal/entities"
	internal "core/internal/entities"
	pkgEntities "core/pkg/entities"
	pkgErrors "core/pkg/errors"
	"errors"
)

// QuestionExists returns error message if question does not exisits in bd.
func QuestionExists(q *entities.Question) error {
	if q == nil {
		return errors.New(pkgErrors.QUESTION_NOT_FOUND)
	}

	return nil
}

// QuestionCanViewQuestion returns error message if user is not authorized to view the question.
func QuestionCanViewQuestion(q *internal.Question, authenticatedUserId pkgEntities.ID) error {
	if q.SendTo != authenticatedUserId && q.SentBy != authenticatedUserId {
		return errors.New(pkgErrors.QUESTION_NOT_AUTHORIZED)
	}

	return nil
}

// IsSendingQuestionToYourself returns error message if users is sending questions to yourself.
func IsSendingQuestionToYourself(sendTo pkgEntities.ID, authenticatedUserId pkgEntities.ID) error {
	if sendTo == authenticatedUserId {
		return errors.New(pkgErrors.SENDING_QUESTION_TO_YOURSELF)
	}

	return nil
}

// ReachedPostsLimitToCreateQuestion returns error message if user is not a pro member and reached posts limit.
func ReachedPostsLimitToCreateQuestion(u *entities.User) error {
	if !u.IsPRO && u.PostsLimit <= 0 {
		return errors.New(pkgErrors.REACHED_QUESTIONS_LIMIT)
	}

	return nil
}
