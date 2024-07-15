package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func register(c *gin.Context){
	var newUser User
	if err := c.BindJSON(&newUser); err != nil{
		c.IndentedJSON(http.StatusBadRequest, "Invalid JSON")
		return
	}

	if newUser.Email == "" || newUser.Name == "" || newUser.Password == ""{
		c.IndentedJSON(http.StatusUnauthorized, "All fields are required.")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, "Failed to hash password")
		return
	}

	if !emailCheck(newUser.Email){
		c.IndentedJSON(http.StatusConflict, "Invalid email format.")
		return
	}

	if !usernameCheck(newUser){
		c.IndentedJSON(http.StatusConflict, "Name must contain only 5 to 15 alphanumerical character.")
		return
	}

	db, err := sql.Open(DRIVER, DB)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, "Connection to database failed.")
		return
	}
	defer db.Close()

	rows, _ := db.Query("SELECT email FROM USERS WHERE email = ?", newUser.Email)

	var emailExists string

	for rows.Next() {
		rows.Scan(&emailExists)
	}
	if emailExists != ""{
		c.IndentedJSON(http.StatusConflict, "Email adress already taken.")
		return
	}

	r := `INSERT INTO USERS VALUES(NULL,?,?,?)`
	_, err = db.Exec(r, newUser.Email, newUser.Name, hashedPassword)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, "Failed to insert user in database.")
		return
	}
	c.IndentedJSON(http.StatusCreated, newUser)
}

func emailCheck(u string) bool{
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(u)
}

func usernameCheck(u User) bool{
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9]{5,15}$`)
	return usernameRegex.MatchString(u.Name)
}