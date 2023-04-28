package healthcheck

import (
	"net/http"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/auth"
	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/internal/users"
	"github.com/quessapp/core-go/pkg/i18n"
	"github.com/quessapp/toolkit/responses"
)

// RunHealthCheckHandler runs a health check on the service.
// It takes four parameters, a HandlerCtx, an AuthRepository, a QuestionsRepository, and a UsersRepository.
// It returns an error if the health check is unsuccessful.
func RunHealthCheckHandler(handlerCtx *configs.HandlersCtx, authRepository *auth.AuthRepository, questionsRepository *questions.QuestionsRepository, usersRepository *users.UsersRepository) error {
	if err := Run(handlerCtx, authRepository, questionsRepository, usersRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, i18n.Translate(handlerCtx, err.Error()))
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, nil)
}
