package model

type RequestSendJob struct {
	Message string `json:"message" validate:"required"`
}
