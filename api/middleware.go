package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
    key   = []byte("super-secret-key")
    store = sessions.NewCookieStore(key)
)

func corsMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
	}
}

func sessionsMiddleware() gin.HandlerFunc{
	return func (c *gin.Context){
		session, _:= store.Get(c.Request, "session-name")
		c.Set("session", session)
		c.Next()
	}
}