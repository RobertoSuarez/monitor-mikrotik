package device

import (
	"github.com/RobertoSuarez/monitor-mikrotik/monitor"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Dev monitor.MapDevice
}

func (c *Controller) DeviceAll(ctx *fiber.Ctx) error {
	return ctx.JSON(c.Dev)
}
