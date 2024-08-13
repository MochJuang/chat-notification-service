package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	e "notification-service/internal/exception"
	"notification-service/internal/model"
	"notification-service/internal/service"
)

type NotificationController struct {
	notificationService service.NotificationService
}

func NewNotificationController(notificationService service.NotificationService) *NotificationController {
	return &NotificationController{notificationService: notificationService}
}

func (h *NotificationController) SendNotification(c *fiber.Ctx) error {
	var req model.RequestSendNotification
	if err := c.BodyParser(&req); err != nil {
		return e.Validation(err)
	}

	res, err := h.notificationService.SendNotification(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *NotificationController) GetNotificationById(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return e.Validation(fmt.Errorf("invalid id"))
	}

	notifications, err := h.notificationService.GetNotificationById(uint(userId))
	if err != nil {
		return err
	}

	return c.JSON(notifications)
}
