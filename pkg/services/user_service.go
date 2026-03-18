package services

import (
	"github.com/EnterpriseGradeSystem/pkg/adapters"
	"github.com/EnterpriseGradeSystem/pkg/models"
)

// UserService represents the user service
//
// @Description: UserService is a struct that represents the user service.
//
// @Methods:
//   - GetAllUsers(): ([]User, error)
//
type UserService struct{}

// NewUserService returns a new user service
//
// @Description: NewUserService is a function that returns a new user service.
//
func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	userRepository := adapters.NewUserRepository()
	return userRepository.GetAllUsers()
}