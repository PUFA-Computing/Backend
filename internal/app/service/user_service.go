package service

import (
	"Backend/internal/app/domain"
	"Backend/internal/app/repository"
	"Backend/internal/utils"
	"github.com/gocql/gocql"
)

type AuthResponse struct {
	User  *domain.User `json:"user"`
	Token string       `json:"token"`
}

type UserServices interface {
	RegisterUser(user *domain.User) error
	AuthenticateUser(email, password string) (*AuthResponse, error)
}

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) RegisterUser(user *domain.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	user.ID = gocql.TimeUUID()
	return s.userRepository.RegisterUser(user)
}

func (s *UserService) AuthenticateUser(email, password string) (*AuthResponse, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := utils.ComparePassword(user.Password, password); err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWTToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	response := &AuthResponse{
		User:  user,
		Token: token,
	}

	return response, nil
}
