package blocks

import (
	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/users"
	i18n "github.com/quessapp/core-go/pkg/i18n"
	toolkitEntities "github.com/quessapp/toolkit/entities"
	"github.com/quessapp/toolkit/responses"

	"net/http"
)

// BlockUserHandler is a handler function that blocks a user given their ID.
// It takes in a HandlersCtx, a UsersRepository, and a BlocksRepository, and returns an error if there is one.
func BlockUserHandler(handlerCtx *configs.HandlersCtx, usersRepository *users.UsersRepository, blocksRepository *BlocksRepository) error {
	payload := BlockUserDTO{}
	ID, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, i18n.Translate(handlerCtx, err.Error()))
	}

	payload.BlockedBy = users.GetUserByToken(handlerCtx).ID
	payload.UserToBlock = ID

	if err := BlockUser(&payload, usersRepository, blocksRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, i18n.Translate(handlerCtx, err.Error()))
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}

// UnblockUserHandler is a handler function that unblocks a user given their ID.
// It takes in a HandlersCtx, a UsersRepository, and a BlocksRepository, and returns an error if there is one.
func UnblockUserHandler(handlerCtx *configs.HandlersCtx, usersRepository *users.UsersRepository, blocksRepository *BlocksRepository) error {
	id, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, i18n.Translate(handlerCtx, err.Error()))
	}

	if err := UnblockUser(id, usersRepository, blocksRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, i18n.Translate(handlerCtx, err.Error()))
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}
