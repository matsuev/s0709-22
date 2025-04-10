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

// ConnectRequest ..
type ConnectRequest struct {
	Client string      `json:"client"`
	Data   UserConnect `json:"data"`
}

// UserConnect ...
type UserConnect struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ConnectResponse ...
type ConnectResponse struct {
	Result     *ConnectResult `json:"result,omitempty"`
	Error      *Error         `json:"error,omitempty"`
	Disconnect *Disconnect    `json:"disconnect,omitempty"`
}

// Error ...
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

// Disconnect ...
type Disconnect struct {
	Code   int    `json:"code"`
	Reason string `json:"reason,omitempty"`
}

// ConnectResult ...
type ConnectResult struct {
	User string `json:"user"`
}

func connectHandler(ctx *fiber.Ctx) error {
	req := &ConnectRequest{}

	if err := ctx.BodyParser(req); err != nil {
		return ctx.JSON(ConnectResponse{
			Error: &Error{
				Code:    107,
				Message: "bad request",
			},
		})
	}

	if req.Data.Username != "alex" || req.Data.Password != "password" {
		return ctx.JSON(ConnectResponse{
			Error: &Error{
				Code:    101,
				Message: "unauthorized",
			},
		})
	}

	return ctx.JSON(ConnectResponse{
		Result: &ConnectResult{
			User: req.Data.Username,
		},
	})
}
