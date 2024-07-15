package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

func login(c *gin.Context){
	session := c.MustGet("session").(*sessions.Session)
	var newLoginRequest LoginRequest
	if err := c.BindJSON(&newLoginRequest); err != nil{
		c.IndentedJSON(http.StatusBadRequest, "Invalid JSON")
		return
	}

	if newLoginRequest.Email == "" || newLoginRequest.Password == ""{
		c.IndentedJSON(http.StatusUnauthorized, "All fields are required.")
		return
	}

	if !emailCheck(newLoginRequest.Email){
		c.IndentedJSON(http.StatusConflict, "Invalid email format.")
		return
	}

	
	db, err := sql.Open(DRIVER, DB)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, "Connection to database failed.")
		return
	}
	defer db.Close()

	r := "SELECT password FROM USERS WHERE email = ?"
	row := db.QueryRow(r, newLoginRequest.Email)
	var hashedpwd string
	err = row.Scan(&hashedpwd)
	if err != nil{
		fmt.Println(err)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedpwd), []byte(newLoginRequest.Password))
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusConflict, "Wrong email or password.")
		return
	}

	session.Values["authenticated"] = true
	session.Save(c.Request, c.Writer)
	c.JSON(http.StatusOK, "Login successful")
}