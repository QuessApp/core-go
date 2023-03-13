package reports

import (
	"core/configs"
	"core/internal/questions"
	"core/internal/users"
	"net/http"

	"github.com/quessapp/toolkit/responses"
)

// CreateReportHandler is a function that handles the creation of a report by receiving information from an HTTP request's body.
// The function uses the handlerCtx to access the request context and the reportsRepository, questionsRepository and usersRepository to handle the creation of the report.
// It starts by parsing the body of the HTTP request into a CreateReportDTO payload using the handlerCtx.C.BodyParser method.
// If an error occurs during the parsing, it returns an error response using the responses.ParseUnsuccesfull method.
// It then retrieves the authenticated user's ID from the request token using the users.GetUserByToken function and sets it as the SentBy field of the payload.
// Finally, the function calls the CreateReport function passing the handlerCtx, the authenticatedUserID, the questionsRepository, usersRepository and reportsRepository to create the report.
// If an error occurs during the creation of the report, it returns an error response using the responses.ParseUnsuccesfull method.
// If the report is created successfully, it returns a successful response with status 201 using the responses.ParseSuccessful method.
func CreateReportHandler(handlerCtx *configs.HandlersCtx, questionsRepository *questions.QuestionsRepository, usersRepository *users.UsersRepository, reportsRepository *ReportsRepository) error {
	payload := CreateReportDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	payload.SentBy = authenticatedUserID

	if err := CreateReport(handlerCtx, &payload, authenticatedUserID, questionsRepository, usersRepository, reportsRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}
