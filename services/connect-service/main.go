package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/connect", connectHandler)

	if err := app.Listen("127.0.0.1:6080"); err != nil {
		log.Fatalln(err)
	}
}

func connectHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"result": fiber.Map{
			"user": "alex",
		},
	})
}
