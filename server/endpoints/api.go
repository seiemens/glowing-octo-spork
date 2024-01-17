package api

import (
	"fmt"
	"fridge/lib"
	"fridge/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"answer": lib.CreateUser(user.Username, user.Email, user.Phone, user.Password)})
}

func AuthUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
	}
	authUser := lib.AuthUser(user.Username, user.Password)
	if authUser.ID == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"answer": 404})
	} else {
		if authUser.Phone != "" {
			lib.ValidatePhoneNumber(authUser.Phone)
			c.IndentedJSON(http.StatusOK, gin.H{"answer": "2FA initialized"})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"answer": "no phone number"})
		}
	}
}
