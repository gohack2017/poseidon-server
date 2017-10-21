package controllers

type CreateDeviceInput struct {
	Password string `json:"password"`
	Address  string `json:"address"`
}

type ShowDeviceInput struct {
	Num      string `json:"num"`
	Password string `json:"password"`
}
