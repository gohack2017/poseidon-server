package controllers

import (
	"github.com/poseidon/app/models"
	"github.com/poseidon/lib/errors"

	"github.com/dolab/gogo"
)

type _Device struct{}

var (
	Device *_Device
)

func (_ *_Device) Create(ctx *gogo.Context) {
	var input *CreateDeviceInput
	if err := ctx.Params.Json(&input); err != nil {
		ctx.Logger.Errorf("ctx.Params.Json(): %v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.MalformedParameter))
		return
	}

	device := models.NewDeviceModel(input.Password, input.Address)
	if err := device.Save(); err != nil {
		ctx.Logger.Errorf("device.save(): %v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
		return
	}
	ctx.Return()
}

func (_ *_Device) Index(ctx *gogo.Context) {
	marker := ctx.Params.Get("maker")
	limit, _ := ctx.Params.GetInt("limit")

	devices, err := models.Device.All(limit, marker)
	if err != nil {
		ctx.Logger.Errorf("models.Device.All(%v, %v): %v", limit, marker, err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
		return
	}
	ctx.Json(devices)
}

func (_ *_Device) Exist(ctx *gogo.Context) {
	var input *ShowDeviceInput
	if err := ctx.Params.Json(&input); err != nil {
		ctx.Logger.Errorf("ctx.Params.Json(): %v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.MalformedParameter))
		return
	}

	device, err := models.Device.FindByNum(input.Num)
	if err != nil {
		ctx.Logger.Errorf("models.Device.FindByNum(%v):%v", input.Num, err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
		return
	}

	if device.Password != input.Password {
		ctx.Logger.Errorf("invalid parameters")

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InvalidParameter))
		return
	}

	ctx.Json(device)
}
