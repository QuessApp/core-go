package reports

import (
	"github.com/quessapp/core-go/configs"

	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/internal/users"

	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// CreateReport is a function that creates a report based on the given payload and saves it in the reportsRepository.
// It also performs some validation checks before creating the report.
// The function receives a handlerCtx which contains information about the request, the payload which contains information about the report to be created, the authenticatedUserID of the user creating the report,
// the questionsRepository and usersRepository for finding the target user or question and the reportsRepository to save the report.
// If the report has already been sent, an error is returned.
// If the report is of type "user", it checks if the target user exists in the usersRepository.
// If the report is of type "question", it checks if the target question exists in the questionsRepository.
// If all checks pass, the payload is saved in the reportsRepository and no error is returned.
func CreateReport(handlerCtx *configs.HandlersCtx, payload *CreateReportDTO, authenticatedUserID toolkitEntities.ID, questionsRepository *questions.QuestionsRepository, usersRepository *users.UsersRepository, reportsRepository *ReportsRepository) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	if err := AlreadySent(reportsRepository.AlreadySent(payload)); err != nil {
		return err
	}

	if payload.Type == "user" {
		u := usersRepository.FindUserByID(payload.SendTo)

		if err := users.UserExists(u); err != nil {
			return err
		}

		if err := IsReportingYourself(authenticatedUserID, u.ID); err != nil {
			return err
		}
	}

	if payload.Type == "question" {
		q := questionsRepository.FindQuestionByID(payload.SendTo)

		if err := questions.QuestionExists(q); err != nil {
			return err
		}

		if err := questions.QuestionCanViewQuestion(q, authenticatedUserID); err != nil {
			return err
		}

		if err := IsReportingYourself(authenticatedUserID, q.SentBy.(toolkitEntities.ID)); err != nil {
			return err
		}
	}

	if err := reportsRepository.Create(payload); err != nil {
		return err
	}

	return nil
}
