package service

import (
	"encoding/json"
	"notification-service/internal/entity"
	e "notification-service/internal/exception"
	"notification-service/internal/model"
	"notification-service/internal/repository"
	"notification-service/internal/utils"
	"time"
)

type JobService interface {
	SendJob(request model.RequestSendJob) (*model.JobResponse, error)
	GetJobById(id int) (*model.JobResponse, error)
}

type jobService struct {
	repo          repository.JobRepository
	rabbitmqUtils *utils.RabbitMQ
}

func NewJobService(repo repository.JobRepository, rabbitmqUtils *utils.RabbitMQ) JobService {
	return &jobService{repo: repo, rabbitmqUtils: rabbitmqUtils}
}

func (s *jobService) SendJob(request model.RequestSendJob) (*model.JobResponse, error) {
	err := utils.Validate(request)
	if err != nil {
		return nil, err
	}

	job := &entity.Job{
		Status:  entity.StatusQueued,
		QueueAt: time.Now(),
		Message: request.Message,
	}

	err = s.repo.CreateJob(job)
	if err != nil {
		return nil, e.Internal(err)
	}

	byteData, err := json.Marshal(job)
	if err != nil {
		return nil, e.Internal(err)
	}
	err = s.rabbitmqUtils.PublishMessage(utils.QUEUE_BROADCAST, string(byteData))
	if err != nil {
		return nil, e.Internal(err)
	}

	return model.ToJobResponse(job), nil
}

func (s *jobService) GetJobById(id int) (*model.JobResponse, error) {
	notification, err := s.repo.GetJobById(id)
	if err != nil {
		return nil, e.NotFound("job not found")
	}

	return model.ToJobResponse(notification), nil
}
