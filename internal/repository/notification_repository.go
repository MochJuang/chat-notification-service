package repository

import "notification-service/internal/entity"

type NotificationRepository interface {
	CreateNotification(notification *entity.Notification) error
	GetNotificationById(id uint) (*entity.Notification, error)
}
