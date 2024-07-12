package main

const(
	DB = "./data.db"
	DRIVER = "sqlite3"
)

type User struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}