package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/viktorbeznosov/job4j_go_lang_base/internal/tracker"
)

type UpdateItemRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (s *Server) UpdayteItem(c *fiber.Ctx) error {
	var req UpdateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}
	if req.Id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	err := s.Repository.Update(c.Context(), tracker.Item{
		ID:   req.Id,
		Name: req.Name,
	})
	if err != nil {
		log.Errorw("s.Repository.Update", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
