package route

import (
	"notification-service/internal/config"
	httpdelivery "notification-service/internal/delivery/http"
	middleware "notification-service/internal/delivery/http/midlleware"
	"notification-service/internal/repository/postgresql"
	"notification-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg config.Config) {

	// Repositories
	notificationRepo := postgresql.NewNotificationRepository(cfg.DB)
	broadcastRepo := postgresql.NewJobRepository(cfg.DB)

	// Services
	notificationService := service.NewNotificationService(notificationRepo, cfg.RabbitMQUtils)
	broadcastService := service.NewJobService(broadcastRepo, cfg.RabbitMQUtils)

	// Controllers
	notificationController := httpdelivery.NewNotificationController(notificationService)
	broadcastController := httpdelivery.NewJobController(broadcastService)
	app.Use(middleware.ErrorControllerMiddleware)

	// Public routes
	app.Post("/notifications", notificationController.SendNotification)
	app.Get("/notifications/:id", notificationController.GetNotificationById)
	app.Post("/notifications/broadcasts", broadcastController.SendJob)
	app.Get("/jobs/:id", broadcastController.GetJobById)
}

func SetupRabbitMQ() {

}
