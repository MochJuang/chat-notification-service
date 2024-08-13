package repository

import "notification-service/internal/entity"

type JobRepository interface {
	CreateJob(job *entity.Job) error
	GetAllJobs() ([]entity.Job, error)
}
