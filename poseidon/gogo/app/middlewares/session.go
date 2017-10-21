package middlewares

import (
	"github.com/dolab/gogo"
	"github.com/dolab/session"
	"github.com/poseidon/app/concerns"
	"github.com/poseidon/lib/errors"
)

func Session(sess *session.Session) gogo.Middleware {
	return func(ctx *gogo.Context) {
		tmpsess, err := sess.Start(ctx.Response, ctx.Request)
		if err != nil {
			ctx.Logger.Errorf("session.Start(?, ?): %v", err)

			ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError, err))
			return
		}

		var currentUser concerns.CurrentUser

		err = tmpsess.GetValue().Unmarshal(concerns.CurrentUserKey, &currentUser)
		if err != nil {
			ctx.Logger.Errorf("session.Unmarshal(%s, %T): %v", concerns.CurrentUserKey, currentUser, err)

			ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.AuthFailure, err))
			return
		}

		ctx.MustSetFinal(concerns.ContextUserKey, &currentUser)

		ctx.Next()
	}
}
