package adapters

import (
	"github.com/EnterpriseGradeSystem/pkg/models"
)

// UserRepository represents the user repository
//
// @Description: UserRepository is a struct that represents the user repository.
//
// @Methods:
//   - GetAllUsers(): ([]User, error)
//
type UserRepository struct{}

// NewUserRepository returns a new user repository
//
// @Description: NewUserRepository is a function that returns a new user repository.
//
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	// Implement the logic to get all users from the database
	// For demonstration purposes, we will return a hardcoded list of users
	users := []models.User{
		{ID: "1", Name: "John Doe", Email: "john.doe@example.com"},
		{ID: "2", Name: "Jane Doe", Email: "jane.doe@example.com"},
	}
	return users, nil
}