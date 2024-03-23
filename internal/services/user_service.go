package services

import (
	"net/http"
	"time"

	"github.com/StefanWellhoner/task-manager-api/internal/errors"

	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	repositories "github.com/StefanWellhoner/task-manager-api/internal/repository"
)

type UserService interface {
	Create(user *model.User) error
	Authenticate(email string, password string) (*TokenDetails, error)
	GetUserByID(id string) (*model.User, error)
	ChangePassword(userID, oldPassword, newPassword string) error
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
		return errors.NewServiceError(errors.ConflictError, "email already in use", http.StatusConflict)
	}

	if err := model.ValidatePassword(user.PasswordHash); err != nil {
		return errors.NewServiceError(errors.ValidationError, err.Error(), http.StatusBadRequest)
	}

	hashedPassword, err := model.HashPassword(user.PasswordHash)
	if err != nil {
		return errors.NewServiceError(errors.InternalError, "Something went wrong", http.StatusInternalServerError)
	}

	user.PasswordHash = hashedPassword
	return s.userRepository.Create(user)
}

func (s *userService) Authenticate(email, password string) (*TokenDetails, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := user.VerifyPassword(password); err != nil {
		return nil, errors.NewServiceError(errors.ValidationError, "invalid email or password", http.StatusUnauthorized)
	}

	tokenDetails, err := s.tokenService.GenerateTokens(user)
	if err != nil {
		return nil, err
	}

	user.LastLogin = time.Now()
	if err := s.userRepository.Update(user); err != nil {
		return nil, errors.NewServiceError(errors.InternalError, "Something went wrong", http.StatusInternalServerError)
	}

	return tokenDetails, nil
}

func (s *userService) GetUserByID(userID string) (*model.User, error) {
	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return nil, errors.NewServiceError(errors.NotFoundError, "User not found", http.StatusNotFound)
	}
	return user, nil
}

func (s *userService) ChangePassword(userID, oldPassword, newPassword string) error {
	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return errors.NewServiceError(errors.NotFoundError, "User not found", http.StatusNotFound)
	}

	if oldPassword == newPassword {
		return errors.NewServiceError(errors.ValidationError, "New password must be different from old password", http.StatusBadRequest)
	}

	if err := user.VerifyPassword(oldPassword); err != nil {
		return errors.NewServiceError(errors.ValidationError, "Invalid old password", http.StatusBadRequest)
	}

	if err := model.ValidatePassword(newPassword); err != nil {
		return errors.NewServiceError(errors.ValidationError, err.Error(), http.StatusBadRequest)
	}

	hashedPassword, err := model.HashPassword(newPassword)
	if err != nil {
		return errors.NewServiceError(errors.InternalError, "Something went wrong", http.StatusInternalServerError)
	}

	user.PasswordHash = hashedPassword
	user.PasswordLastChanged = time.Now()

	return s.userRepository.Update(user)
}
