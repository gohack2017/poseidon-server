package controllers

import (
	"sync"

	"github.com/dolab/gogo"
	"github.com/dolab/session"
	"github.com/poseidon/app/models"

	"github.com/poseidon/app/middlewares"
)

var (
	APP    *Application
	Config *AppConfig
)

type Application struct {
	*gogo.AppServer

	mux        sync.Mutex
	appSession *session.Session
	appLogger  gogo.Logger
	guest      *gogo.AppRoute
	user       *gogo.AppRoute
	admin      *gogo.AppRoute

	_ bool
}

func New(runMode, srcPath string) *Application {
	appServer := gogo.New(runMode, srcPath)

	err := NewAppConfig(appServer.Config())
	if err != nil {
		panic(err.Error())
	}

	appLogger := gogo.NewAppLogger(Config.Logger.Output, "")

	// setup model
	models.SetupModelWithConfig(Config.Mongo, appLogger)

	APP = &Application{
		AppServer: appServer,
		appLogger: appLogger,
		guest:     appServer.Group("v1.0"),
		user:      appServer.Group("v1.0"),
		admin:     appServer.Group("v1.0"),
	}

	return APP
}

func (app *Application) Session() *session.Session {
	if app.appSession != nil {
		return app.appSession
	}

	app.mux.Lock()
	defer app.mux.Unlock()

	if app.appSession != nil {
		return app.appSession
	}

	app.appSession = session.New(models.Session, Config.Cookie)

	return app.appSession
}

func (app *Application) V1Use(route string, middlewares ...gogo.Middleware) {
	switch route {
	case "*":
		app.guest.Use(middlewares...)
		app.user.Use(middlewares...)

	case "guest":
		app.guest.Use(middlewares...)

	case "user":
		app.user.Use(middlewares...)

	default:
		panic("Unknown route of " + route)
	}
}

// Middlerwares implements gogo.Middlewarer
// NOTE: DO NOT change the method name, its required by gogo!
func (app *Application) Middlewares() {
	// apply your middlewares

	// panic recovery
	app.Use(middlewares.Recovery())
}

// Resources implements gogo.Resourcer
// NOTE: DO NOT change the method name, its required by gogo!
func (app *Application) Resources() {
	// register your resources
	// app.GET("/", handler)

	app.GET("/@getting_start/hello", GettingStart.Hello)
}

// Run runs application after registering middelwares and resources
func (app *Application) Run() {
	// register middlewares
	app.Middlewares()

	// register resources
	app.Resources()

	// run server
	app.AppServer.Run()
}
