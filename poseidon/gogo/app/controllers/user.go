package controllers

import (
	"net/http"

	"github.com/dolab/gogo"
	"github.com/dolab/session"
	"github.com/poseidon/app/concerns"
	"github.com/poseidon/app/models"
	"github.com/poseidon/lib/errors"
)

type _User struct{}

var (
	User *_User
)

func (_ *_User) Create(ctx *gogo.Context) {
	var input *CreateUserInput
	if err := ctx.Params.Json(&input); err != nil {
		ctx.Logger.Errorf("ctx.Parmas.Json(): %v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.MalformedParameter))
		return
	}

	user := models.NewUserModel(input.Email, input.Password)
	if err := user.Save(); err != nil {
		ctx.Logger.Errorf("user.Save(): %v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
		return
	}
	ctx.Return()
}

func (_ *_User) Login(ctx *gogo.Context) {
	sess := APP.Session()

	//destroy session first
	sess.Destroy(ctx.Response, ctx.Request)
	sess.New(ctx.Response, ctx.Request)

	sto, err := sess.Start(ctx.Response, ctx.Request)
	if err != nil {
		ctx.Logger.Errorf("session.Start(?, ?): %v", err)

		if err == session.ErrCookieExpired {

		}
		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.AuthFailure))
		return
	}

	var input UserLoginInput
	if err = ctx.Params.Json(&input); err != nil {
		ctx.Logger.Errorf("ctx.Paras.Json(%T): %v", input, err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.MalformedParameter))
		return
	}
	if !input.IsValid() {
		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InvalidParameter))
		return
	}

	user, err := models.User.FindByEmail(input.Email)
	if err != nil {
		ctx.Logger.Errorf("user.FindByEmail(%s): %v", input.Email, err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalFailure))
		return
	}

	if !user.IsValidPassword(input.Password) {
		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.AccessDenied))
		return
	}

	value := sto.GetValue()
	err = value.Set(concerns.CurrentUserKey, concerns.NewCurrentUser(user))
	if err != nil {
		ctx.Logger.Errorf("session.Set(%s, %T): %v", concerns.CurrentUserKey, concerns.CurrentUser{}, err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalFailure))
		return
	}

	err = sto.SetValue(value)
	if err != nil {
		ctx.Logger.Errorf("session.SetValue(): %v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalFailure))
		return
	}
}

func (_ *_User) Logout(ctx *gogo.Context) {
	err := APP.Session().Destroy(ctx.Response, ctx.Request)
	if err != nil {
		ctx.Logger.Errorf("APP.Session().Destroy(?,?): %v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
		return
	}
	ctx.Return(errors.NewResponse(ctx.RequestID(), ctx.RequestURI(), err, http.StatusNoContent))
}

func (_ *_User) Access(ctx *gogo.Context) {
	currentUser := concerns.NewCurrentUserFromContext(ctx)
	ctx.Json(NewUserLoginOutput(currentUser.User))
}
