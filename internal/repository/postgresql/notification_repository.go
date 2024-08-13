package postgresql

import (
	"gorm.io/gorm"
	"notification-service/internal/entity"
	"notification-service/internal/repository"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) repository.NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) CreateNotification(notification *entity.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) GetNotificationById(id uint) (*entity.Notification, error) {
	var notifications entity.Notification
	err := r.db.Where("id = ?", id).First(&notifications).Error

	return &notifications, err
}
