package app

import (
	"fmt"
	"os"
	"os/signal"
	"payments/logs"

	"github.com/gofiber/fiber/v2"
)

func serve(app *fiber.App) {
	addr := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	go func() {
		logs.Info(fmt.Sprintf("Starting server on %s:%s ...", addr, port))
		err := app.Listen(fmt.Sprintf("%s:%s", addr, port))
		logs.Fatal(err.Error())
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
