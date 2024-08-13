package model

type RequestSendNotification struct {
	UserID  uint   `json:"user_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}
