package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) list(ctx *fiber.Ctx) error {
	list, err := h.services.GetAllList()

	if err != nil {
		return ctx.JSON(err.Error())
	}

	return ctx.JSON(list.GetList())
}
