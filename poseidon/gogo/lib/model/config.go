package model

import (
	"github.com/dolab/gogo"
)

type Config struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Passwd   string `json:"password"`
	Database string `json:"database"`
	Mode     string `json:"mode"`
	Pool     int    `json:"pool"`
	Timeout  int    `json:"timeout"`
}

func NewConfig(config *gogo.AppConfig) (*Config, error) {
	var mongoConfig struct {
		Config *Config `json:"mongodb"`
	}

	err := config.UnmarshalJSON(&mongoConfig)
	if err != nil {
		return nil, err
	}

	return mongoConfig.Config, err
}

func (c *Config) Copy() *Config {
	config := *c

	return &config
}
