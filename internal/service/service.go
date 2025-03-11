package service

import "github.com/gofiber/fiber/v2"

// Service ...
type Service struct{}

// New ...
func New() *Service {
	return new(Service)
}

// RootHandler ...
func (s *Service) RootHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello!!!")
}
