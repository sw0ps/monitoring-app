package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sw0ps/monitoring-app/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(app *fiber.App) *fiber.App {

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, User")
	})
	app.Get("/stat", h.stat)
	app.Get("/list", h.list)

	return app
}