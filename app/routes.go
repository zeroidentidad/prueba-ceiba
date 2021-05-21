package app

import (
	"payments/app/handlers"
	"payments/domain"
	"payments/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func routes(prefix string) *fiber.App {
	router := fiber.New()
	router.Use(logger.New())
	router.Use(cors.New(cors.Config{AllowCredentials: true}))

	db := dbclient()
	payStorage := domain.NewPayStorageDb(db)
	hp := handlers.HandlerPay{Svc: service.NewPayService(payStorage)}

	router.Post(prefix+"/pagos", hp.Pay)
	router.Get(prefix+"/pagos", hp.Payments)

	return router
}
