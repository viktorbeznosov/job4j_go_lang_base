package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/viktorbeznosov/job4j_go_lang_base/internal/tracker"
)

type GetItemResponse struct {
	Item tracker.Item `json:"item"`
}

func (s *Server) GetItem(c *fiber.Ctx) error {
	id := c.Params("id")

	item, err := s.Repository.Get(c.Context(), id)
	if err != nil {
		log.Errorw("s.Repository.Get", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(GetItemResponse{Item: item})

}
