package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) stat(ctx *fiber.Ctx) error {
	list := tryYoutubeRequest()
	if err := h.services.UpdateList(list); err != nil {
		return ctx.JSON(err.Error())
	}

	return ctx.JSON("Success update")
}
