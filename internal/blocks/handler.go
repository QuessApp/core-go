package blocks

import (
	"core/configs"
	"core/internal/users"
	"core/pkg/jwt"
	"net/http"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"

	"github.com/kuriozapp/toolkit/responses"
)

// BlockUserHandler is a handler to block an user.
func BlockUserHandler(handlerCtx *configs.HandlersCtx, usersRepository *users.UsersRepository, blocksRepository *BlocksRepository) error {
	payload := BlockUserDTO{}
	id, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	payload.BlockedBy = jwt.GetUserByToken(handlerCtx.C).ID
	payload.UserToBlock = id

	if err := BlockUser(&payload, usersRepository, blocksRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}

// UnblockUserHandler is a handler to unblock an user.
func UnblockUserHandler(handlerCtx *configs.HandlersCtx, usersRepository *users.UsersRepository, blocksRepository *BlocksRepository) error {
	id, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	if err := UnblockUser(id, usersRepository, blocksRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}
