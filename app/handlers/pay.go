package handlers

import (
	"net/http"
	"payments/dto"

	"github.com/gofiber/fiber/v2"
)

type HandlerPay struct {
	Svc service.RoleService
}

func (h *HandlerPay) Pay(c *fiber.Ctx) error {
	body := new(dto.RequestPay)
	if err := parseBody(c, body); err != nil {
		return err
	}

	var msg fiber.Map
	message, err := h.Svc.Pay(*body)
	if err == nil {
		msg = *setMessage(message)
	}

	return resJSON(c, msg, err, http.StatusOK)
}

func (h *HandlerPay) Payments(c *fiber.Ctx) error {
	payments, err := h.Svc.Payments()

	return resJSON(c, payments, err, http.StatusOK)
}
