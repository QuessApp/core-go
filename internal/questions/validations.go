package questions

import (
	"errors"

	"github.com/quessapp/core-go/internal/users"

	pkgErrors "github.com/quessapp/core-go/pkg/errors"

	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// QuestionExists validates whether a question exists or not based on its ID.
func QuestionExists(q *Question) error {
	if toolkitEntities.IsZeroID(q.ID) {
		return errors.New(pkgErrors.QUESTION_NOT_FOUND)
	}

	return nil
}

// QuestionCanViewQuestion validates whether the authenticated user is authorized to view the question.
func QuestionCanViewQuestion(q *Question, authenticatedUserID toolkitEntities.ID) error {
	if q.SendTo != authenticatedUserID && q.SentBy != authenticatedUserID {
		return errors.New(pkgErrors.QUESTION_NOT_AUTHORIZED)
	}

	return nil
}

// IsSendingQuestionToYourself validates whether the user is trying to send a question to themselves.
func IsSendingQuestionToYourself(sendTo toolkitEntities.ID, authenticatedUserID toolkitEntities.ID) error {
	if sendTo == authenticatedUserID {
		return errors.New(pkgErrors.SENDING_QUESTION_TO_YOURSELF)
	}

	return nil
}

// ReachedPostsLimitToCreateQuestion validates whether the user has reached their monthly post limit and is not a PRO member
func ReachedPostsLimitToCreateQuestion(u *users.User) error {
	if !u.IsPRO && u.PostsLimit <= 0 {
		return errors.New(pkgErrors.REACHED_QUESTIONS_LIMIT)
	}

	return nil
}

// CanUserDeleteQuestion validates whether the user who is trying to delete the question is the question owner.
func CanUserDeleteQuestion(q *Question, authenticatedUserID toolkitEntities.ID) error {
	if q.SentBy != authenticatedUserID {
		return errors.New(pkgErrors.CANT_DELETE_QUESTION_NOT_SENT_BY_YOU)
	}

	return nil
}

// CanHideQuestion validates whether the question is sent to the authenticated user.
func CanHideQuestion(q *Question, authenticatedUserID toolkitEntities.ID) error {
	if q.SendTo != authenticatedUserID {
		return errors.New(pkgErrors.QUESTION_NOT_SENT_FOR_ME)
	}

	return nil
}

// IsHiddenByReceiver validates whether the question is already hidden by the receiver.
func IsHiddenByReceiver(isHiddenByReceiver bool) error {
	if isHiddenByReceiver {
		return errors.New(pkgErrors.CANT_HIDE_ALREADY_HIDDEN)
	}

	return nil
}

// IsInvalidSendToID validates whether the user ID to send the question is valid.
func IsInvalidSendToID(payload *CreateQuestionDTO) error {
	if toolkitEntities.IsZeroID(payload.SendTo) {
		return errors.New(pkgErrors.CANT_SEND_INVALID_ID)
	}

	return nil
}

// IsAlreadyReplied validates whether the question is already replied.
func IsAlreadyReplied(q *Question) error {
	if q.IsReplied {
		return errors.New(pkgErrors.QUESTION_ALREADY_REPLIED)
	}

	return nil
}

// CanReply validates whether the authenticated user can reply to the question.
func CanReply(q *Question, authenticatedUserID toolkitEntities.ID) error {
	if q.SendTo != authenticatedUserID {
		return errors.New(pkgErrors.QUESTION_NOT_SENT_FOR_ME)
	}

	return nil
}

// IsQuestionNotRepliedYet validates whether the user is trying to edit a reply that does not exist.
func IsQuestionNotRepliedYet(q *Question) error {
	if !q.IsReplied {
		return errors.New(pkgErrors.CANT_EDIT_REPLY_NOT_REPLIED_YET)
	}

	return nil
}

// ReachedLimitToEditReply validates whether the user has reached their limit to edit the question reply.
func ReachedLimitToEditReply(q *Question) error {
	// max is 5 but we add one more because when we edit a reply we add the prev content
	if len(q.RepliesHistory) >= 6 {
		return errors.New(pkgErrors.CANT_EDIT_REPLY_REACHED_LIMIT)
	}

	return nil
}
