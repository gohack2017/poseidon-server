package controllers

import (
	"github.com/poseidon/app/concerns/kodo"
	uuid "github.com/satori/go.uuid"

	"github.com/dolab/gogo"
	"github.com/poseidon/app/models"
	"github.com/poseidon/lib/errors"
)

type _BuKong struct{}

var (
	BuKong *_BuKong
)

func (_ *_BuKong) Create(ctx *gogo.Context) {
	var input *CreateBukongInput
	if err := ctx.Params.Json(&input); err != nil {
		ctx.Logger.Errorf("ctx.Params.Json(): %v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.MalformedParameter))
		return
	}

	//send img to kodo bucket
	data, err := input.FaceImage()
	kodoclient := kodo.New(Config.Qiniu.Kodo)
	key := uuid.NewV4().String()
	err = kodoclient.Put(key, data)
	if err != nil {
		ctx.Logger.Errorf("kodo.client.Put():%v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
		return
	}

	//store record to mongo
	uri := "http://" + Config.Qiniu.Kodo.BucketDomain + "/" + key
	bukong := models.NewBukongModel(uri, input.Name, input.Phone, input.MonitorClass)
	if err := bukong.Save(); err != nil {
		ctx.Logger.Errorf("bukong.Save(): %v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
		return
	}

	//create index
	if err = FaceX.AddFace(uri, bukong.ID.Hex()); err != nil {
		ctx.Logger.Errorf("Facex.AddFace(%s, %s): %v", uri, bukong.ID.Hex(), err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
		return
	}

	ctx.Return()
}

func (_ *_BuKong) Index(ctx *gogo.Context) {
	marker := ctx.Params.Get("maker")
	limit, _ := ctx.Params.GetInt("limit")

	bukongs, err := models.BuKong.All(limit, marker)
	if err != nil {
		ctx.Logger.Errorf("models.Device.All(%v, %v): %v", limit, marker, err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
		return
	}
	ctx.Json(bukongs)
}

func (_ *_BuKong) Check(ctx *gogo.Context) {
	var input *UploadFileInput
	if err := ctx.Params.Json(&input); err != nil {
		ctx.Logger.Errorf("ctx.Params.Json(): %v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.MalformedParameter))
		return
	}

	device, err := models.Device.Find(input.DeviceID)
	if err != nil {
		ctx.Logger.Errorf("models.Device.Find(%v): %v", input.DeviceID, err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
		return
	}

	data, err := input.FaceImage()
	kodoclient := kodo.New(Config.Qiniu.Kodo)
	key := uuid.NewV4().String()
	err = kodoclient.Put(key, data)
	if err != nil {
		ctx.Logger.Errorf("kodo.client.Put():%v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
		return
	}

	uri := "http://" + Config.Qiniu.Kodo.BucketDomain + "/" + key
	go checkFace(ctx.Logger, uri, device.Address)

	ctx.Return()
}

func checkFace(logger gogo.Logger, uri, address string) {
	result, err := FaceX.Search(uri)
	if err != nil {
		logger.Errorf("Facex.Search(%s): %v", uri, err)

		return
	}

	if result.IsOK() {
		//step 1: create alert record
		bukongId := result.Name()
		bukong, err := models.BuKong.Find(bukongId)
		if err != nil {
			logger.Errorf("models.Bukong.Find(%v):%v", bukongId, err)
			return
		}

		alert := models.NewAlertModel(address, uri, bukong.URI, bukong.MonitorClass)
		if err = alert.Save(); err != nil {
			logger.Errorf("alert.Save():%v", err)
			return
		}

		//step 2: send alert message

	}
}

// func (_ *_BuKong) Upload(ctx *gogo.Context) {
// 	file, _, err := ctx.Request.FormFile("uploadfile")
// 	if err != nil {
// 		ctx.Logger.Errorf("ctx.Request.FormFile(uploadfile): %v", err)

// 		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
// 		return
// 	}
// 	defer file.Close()

// 	name := uuid.NewV4().String()
// 	data, err := ioutil.ReadAll(file)
// 	kodoclient := kodo.New(Config.Qiniu.Kodo)
// 	err = kodoclient.Put(name, data)
// 	if err != nil {
// 		ctx.Logger.Errorf("kodo.client.Put():%v", err)

// 		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InternalError))
// 		return
// 	}

// 	ctx.Return()
// }
