package models

// User represents a user
//
// @Description: User is a struct that represents a user.
// @Fields:
//   - ID: string
//   - Name: string
//   - Email: string
//
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}