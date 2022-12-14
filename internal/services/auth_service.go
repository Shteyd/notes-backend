package services

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"github.com/Shteyd/notes-backend/internal/models"
	"github.com/Shteyd/notes-backend/internal/repository"
	"github.com/Shteyd/notes-backend/pkg/env"
	"github.com/dgrijalva/jwt-go"
)

const tokenSSL = 12 * time.Hour

type AuthService struct {
	config *env.EnvConfig
	repo   repository.UsersRepository
}

func NewAuthService(config *env.EnvConfig, repo repository.UsersRepository) *AuthService {
	return &AuthService{config: config, repo: repo}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) CreateUser(user models.InputUser) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	userId, err := s.repo.GetUserID(email, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenSSL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(s.config.AppConfig.SignInKey))
}

func (s *AuthService) ParseToken(token string) (int, error) {
	res, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.config.AppConfig.SignInKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := res.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.config.AppConfig.Salt)))
}
