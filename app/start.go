package app

import (
	"github.com/gofiber/fiber/v2"
)

func Start() {
	config()
	app := fiber.New()
	Routes(app)
	serve(app)
}
