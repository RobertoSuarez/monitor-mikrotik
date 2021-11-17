package procesos

import (
	"github.com/RobertoSuarez/monitor-mikrotik/config/application"
	"github.com/RobertoSuarez/monitor-mikrotik/monitor"
)

type App struct {
	*Config
}

type Config struct {
	Pro *monitor.MapProcess
}

// implementa la interface MicroApp
func (app App) ConfigureApplication(application *application.Application) {
	process := application.AppFiber.Group("/process")
	ctr := Controller{
		Pro: monitor.Pro,
		Dev: monitor.Devs,
	}

	process.Get("/", ctr.AllProcess)
	process.Delete("/:id", ctr.DeleteProcess)
	process.Post("/:id", ctr.InitProcess)

}

// Contructur de la microapp
func New(cfg *Config) *App {
	return &App{cfg}
}
