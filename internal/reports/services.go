package reports

import (
	"strings"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/internal/queues/emails"
	"github.com/quessapp/core-go/internal/users"
	pkgReports "github.com/quessapp/core-go/pkg/reports"
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

		if err := questions.CanViewQuestion(q, authenticatedUserID); err != nil {
			return err
		}

		if err := IsReportingYourself(authenticatedUserID, q.SentBy.(toolkitEntities.ID)); err != nil {
			return err
		}
	}

	if err := reportsRepository.Create(payload); err != nil {
		return err
	}

	u := usersRepository.FindUserByID(authenticatedUserID)
	go emails.SendEmailThanksForReporting(handlerCtx.Cfg, handlerCtx.MessageQueueCh, handlerCtx.EmailsQueue, u)

	return nil
}

// FindReportByID retrieves a report with the given ID and verifies whether the authenticated user is authorized to view it.
// If the report is about a user or a question, this function will also fetch additional data
// about the reported user/question to be displayed in the UI
// It takes four parameters: handlerCtx, reportID, authenticatedUserID, and reportsRepository.
// handlerCtx is an instance of the HandlersCtx struct, which contains the fiber context and application context.
// reportID is the ID of the report to retrieve.
// authenticatedUserID is the ID of the authenticated user.
// reportsRepository is an instance of the ReportsRepository struct, which is used to access and modify report data.
// It returns a pointer to the Report struct and an error.
func FindReportByID(handlerCtx *configs.HandlersCtx, reportID, authenticatedUserID toolkitEntities.ID, reportsRepository *ReportsRepository, usersRepository *users.UsersRepository, questionsRepository *questions.QuestionsRepository) (*Report, error) {
	r, err := reportsRepository.FindByID(reportID)

	if err != nil {
		return nil, err
	}

	if err := ReportExists(r); err != nil {
		return nil, err
	}

	if err := CanViewReport(r, authenticatedUserID); err != nil {
		return nil, err
	}

	// if the user reported an user
	// we would like to show who is the reported user
	// this will help to show in UI
	if r.Type == "user" {
		u := usersRepository.FindUserByID(r.SendTo.(toolkitEntities.ID))

		r.SendTo = users.User{
			ID:        u.ID,
			Nick:      u.Nick,
			Name:      u.Name,
			AvatarURL: u.AvatarURL,
		}
	}

	if r.Type == "question" {
		q := questionsRepository.FindQuestionByID(r.SendTo.(toolkitEntities.ID))
		u := usersRepository.FindUserByID(q.SentBy.(toolkitEntities.ID))

		// if the user reported a question
		// we would like to show who sent the reported question
		// this will help to show in UI
		r.SendTo = questions.Question{
			ID:          q.ID,
			Content:     q.Content,
			IsAnonymous: q.IsAnonymous,
			CreatedAt:   q.CreatedAt,
			SentBy: users.User{
				ID:        u.ID,
				Nick:      u.Name,
				AvatarURL: u.AvatarURL,
				Name:      u.Name,
			},
		}.MapAnonymousFields()
	}

	return r, nil
}

// FindAllSent returns a paginated list of reports sent by the authenticated user, sorted by a specified field and order.
// If the page or sort parameters are nil, they will be set to their default values.
// The reports are retrieved from the provided ReportsRepository, and the user and question data is retrieved from their respective repositories.
// If a report refers to a user, its SendTo field will be replaced with the user's data, to be shown in the UI.
// If a report refers to a question, its SendTo field will be replaced with the question's data, along with the user who sent it, to be shown in the UI.
// The resulting paginated list of reports is returned, along with an error if one occurs during the retrieval process.
func FindAllSent(handlerCtx *configs.HandlersCtx, page *int64, sort *string, authenticatedUserID toolkitEntities.ID, reportsRepository *ReportsRepository, usersRepository *users.UsersRepository, questionsRepository *questions.QuestionsRepository) (*PaginatedReports, error) {
	if *page == 0 {
		*page = 1
	}

	if *sort == "" {
		*sort = "asc"
	}

	reports, err := reportsRepository.FindAllSentReports(authenticatedUserID, page, sort)

	if err != nil {
		return nil, err
	}

	if len(*reports.Reports) == 0 {
		return &PaginatedReports{
			Reports:    &[]Report{},
			TotalCount: 0,
		}, nil
	}

	var allReports []Report

	for _, r := range *reports.Reports {
		if r.Type == "user" {
			u := usersRepository.FindUserByID(r.SendTo.(toolkitEntities.ID))

			// if the user reported an user
			// we would like to show who is the reported user
			// this will help to show in UI
			r.SendTo = users.User{
				ID:        u.ID,
				Nick:      u.Nick,
				Name:      u.Name,
				AvatarURL: u.AvatarURL,
			}

			allReports = append(allReports, r)
		}

		if r.Type == "question" {
			q := questionsRepository.FindQuestionByID(r.SendTo.(toolkitEntities.ID))

			// id is zero, means that the question was deleted
			if toolkitEntities.IsZeroID(q.ID) {
				return nil, nil
			}

			// get sender question data
			u := usersRepository.FindUserByID(q.SentBy.(toolkitEntities.ID))

			// if the user reported a question
			// we would like to show who sent the reported question
			// this will help to show in UI
			r.SendTo = questions.Question{
				ID:          q.ID,
				Content:     q.Content,
				IsAnonymous: q.IsAnonymous,
				CreatedAt:   q.CreatedAt,
				SentBy: users.User{
					ID:        u.ID,
					Nick:      u.Name,
					AvatarURL: u.AvatarURL,
					Name:      u.Name,
				},
			}.MapAnonymousFields()

			allReports = append(allReports, r)
		}
	}

	result := PaginatedReports{
		Reports:    &allReports,
		TotalCount: reports.TotalCount,
	}

	return &result, nil
}

// DeleteReport is responsible for deleting a report.
// It receives four parameters: handlerCtx, reportID, authenticatedUserID, and reportsRepository.
// handlerCtx is an instance of the HandlersCtx struct, which contains the request context and other useful information.
// reportID is an instance of the toolkitEntities.ID struct, which represents the ID of the report to be deleted.
// authenticatedUserID is an instance of the toolkitEntities.ID struct, which represents the ID of the user making the request.
// reportsRepository is an instance of the ReportsRepository struct, which is used to access and modify report data.
// It returns an error if there was an issue deleting the report or if the user is not authorized to perform this action.
func DeleteReport(handlerCtx *configs.HandlersCtx, reportID, authenticatedUserID toolkitEntities.ID, reportsRepository *ReportsRepository) error {
	r, err := reportsRepository.FindByID(reportID)

	if err != nil {
		return err
	}

	if err := ReportExists(r); err != nil {
		return err
	}

	if err := CanUserDeleteReport(r, authenticatedUserID); err != nil {
		return err
	}

	if err := CanViewReport(r, authenticatedUserID); err != nil {
		return err
	}

	if err := reportsRepository.Delete(reportID); err != nil {
		return err
	}

	return nil
}

func ListAllKindReasonsOfReports() []string {
	return strings.Split(pkgReports.REASONS, ", ")
}
