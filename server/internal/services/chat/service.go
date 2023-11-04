package chat

import (
	"context"

	"github.com/mieltn/keepintouch/internal/dto"
	pswd "github.com/mieltn/keepintouch/internal/services/password"
)

type ChatRepository interface {
	List(context.Context, *dto.ChatListReq) ([]*dto.Chat, error)
	Create(context.Context, *dto.ChatCreateReq) (*dto.Chat, error)
}

type Service struct {
	repo ChatRepository
}

func NewService(repo ChatRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List(ctx context.Context, in *dto.ChatListReq) ([]*dto.Chat, error) {
	return s.repo.List(ctx, in)
}

func (s *Service) Create(ctx context.Context, in *dto.ChatCreateReq) (*dto.Chat, error) {
	hash, err := pswd.HashPassword(in.Password)
	if err != nil {
		return nil, err
	}
	in.Password = hash
	chat, err := s.repo.Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return chat, nil
}