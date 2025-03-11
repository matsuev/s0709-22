package main

import (
	"fmt"
	"log"
	"s0709-22/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("App started")

	api := fiber.New()

	svc := service.New()

	api.Get("/", svc.RootHandler)

	if err := api.Listen(":7080"); err != nil {
		log.Fatalln(err)
	}
}
