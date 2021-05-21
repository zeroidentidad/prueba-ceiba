package app

import (
	"payments/app/handlers"
	"payments/domain"
	"payments/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Routes(app *fiber.App) {
	db := dbclient()
	payStorage := domain.NewPayStorageDb(db)
	hp := handlers.HandlerPay{Svc: service.NewPayService(payStorage)}

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	route := app.Group("/api")
	route.Post("/pagos", hp.Pay)
	route.Get("/pagos", hp.Payments)
}
