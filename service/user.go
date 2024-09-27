package service

import "time"

type UserResponse struct {
	Code      string    `json:"code"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_ at"`
}
type TokenResponse struct {
	Token string `json:"token"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	CreateUser(UserRequest) error
	UserLogin(UserRequest) (*TokenResponse, error)
}
