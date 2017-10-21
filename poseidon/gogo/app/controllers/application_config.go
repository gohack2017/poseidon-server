package controllers

import (
	"github.com/dolab/gogo"
	"github.com/dolab/session"
	"github.com/poseidon/lib/model"
)

// Application configuration specs
type AppConfig struct {
	Domain       string              `json:"domain"`
	GettingStart *GettingStartConfig `json:"getting_start"`

	Mongo  *model.Config      `json:"mongo"`
	Logger *gogo.LoggerConfig `json:"logger"`
	Cookie *session.Config    `json:"cookie"`
}

// NewAppConfig apply application config from *gogo.AppConfig
func NewAppConfig(config *gogo.AppConfig) error {
	return config.UnmarshalJSON(&Config)
}

// Sample application config for illustration
type GettingStartConfig struct {
	Greeting string `json:"greeting"`
}
