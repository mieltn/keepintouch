package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mieltn/keepintouch/internal/config"
	"github.com/mieltn/keepintouch/internal/dto"
)

const (
	tokenTypeAccess = "access"
	tokenTypeRefresh = "refresh"
)

var (
	errInvalidClaims = errors.New("claims type assertion failed")
	errInvalidToken = errors.New("token is invalid")
	errUnexpectedSigningMethod = errors.New("unexpected signing method")
)

type Service struct {
	cfg *config.Config
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		cfg: cfg,
	}
}

func (s *Service) CreateTokens(user *dto.User) (*dto.UserAuthRes, error) {
	access, err := s.generateToken(
		tokenTypeAccess,
		user.Id,
		user.Email,
		user.Username,
		s.cfg.JWTAccessLifetimeMins,
	)
	if err != nil {
		return nil, err
	}

	refresh, err := s.generateToken(
		tokenTypeRefresh,
		user.Id,
		user.Email,
		user.Username,
		s.cfg.JWTRefreshLifetimeMins,
	)
	if err != nil {
		return nil, err
	}

	return &dto.UserAuthRes{
		AccessToken: access,
		RefreshToken: refresh,
		Id: user.Id,
		Username: user.Username,
		Email: user.Email,
	}, nil
}

func (s *Service) Refresh(refreshToken string) (*dto.UserAuthRes, error) {
	_, claims, err := s.decodeToken(refreshToken)
	if err != nil {
		return nil, err
	}
	return s.CreateTokens(&dto.User{
		Id: claims["user_id"].(string),
		Username: claims["username"].(string),
		Email: claims["user_email"].(string),
	})
}

func (s *Service) Validate(accessToken string) (bool, error) {
	_, _, err := s.decodeToken(accessToken)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) generateToken(tokenType, id, email, username string, lifetime int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["typ"] = tokenType
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(lifetime)).Unix()
	claims["user_id"] = id
	claims["username"] = username
	claims["user_email"] = email

	tokenString, err := token.SignedString([]byte(s.cfg.JWTSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *Service) decodeToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errUnexpectedSigningMethod
		}

		return []byte(s.cfg.JWTSecretKey), nil
	})

	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, errInvalidClaims
	}

	if !token.Valid {
		return nil, nil, errInvalidToken
	}

	return token, claims, nil
}
