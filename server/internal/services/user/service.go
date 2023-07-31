package user

import (
	"context"
	"errors"

	"github.com/mieltn/keepintouch/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

var errIncorrectPassword = errors.New("Incorrect password")

type UserRepository interface {
	GetUserByEmail(context.Context, string) (*dto.User, error)
	Create(context.Context, *dto.UserCreateReq) (*dto.User, error)
}

type JWTService interface {
	CreateTokens(*dto.User) (*dto.UserAuthRes, error)
	Refresh(string) (*dto.UserAuthRes, error)
	Validate(string) (bool, error)
}

type Service struct {
	repo UserRepository
	JWTService JWTService
}

func NewService(repo UserRepository, JWTService JWTService) *Service {
	return &Service{
		repo: repo,
		JWTService: JWTService,
	}
}

func (s *Service) Register(ctx context.Context, in *dto.UserCreateReq) (*dto.UserCreateRes, error) {
	hash, err := hashPassword(in.Password)
	if err != nil {
		return nil, err
	}
	in.Password = hash
	user, err := s.repo.Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return &dto.UserCreateRes{
		Id: user.Id,
		Username: user.Username,
		Email: user.Email,
	}, nil
}

func (s *Service) Login(ctx context.Context, in *dto.UserLoginReq) (*dto.UserAuthRes, error) {
	user, err := s.repo.GetUserByEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if !checkPasswordHash(in.Password, user.Password) {
		return nil, errIncorrectPassword
	}
	return s.JWTService.CreateTokens(user)
}

func (s *Service) Refresh(ctx context.Context, in *dto.UserRefreshReq) (*dto.UserAuthRes, error) {
	return s.JWTService.Refresh(in.RefreshToken)
}

func (s *Service) Validate(ctx context.Context, token string) (bool, error) {
	return s.JWTService.Validate(token)
}

func (s *Service) Logout(ctx context.Context) {}

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}