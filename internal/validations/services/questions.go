package validations

import (
	"core/internal/entities"
	internal "core/internal/entities"
	pkgEntities "core/pkg/entities"
	pkgErrors "core/pkg/errors"
	"errors"
)

// ValidateQuestionExists returns error message if question does not exisits in bd.
func ValidateQuestionExists(q entities.Question) error {
	if pkgEntities.IsZeroID(q.ID) && q.Content == "" {
		return errors.New(pkgErrors.QUESTION_NOT_FOUND)
	}

	return nil
}

// ValidateQuestionIsSentForMe returns error message if question is not sent for me (authenticated user).
func ValidateQuestionIsSentForMe(q internal.Question, authenticatedUserId pkgEntities.ID) error {
	if q.SendTo != authenticatedUserId {
		return errors.New(pkgErrors.QUESTION_NOT_SENT_FOR_ME)
	}

	return nil
}

// ValidateIsSendingQuestionToYourself returns error message if users is sending questions to yourself.
func ValidateIsSendingQuestionToYourself(sendTo pkgEntities.ID, authenticatedUserId pkgEntities.ID) error {
	if sendTo == authenticatedUserId {
		return errors.New(pkgErrors.SENDING_QUESTION_TO_YOURSELF)
	}

	return nil
}
