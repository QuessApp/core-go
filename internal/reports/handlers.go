package reports

import (
	"net/http"
	"strconv"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/internal/users"
	toolkitEntities "github.com/quessapp/toolkit/entities"
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

// FindReportByIDHandler is responsible for handling requests to retrieve a report by its ID.
// It receives two parameters: handlerCtx and reportsRepository.
// handlerCtx is an instance of the HandlersCtx struct, which contains the fiber context and other data.
// reportsRepository is an instance of the ReportsRepository struct, which is used to access and modify report data.
// The function first parses the report ID from the request parameters and the authenticated user ID from the request context.
// It then calls the FindReportByID function to retrieve the report with the given ID, and checks if the user is authorized to view the report.
// Finally, it returns the report data in a successful response or an error response in case of failures.
func FindReportByIDHandler(handlerCtx *configs.HandlersCtx, reportsRepository *ReportsRepository, usersRepository *users.UsersRepository, questionsRepository *questions.QuestionsRepository) error {
	ID, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	r, err := FindReportByID(handlerCtx, ID, authenticatedUserID, reportsRepository, usersRepository, questionsRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, r)
}

// DeleteReportHandler is a function responsible for handling HTTP requests to delete a report.
// It receives two parameters: handlerCtx and reportsRepository.
// handlerCtx is an instance of the HandlersCtx struct, which contains the fiber.Ctx and other context information.
// reportsRepository is an instance of the ReportsRepository struct, which is used to access and modify report data.
// It parses the ID parameter from the request context and the authenticated user ID from the token.
// It calls the DeleteReport function passing the handlerCtx, report ID, authenticated user ID and reportsRepository as parameters.
// If the DeleteReport function returns an error, it returns a bad request response.
// Otherwise, it returns a successful response with status code 201.
func DeleteReportHandler(handlerCtx *configs.HandlersCtx, reportsRepository *ReportsRepository) error {
	ID, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	if err := DeleteReport(handlerCtx, ID, authenticatedUserID, reportsRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}

// FindAllSentReportsHandler handles the HTTP request to get all sent reports from a given user.
// It reads the authenticated user ID from the token, parses the "page" query parameter,
// and calls the FindAllSent function passing the necessary parameters to retrieve and sort the reports.
// If an error occurs during parsing or retrieving the reports, it returns an HTTP response with the error message.
// Otherwise, it returns an HTTP response with the retrieved reports.
func FindAllSentReportsHandler(handlerCtx *configs.HandlersCtx, reportsRepository *ReportsRepository, usersRepository *users.UsersRepository, questionsRepository *questions.QuestionsRepository) error {
	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	p, err := strconv.Atoi(handlerCtx.C.Query("page"))

	page := int64(p)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	sort := handlerCtx.C.Query("sort")

	reports, err := FindAllSent(handlerCtx, &page, &sort, authenticatedUserID, reportsRepository, usersRepository, questionsRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, reports)
}
