package local_tasks

import (
	"context"
	"my_pocket_taskbook/internal/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]models.Task, error)
	GetByID(ctx context.Context, id int) (*models.Task, error)
	GetAllCurrent(ctx context.Context) ([]models.Task, error)
	Create(ctx context.Context, t *models.Task) (*models.Task, error)
	Edit(ctx context.Context, t *models.Task, id int) (*models.Task, error)
	ChangeStatus(ctx context.Context, id int, status string) (*models.Task, error)
}

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) GetAll(ctx context.Context) ([]models.Task, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id int) (*models.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetAllCurrent(ctx context.Context) ([]models.Task, error) {
	return s.repo.GetAllCurrent(ctx)
}

func (s *Service) Create(ctx context.Context, t *models.Task) (*models.Task, error) {
	return s.repo.Create(ctx, t)
}

func (s *Service) Edit(ctx context.Context, t *models.Task, id int) (*models.Task, error) {
	return s.repo.Edit(ctx, t, id)
}

func (s *Service) ChangeStatus(ctx context.Context, id int, status string) (*models.Task, error) {
	return s.repo.ChangeStatus(ctx, id, status)
}
