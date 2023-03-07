package questions

import (
	"core/internal/users"
	"errors"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// QuestionExists returns error message if question does not exisits in bd.
func QuestionExists(q *Question) error {
	if toolkitEntities.IsZeroID(q.ID) {
		return errors.New(QUESTION_NOT_FOUND)
	}

	return nil
}

// QuestionCanViewQuestion returns error message if user is not authorized to view the question.
func QuestionCanViewQuestion(q *Question, authenticatedUserId toolkitEntities.ID) error {
	if q.SendTo != authenticatedUserId && q.SentBy != authenticatedUserId {
		return errors.New(QUESTION_NOT_AUTHORIZED)
	}

	return nil
}

// IsSendingQuestionToYourself returns error message if users is sending questions to yourself.
func IsSendingQuestionToYourself(sendTo toolkitEntities.ID, authenticatedUserId toolkitEntities.ID) error {
	if sendTo == authenticatedUserId {
		return errors.New(SENDING_QUESTION_TO_YOURSELF)
	}

	return nil
}

// ReachedPostsLimitToCreateQuestion returns error message if user is not a pro member and reached posts limit.
func ReachedPostsLimitToCreateQuestion(u *users.User) error {
	if !u.IsPRO && u.PostsLimit <= 0 {
		return errors.New(REACHED_QUESTIONS_LIMIT)
	}

	return nil
}

// CanUserDeleteQuestion returns error message if the user who is trying to delete the question is not the question owner.
func CanUserDeleteQuestion(q *Question, authenticatedUserId toolkitEntities.ID) error {
	if q.SentBy != authenticatedUserId {
		return errors.New(CANT_DELETE_QUESTION_NOT_SENT_BY_YOU)
	}

	return nil
}

// CanHideQuestion returns error message if question is not sent to authenticated user.
func CanHideQuestion(q *Question, authenticatedUserId toolkitEntities.ID) error {
	if q.SendTo != authenticatedUserId {
		return errors.New(QUESTION_NOT_SENT_FOR_ME)
	}

	return nil
}

// IsHiddenByReceiver returns error message if the question is already hidden by receiver.
func IsHiddenByReceiver(isHiddenByReceiver bool) error {
	if isHiddenByReceiver {
		return errors.New(CANT_HIDE_ALREADY_HIDDEN)
	}

	return nil
}

// IsInvalidSendToID returns error message if the user to send id is invalid.
func IsInvalidSendToID(payload *CreateQuestionDTO) error {
	if toolkitEntities.IsZeroID(payload.SendTo) {
		return errors.New(CANT_SEND_INVALID_ID)
	}

	return nil
}

// IsAlreadyReplied returns error message if the question is already replied.
func IsAlreadyReplied(q *Question) error {
	if q.IsReplied {
		return errors.New(QUESTION_ALREADY_REPLIED)
	}

	return nil
}

// CanReply returns error message if the user can reply the question.
func CanReply(q *Question, authenticatedUserId toolkitEntities.ID) error {
	if q.SendTo != authenticatedUserId {
		return errors.New(QUESTION_NOT_SENT_FOR_ME)
	}

	return nil
}

// IsQuestionNotRepliedYet returns error message if user try to edit a reply that does not exists.
func IsQuestionNotRepliedYet(q *Question) error {
	if !q.IsReplied {
		return errors.New(CANT_EDIT_REPLY_NOT_REPLIED_YET)
	}

	return nil
}
