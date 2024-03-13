package services

import (
	"errors"

	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repository"
)

type UserService interface {
	Create(user *model.User) error
	Authenticate(email string, password string) (*TokenDetails, error)
}

type userService struct {
	userRepository repositories.UserRepository
	tokenService   TokenService
}

func NewUserService(userRepo repositories.UserRepository, tokenService TokenService) UserService {
	return &userService{userRepository: userRepo, tokenService: tokenService}
}

func (s *userService) Create(user *model.User) error {
	existingUser, _ := s.userRepository.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	return s.userRepository.Create(user)
}

func (s *userService) Authenticate(email, password string) (*TokenDetails, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := user.VerifyPassword(password); err != nil {
		return nil, errors.New("invalid login information")
	}

	tokenDetails, err := s.tokenService.GenerateTokens(user)
	if err != nil {
		return nil, err
	}

	return tokenDetails, nil
}
