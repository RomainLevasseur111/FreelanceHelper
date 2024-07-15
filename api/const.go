package main

const(
	DB = "./data.db"
	DRIVER = "sqlite3"
)

type APIerror struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type User struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}