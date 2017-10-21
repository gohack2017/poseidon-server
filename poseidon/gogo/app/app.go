package app

import (
	"github.com/poseidon/app/controllers"
	"github.com/poseidon/app/middlewares"
)

type Application struct {
	*controllers.Application

	_ bool
}

func New(runMode, cfgPath string) *Application {
	app := &Application{
		Application: controllers.New(runMode, cfgPath),
	}

	return app
}

// Middlerwares implements gogo.Middlewarer
// NOTE: DO NOT change the method name, its required by gogo!
func (app *Application) Middlewares() {
	// apply your middlewares

	// panic recovery
	app.V1Use("*", middlewares.Recovery())
	app.V1Use("user", middlewares.Session(app.Session()))
}

// Run runs application after registering middelwares and resources
func (app *Application) Run() {
	// register middlewares
	app.Middlewares()

	// register resources
	app.Resources()

	// run server
	app.Application.Run()
}
