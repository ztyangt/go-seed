package system

import (
	"go-seed/app/config"

	"github.com/gofiber/fiber/v3"
)

func Setup(routerGroup fiber.Router, cfg *config.Cfg) {

	h := newHandler(cfg)

	systems := routerGroup.Group("/system")
	systems.Get("/", h.SystemInfo)
}
