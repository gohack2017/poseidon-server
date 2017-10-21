package controllers

import (
	"encoding/base64"
	"strings"
)

type UploadFileInput struct {
	DeviceID string `json:"device_id"`
	Face     string `json:"face"`
}

func (in *UploadFileInput) FaceImage() (face []byte, err error) {
	tmp := strings.SplitN(in.Face, "base64,", 2)
	if len(tmp) == 2 {
		in.Face = tmp[1]
	}

	return base64.StdEncoding.DecodeString(in.Face)
}

type CreateBukongInput struct {
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	MonitorClass string `json:"class"`
	Face         string `json:"face"`
}

func (in *CreateBukongInput) FaceImage() (face []byte, err error) {
	tmp := strings.SplitN(in.Face, "base64,", 2)
	if len(tmp) == 2 {
		in.Face = tmp[1]
	}

	return base64.StdEncoding.DecodeString(in.Face)
}
