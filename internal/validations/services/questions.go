package validations

import (
	internalEntities "core/internal/entities"
	pkgEntities "core/pkg/entities"
	pkgErrors "core/pkg/errors"
	"errors"
)

// QuestionExists returns error message if question does not exisits in bd.
func QuestionExists(q *internalEntities.Question) error {
	if pkgEntities.IsZeroID(q.ID) {
		return errors.New(pkgErrors.QUESTION_NOT_FOUND)
	}

	return nil
}

// QuestionCanViewQuestion returns error message if user is not authorized to view the question.
func QuestionCanViewQuestion(q *internalEntities.Question, authenticatedUserId pkgEntities.ID) error {
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
func ReachedPostsLimitToCreateQuestion(u *internalEntities.User) error {
	if !u.IsPRO && u.PostsLimit <= 0 {
		return errors.New(pkgErrors.REACHED_QUESTIONS_LIMIT)
	}

	return nil
}

// CanUserDeleteQuestion returns error message if the user who is trying to delete the question is not the question owner.
func CanUserDeleteQuestion(q *internalEntities.Question, authenticatedUserId pkgEntities.ID) error {
	if q.SentBy != authenticatedUserId {
		return errors.New(pkgErrors.CANT_DELETE_QUESTION_NOT_SENT_BY_YOU)
	}

	return nil
}

// IsHiddenByReceiver returns error message if the question is already hidden by receiver.
func IsHiddenByReceiver(isHiddenByReceiver bool) error {
	if isHiddenByReceiver {
		return errors.New(pkgErrors.CANT_HIDE_ALREADY_HIDDEN)
	}

	return nil
}
