package service

import (
	"notification-service/internal/entity"
	"notification-service/internal/repository"
)

type JobService interface {
	SendJob(message string) error
	GetAllJobs() ([]entity.Job, error)
}

type jobService struct {
	repo repository.JobRepository
}

func NewJobService(repo repository.JobRepository) JobService {
	return &jobService{repo: repo}
}

func (s *jobService) SendJob(message string) error {
	//job := &entity.Job{
	//	Message: message,
	//}
	//return s.repo.CreateJob(job)

	return nil
}

func (s *jobService) GetAllJobs() ([]entity.Job, error) {
	return s.repo.GetAllJobs()
}
