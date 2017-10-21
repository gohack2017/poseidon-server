package models

import (
	"github.com/dolab/gogo"
	"github.com/poseidon/lib/model"
)

var (
	mongo *model.Model
)

func SetupModel(model *model.Model) {
	mongo = model
}

func SetupModelWithConfig(config *model.Config, logger gogo.Logger) {
	mongo = model.NewModel(config, logger)
}

func Model() *model.Model {
	return mongo
}
