package models

import "time"

type Users struct {
	Id          string
	Name    string    `json:"Name"`
	Email       string    `json:"email"`
	Timestamp   time.Time `json:"timestamp"`
}

type RegisterRequest struct {
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
