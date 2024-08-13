package http

import (
	"notification-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type jobController struct {
	jobService service.JobService
}

func NewJobController(jobService service.JobService) *jobController {
	return &jobController{jobService: jobService}
}

func (h *jobController) SendJob(c *fiber.Ctx) error {
	message := c.FormValue("message")

	err := h.jobService.SendJob(message)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "job sent"})
}

func (h *jobController) GetAllJobs(c *fiber.Ctx) error {
	jobs, err := h.jobService.GetAllJobs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(jobs)
}
