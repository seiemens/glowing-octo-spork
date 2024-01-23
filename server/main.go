package main

import (
	api "fridge/endpoints"
	"fridge/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	lib.ConnectToDb()
	//lib.CreateSessionToken("5Ygiw8N8")
	router := gin.New()
	router.Use(CORSMiddleware())
	lib.ValidateAPIKey("asdfasdfasdf")
	router.POST("/api/user/register", api.CreateUser)
	router.POST("/api/user/login", api.LoginUser)
	router.GET("/api/user/auth", api.AuthUser)
	router.POST("/api/user/sms", api.VerifySMS)
	router.GET("/api/user/info", api.GetUserData)
	router.POST("/api/user/telefon", api.ChangePhone)
	router.GET("/api/user/isadmin", api.IsUserAdmin)
	router.GET("/api/user/logout", api.Logout)
	router.POST("/api/user/naughty", api.NaughtyUser)

	router.POST("/api/posts/create", api.CreateNote)
	router.POST("/api/posts/visibility", api.ChangeVisibility)
	router.POST("/api/posts/get", api.GetNotes)
	router.POST("/api/posts/comment", api.AddComment)

	router.GET("/api/posts", api.GetPublicPosts) // only accessible with user api key

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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
