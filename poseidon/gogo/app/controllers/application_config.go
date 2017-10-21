package controllers

import (
	"github.com/dolab/gogo"
	"github.com/dolab/session"
	"github.com/poseidon/app/concerns/facex"
	"github.com/poseidon/app/concerns/kodo"
	"github.com/poseidon/app/concerns/sms"
	"github.com/poseidon/lib/model"
)

// Application configuration specs
type AppConfig struct {
	Domain       string              `json:"domain"`
	GettingStart *GettingStartConfig `json:"getting_start"`
	Qiniu        *QiniuConfig        `json:"qiniu"`

	Mongo  *model.Config      `json:"mongo"`
	Logger *gogo.LoggerConfig `json:"logger"`
	Cookie *session.Config    `json:"cookie"`
	Facex  *facex.Config      `json:"facex"`
	SmsCfg *sms.Config        `json:"sms"`
}

// NewAppConfig apply application config from *gogo.AppConfig
func NewAppConfig(config *gogo.AppConfig) error {
	return config.UnmarshalJSON(&Config)
}

// Qiniu config
type QiniuConfig struct {
	Kodo *kodo.Config `json:"kodo"`
}

// Sample application config for illustration
type GettingStartConfig struct {
	Greeting string `json:"greeting"`
}
