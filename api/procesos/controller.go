package procesos

import (
	"net/http"

	"github.com/RobertoSuarez/monitor-mikrotik/monitor"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Controller struct {
	Pro monitor.MapProcess
	Dev monitor.MapDevice
}

func (c *Controller) AllProcess(ctx *fiber.Ctx) error {
	return ctx.JSON(c.Pro.ToSlice())
}

// DleteProcess le da de baja a un proceso y lo elimina del mapa
// /process/:id DELETE
func (c *Controller) DeleteProcess(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Se requiere el id del proceso")
	}

	pro, ok := monitor.Pro[id]
	if !ok {
		return ctx.Status(http.StatusBadRequest).SendString("El proceso no existe")
	}

	pro.Cancel()
	return ctx.Status(http.StatusCreated).SendString("El proceso se ha eliminado")
}

// InitProcess inicia a monitorizar el divece que se pasa por id
// POST process/:id
func (c *Controller) InitProcess(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Se requiere el id del proceso")
	}

	dev, ok := c.Dev[id]
	if !ok {
		return ctx.Status(http.StatusBadRequest).SendString("El device no existe")
	}

	monitor.Monitorizar(dev)
	return ctx.Status(http.StatusCreated).SendString("El device se inicio a monitorizar")
}
