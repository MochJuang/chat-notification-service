package postgresql

import (
	"gorm.io/gorm"
	"notification-service/internal/entity"
	"notification-service/internal/repository"
)

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) repository.JobRepository {
	return &jobRepository{db: db}
}

func (r *jobRepository) CreateJob(job *entity.Job) error {
	return r.db.Create(job).Error
}

func (r *jobRepository) GetJobById(id int) (*entity.Job, error) {
	var jobs entity.Job
	err := r.db.Where("id = ?", id).First(&jobs).Error
	return &jobs, err
}
