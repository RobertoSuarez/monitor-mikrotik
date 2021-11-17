package application

import (
	"github.com/gofiber/fiber/v2"
)

type MicroApp interface {
	ConfigureApplication(*Application)
}

type Config struct {
	AppFiber *fiber.App
	//MongoDB  *mongo.Client
}

type Application struct {
	*Config
}

func New(cfg *Config) *Application {
	if cfg == nil {
		cfg = &Config{}
	}

	if cfg.AppFiber == nil {
		cfg.AppFiber = fiber.New()
	}

	return &Application{
		Config: cfg,
	}
}

func (app *Application) Use(application MicroApp) {
	application.ConfigureApplication(app)
}
