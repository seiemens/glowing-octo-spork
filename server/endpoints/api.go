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
	lib.CreateUser(user.Username, user.Email, user.Phone, user.Password)
	c.IndentedJSON(http.StatusOK, gin.H{"answer": "user created successfully"})
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
			pid := lib.CreateSMSToken(authUser.Phone, authUser.ID)
			c.IndentedJSON(http.StatusOK, gin.H{"id": pid, "phone": authUser.Phone})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"answer": "no phone number"})
		}
	}
}

func VerifySMS(c *gin.Context) {
	var sms models.Smstoken
	err := c.BindJSON(&sms)
	if err != nil {
		fmt.Println(err)
	}
	isOk, userId := lib.VerifySMSToken(sms.ProcessID, sms.AccessToken)
	if isOk {
		cookie := lib.CreateSessionToken(userId)
		c.IndentedJSON(http.StatusOK, gin.H{"cookie": cookie})
	} else {
		c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "wrong token"})
	}

}
