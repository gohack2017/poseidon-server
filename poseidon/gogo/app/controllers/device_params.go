package controllers

type CreateDeviceInput struct {
	Password string `json:"password"`
	Address  string `json:"address"`
}
