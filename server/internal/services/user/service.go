package user

import "context"

type UserRepository interface {
	List(context.Context)
	Create(context.Context)
}

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Register(ctx context.Context) {}
func (s *Service) Login(ctx context.Context) {}
func (s *Service) Logout(ctx context.Context) {}
