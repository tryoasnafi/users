package users

import (
	"time"
)

type CreateUserRequest struct {
	Username    string    `json:"username" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	DOB         time.Time `json:"dob" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"required"`
	Address     string    `json:"address" validate:"required"`
}

type UpdateUserRequest struct {
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	DOB         time.Time `json:"dob" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"required"`
	Address     string    `json:"address" validate:"required"`
}

type UserResponse struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Address     string    `json:"address"`
	DOB         time.Time `json:"dob"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
}
