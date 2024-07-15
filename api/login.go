package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func login(c *gin.Context){
	session := c.MustGet("session").(*sessions.Session)
	var newLoginRequest LoginRequest
	if err := c.BindJSON(newLoginRequest); err != nil{
		c.IndentedJSON(http.StatusBadRequest, "coucou")
	}




	session.Values["authenticated"] = true
	session.Save(c.Request, c.Writer)
	c.JSON(http.StatusOK, "Login successful")
}