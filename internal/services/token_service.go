package services

import (
	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repository"
)

type TokenService interface {
	GenerateTokens(user *model.User) (*TokenDetails, error)
	SaveRefreshToken(token *model.RefreshToken) error
	ValidateToken(token string) (*model.RefreshToken, error)
	DeleteRefreshToken(token string) error
}

type tokenService struct {
	tokenRepository repositories.TokenRepository
}

func NewTokenService(tokenRepo repositories.TokenRepository) TokenService {
	return &tokenService{tokenRepository: tokenRepo}
}

func (s *tokenService) GenerateTokens(user *model.User) (*TokenDetails, error) {
	return nil, nil
}

func (s *tokenService) SaveRefreshToken(token *model.RefreshToken) error {
	return nil
}

func (s *tokenService) ValidateToken(token string) (*model.RefreshToken, error) {
	return nil, nil
}

func (s *tokenService) DeleteRefreshToken(token string) error {
	return nil
}
