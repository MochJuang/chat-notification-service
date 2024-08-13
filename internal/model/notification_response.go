package model

import (
	"notification-service/internal/entity"
	"time"
)

type NotificationResponse struct {
	Id      uint      `json:"id"`
	UserId  uint      `json:"user_id"`
	Message string    `json:"message"`
	SendAt  time.Time `json:"send_at"`
}

func ToNotificationResponse(notification *entity.Notification) *NotificationResponse {
	return &NotificationResponse{
		Id:      notification.ID,
		UserId:  notification.UserID,
		Message: notification.Message,
		SendAt:  notification.CreatedAt,
	}
}
