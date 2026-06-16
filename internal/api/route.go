package api

import "github.com/gofiber/fiber/v2"

func (s *Server) Route(route fiber.Router) {
	route.Post("/item/", s.CreateItem)
	route.Put("item", s.UpdayteItem)
	route.Get("/items/", s.GetItems)
	route.Get("/items/:id", s.GetItem)
	route.Delete("item/:id", s.DeleteItem)
}
