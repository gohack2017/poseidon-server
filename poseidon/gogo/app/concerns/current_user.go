package concerns

import (
	"github.com/dolab/gogo"
	"github.com/poseidon/app/models"
)

const (
	CurrentUserKey = "user"
	ContextUserKey = "current_user"
)

type CurrentUser struct {
	User *models.UserModel
}

func NewCurrentUser(user *models.UserModel) *CurrentUser {
	return &CurrentUser{
		User: user,
	}
}

func NewCurrentUserFromContext(ctx *gogo.Context) (user *CurrentUser) {
	return ctx.MustGetFinal(ContextUserKey).(*CurrentUser)
}
