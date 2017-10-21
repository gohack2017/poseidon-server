package model

import (
	"fmt"
	"testing"

	"github.com/dolab/gogo"
	"github.com/golib/assert"
)

func Test_NewConfig(t *testing.T) {
	assertion := assert.New(t)
	appConfig, _ := gogo.NewStringAppConfig(`{
    "sections": {
        "test": {
            "mongodb": {
                "host": "localhost:27017",
                "user": "root",
                "password": "",
                "database": "testing_model",
                "mode": "Strong",
                "pool": 5,
                "timeout": 5
            }
        }
    }
}`)
	appConfig.SetMode("test")

	modelConfig, err := NewConfig(appConfig)
	assertion.Nil(err)
	assertion.Equal("localhost:27017", modelConfig.Host)
	assertion.Equal("root", modelConfig.User)
	assertion.Equal("", modelConfig.Passwd)
	assertion.Equal("testing_model", modelConfig.Database)
	assertion.Equal("Strong", modelConfig.Mode)
	assertion.Equal(5, modelConfig.Pool)
	assertion.Equal(5, modelConfig.Timeout)
}

func Test_ConfigCopy(t *testing.T) {
	config := new(Config)
	copiedConfig := config.Copy()

	assert.Condition(t, func() bool {
		return fmt.Sprintf("%p", config) != fmt.Sprintf("%p", copiedConfig)
	})
}
