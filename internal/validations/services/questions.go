package validations

import (
	"core/internal/dtos"
	internalEntities "core/internal/entities"
	internalErrors "core/internal/errors"
	"errors"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// QuestionExists returns error message if question does not exisits in bd.
func QuestionExists(q *internalEntities.Question) error {
	if toolkitEntities.IsZeroID(q.ID) {
		return errors.New(internalErrors.QUESTION_NOT_FOUND)
	}

	return nil
}

// QuestionCanViewQuestion returns error message if user is not authorized to view the question.
func QuestionCanViewQuestion(q *internalEntities.Question, authenticatedUserId toolkitEntities.ID) error {
	if q.SendTo != authenticatedUserId && q.SentBy != authenticatedUserId {
		return errors.New(internalErrors.QUESTION_NOT_AUTHORIZED)
	}

	return nil
}

// IsSendingQuestionToYourself returns error message if users is sending questions to yourself.
func IsSendingQuestionToYourself(sendTo toolkitEntities.ID, authenticatedUserId toolkitEntities.ID) error {
	if sendTo == authenticatedUserId {
		return errors.New(internalErrors.SENDING_QUESTION_TO_YOURSELF)
	}

	return nil
}

// ReachedPostsLimitToCreateQuestion returns error message if user is not a pro member and reached posts limit.
func ReachedPostsLimitToCreateQuestion(u *internalEntities.User) error {
	if !u.IsPRO && u.PostsLimit <= 0 {
		return errors.New(internalErrors.REACHED_QUESTIONS_LIMIT)
	}

	return nil
}

// CanUserDeleteQuestion returns error message if the user who is trying to delete the question is not the question owner.
func CanUserDeleteQuestion(q *internalEntities.Question, authenticatedUserId toolkitEntities.ID) error {
	if q.SentBy != authenticatedUserId {
		return errors.New(internalErrors.CANT_DELETE_QUESTION_NOT_SENT_BY_YOU)
	}

	return nil
}

// CanHideQuestion returns error message if question is not sent to authenticated user.
func CanHideQuestion(q *internalEntities.Question, authenticatedUserId toolkitEntities.ID) error {
	if q.SendTo != authenticatedUserId {
		return errors.New(internalErrors.QUESTION_NOT_SENT_FOR_ME)
	}

	return nil
}

// IsHiddenByReceiver returns error message if the question is already hidden by receiver.
func IsHiddenByReceiver(isHiddenByReceiver bool) error {
	if isHiddenByReceiver {
		return errors.New(internalErrors.CANT_HIDE_ALREADY_HIDDEN)
	}

	return nil
}

// IsInvalidSendToID returns error message if the user to send id is invalid.
func IsInvalidSendToID(payload *dtos.CreateQuestionDTO) error {
	if toolkitEntities.IsZeroID(payload.SendTo) {
		return errors.New(internalErrors.CANT_SEND_INVALID_ID)
	}

	return nil
}
