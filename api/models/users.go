package models

import "time"

type User struct {
	Id          string
	Name    string    `json:"name"`
	Email       string    `json:"email"`
	Timestamp   time.Time `json:"timestamp"`
}

type RegisterRequest struct {
	Name    string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
