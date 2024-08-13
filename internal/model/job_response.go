package model

import (
	"notification-service/internal/entity"
	"time"
)

type JobResponse struct {
	JobID       uint      `json:"job_id"`
	Status      string    `json:"status"`
	QueuedAt    time.Time `json:"queued_at"`
	CompletedAt time.Time `json:"completed_at"`
}

func ToJobResponse(job *entity.Job) *JobResponse {
	return &JobResponse{
		JobID:       job.ID,
		Status:      job.Status,
		QueuedAt:    job.QueueAt,
		CompletedAt: job.CompletedAt,
	}
}
