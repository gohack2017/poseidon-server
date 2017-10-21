package controllers

import (
	"github.com/dolab/gogo"
	"github.com/poseidon/app/models"
	"github.com/poseidon/lib/errors"
)

type _Alert struct{}

var (
	Alert *_Alert
)

func (_ *_Alert) Index(ctx *gogo.Context) {
	marker := ctx.Params.Get("maker")
	limit, _ := ctx.Params.GetInt("limit")

	alerts, err := models.Alert.All(limit, marker)
	if err != nil {
		ctx.Logger.Errorf("models.Alert.All(%v, %v): %v", limit, marker, err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
		return
	}
	if alerts == nil {
		ctx.Json(make([]*models.AlertModel, 0))
		return
	}
	ctx.Json(alerts)
}
