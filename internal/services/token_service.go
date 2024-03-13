package services

import (
	"fmt"
	"time"

	"github.com/StefanWellhoner/task-manager-api/internal/config"
	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repository"
	"github.com/golang-jwt/jwt"
)

var (
	SigningKey         = []byte(config.Get().Secrets.Jwt)
	AccessTokenExpiry  = 15 * time.Minute
	RefreshTokenExpiry = 7 * 24 * time.Hour
)

type Claims struct {
	jwt.StandardClaims
	UserID    string `json:"user_id"`
	TokenType string `json:"token_type"`
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
}

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
	accessTokenClaims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpiry).Unix(),
		},
		UserID:    user.ID.String(),
		TokenType: "access",
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(SigningKey)
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(RefreshTokenExpiry).Unix(),
		},
		UserID:    user.ID.String(),
		TokenType: "refresh",
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(SigningKey)
	if err != nil {
		return nil, err
	}

	refreshTokenModel := &model.RefreshToken{
		Token:     refreshTokenString,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(RefreshTokenExpiry),
	}

	if err := s.SaveRefreshToken(refreshTokenModel); err != nil {
		return nil, err
	}

	return &TokenDetails{AccessToken: accessTokenString, RefreshToken: refreshTokenString}, nil
}

func (s *tokenService) SaveRefreshToken(token *model.RefreshToken) error {
	if err := s.tokenRepository.Create(token); err != nil {
		return err
	}
	return nil
}

func (s *tokenService) ValidateToken(token string) (*model.RefreshToken, error) {
	refreshToken, err := s.tokenRepository.FindByToken(token)
	if err != nil {
		return nil, err
	}
	if refreshToken == nil {
		return nil, fmt.Errorf("invalid token")
	}

	if time.Now().After(refreshToken.ExpiresAt) {
		return nil, fmt.Errorf("token expired")
	}

	return refreshToken, nil
}

func (s *tokenService) DeleteRefreshToken(token string) error {
	if err := s.tokenRepository.Delete(token); err != nil {
		return err
	}
	return nil
}
