package service

import (
	"encoding/json"
	"notification-service/internal/entity"
	e "notification-service/internal/exception"
	"notification-service/internal/model"
	"notification-service/internal/repository"
	"notification-service/internal/utils"
)

type NotificationService interface {
	SendNotification(request model.RequestSendNotification) (*model.NotificationResponse, error)
	GetNotificationById(id uint) (*model.NotificationResponse, error)
}

type notificationService struct {
	repo          repository.NotificationRepository
	rabbitmqUtils *utils.RabbitMQ
}

func NewNotificationService(repo repository.NotificationRepository, rabbitmqUtils *utils.RabbitMQ) NotificationService {
	return &notificationService{repo: repo, rabbitmqUtils: rabbitmqUtils}
}

func (s *notificationService) SendNotification(request model.RequestSendNotification) (*model.NotificationResponse, error) {
	err := utils.Validate(request)
	if err != nil {
		return nil, err
	}

	notification := &entity.Notification{
		UserID:  request.UserID,
		Message: request.Message,
	}
	err = s.repo.CreateNotification(notification)
	if err != nil {
		return nil, e.Internal(err)
	}

	byteData, err := json.Marshal(notification)
	if err != nil {
		return nil, e.Internal(err)
	}
	err = s.rabbitmqUtils.PublishMessage(utils.QUEUE_NOTIFICATION, string(byteData))
	if err != nil {
		return nil, e.Internal(err)
	}

	return model.ToNotificationResponse(notification), nil
}

func (s *notificationService) GetNotificationById(id uint) (*model.NotificationResponse, error) {
	notification, err := s.repo.GetNotificationById(id)
	if err != nil {
		return nil, e.NotFound("notification not found")
	}

	return model.ToNotificationResponse(notification), nil
}
