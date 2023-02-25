package handlers

import (
	"core/cmd/app/entities"
	"core/internal/dtos"
	"core/internal/services"
	pkg "core/pkg/entities"
	"core/pkg/jwt"
	"core/pkg/responses"
	"net/http"
)

// BlockUserHandler is a handler to block an user.
func BlockUserHandler(handlerCtx *entities.HandlersContext) error {
	payload := dtos.BlockUserDTO{}
	id, err := pkg.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	payload.BlockedBy = jwt.GetUserByToken(handlerCtx.C).ID
	payload.UserToBlock = id

	if err := services.BlockUser(&payload, handlerCtx.UsersRepository, handlerCtx.BlocksRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}

// UnblockUserHandler is a handler to unblock an user.
func UnblockUserHandler(handlerCtx *entities.HandlersContext) error {
	id, err := pkg.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	if err := services.UnblockUser(id, handlerCtx.UsersRepository, handlerCtx.BlocksRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}
