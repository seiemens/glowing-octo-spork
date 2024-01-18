package main

import (
	api "fridge/endpoints"
	"fridge/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	lib.ConnectToDb()

	lib.CreateSMSToken("41788461103")
	router := gin.New()
	router.Use(CORSMiddleware())

	router.POST("/api/user/register", api.CreateUser)
	router.POST("/api/user/auth", api.AuthUser)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
