package device

import (
	"github.com/RobertoSuarez/monitor-mikrotik/config/application"
	"github.com/RobertoSuarez/monitor-mikrotik/monitor"
)

type App struct {
	*Config
}

type Config struct {
	Dev monitor.MapDevice
}

//Implementamos la interfas
// implementa la interface MicroApp
func (app App) ConfigureApplication(application *application.Application) {
	dev := application.AppFiber.Group("/dev")
	ctr := Controller{Dev: app.Config.Dev}

	dev.Get("/", ctr.DeviceAll)
}

func New(cfg *Config) *App {
	return &App{cfg}
}
