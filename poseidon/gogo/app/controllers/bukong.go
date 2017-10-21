package controllers

import (
	"github.com/poseidon/app/concerns/kodo"
	uuid "github.com/satori/go.uuid"

	"github.com/dolab/gogo"
	"github.com/poseidon/lib/errors"
)

type _BuKong struct{}

var (
	BuKong *_BuKong
)

func (_ *_BuKong) Check(ctx *gogo.Context) {
	var input *UploadFileInput
	if err := ctx.Params.Json(&input); err != nil {
		ctx.Logger.Errorf("ctx.Params.Json(): %v", err)

		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.MalformedParameter))
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
	go checkFace(ctx.Logger, uri)

	ctx.Return()
}

func checkFace(logger gogo.Logger, uri string) {
	result, err := FaceX.Search(uri)
	if err != nil {
		logger.Errorf("Facex.Search(%s): %v", uri, err)

		return
	}

	if result.IsOK() {
		//step 1: store record to mongo

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
