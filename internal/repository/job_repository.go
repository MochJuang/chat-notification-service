package repository

import (
	"notification-service/internal/entity"
)

type JobRepository interface {
	CreateJob(job *entity.Job) error
	GetJobById(id int) (*entity.Job, error)
}
