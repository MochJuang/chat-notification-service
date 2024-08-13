package http

import (
	"fmt"
	e "notification-service/internal/exception"
	"notification-service/internal/model"
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
	var req model.RequestSendJob
	if err := c.BodyParser(&req); err != nil {
		return e.Validation(err)
	}

	res, err := h.jobService.SendJob(req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *jobController) GetJobById(c *fiber.Ctx) error {
	jobId, err := c.ParamsInt("id")
	if err != nil {
		return e.Validation(fmt.Errorf("invalid id"))
	}

	jobs, err := h.jobService.GetJobById(jobId)
	if err != nil {
		return err
	}

	return c.JSON(jobs)
}
