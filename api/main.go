package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

    createDB()
    router := gin.Default()
    router.Use(corsMiddleware())
    router.Use(sessionsMiddleware())
    router.POST("/api/register", register)
    router.POST("api/login", login)
    //router.GET("api/user/{id}", getUserById)

    router.Run("localhost:8040")
}
