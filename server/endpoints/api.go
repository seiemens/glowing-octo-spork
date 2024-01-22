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
			cookie := lib.CreateSessionToken(authUser.ID)
			c.SetCookie("user", cookie.Cookie, 3600, "/", "localhost", false, true)
			c.IndentedJSON(http.StatusOK, gin.H{"cookie": cookie.Cookie, "created": cookie.CreatedAt, "expire": cookie.ExpireOn})
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
		c.SetCookie("user", cookie.Cookie, 3600, "/", "localhost", false, true)

		c.IndentedJSON(http.StatusOK, gin.H{"cookie": cookie.Cookie, "created": cookie.CreatedAt, "expire": cookie.ExpireOn})
	} else {
		c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "wrong token"})
	}

}

func GetUserData(c *gin.Context) {
	var user models.User

	cookie := c.Request.Header.Get("Cookie")
	// cookie, err := c.Cookie("user")
	// if err != nil {
	//     c.String(http.StatusNotFound, "Cookie not found")
	//     return
	// }
	ok, userID := lib.VerifySessionToken(cookie)
	if ok {
		users := lib.GetUserByKey("id", userID)
		user = users[0]
		c.IndentedJSON(http.StatusOK, gin.H{"id": user.ID, "username": user.Username, "email": user.Email, "phone": user.Phone, "isAdmin": user.Admin})
	} else {
		c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "unauthorized"})
	}
}

func CreateNote(c *gin.Context) {
	cookie := c.Request.Header.Get("Cookie")
	// cookie, err := c.Cookie("user")
	// if err != nil {
	//     c.String(http.StatusNotFound, "Cookie not found")
	//     return
	// }
	ok, userID := lib.VerifySessionToken(cookie)
	if ok {
		var note models.Note
		err := c.BindJSON(&note)
		if err != nil {
			fmt.Println(err)
		}
		lib.CreateNote(userID, note.Title, note.Content)
		c.IndentedJSON(http.StatusOK, gin.H{"answer": "note created successfully"})
	} else {
		c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "unauthorized"})
	}
}

func GetNotes(c *gin.Context) {
	var mode models.Mode
	err := c.BindJSON(&mode)
	if err != nil {
		fmt.Println(err)
	}

	if mode.Mode == "public" {
		notes := lib.GetNoteByStatus("status", models.Published)
		c.IndentedJSON(http.StatusOK, gin.H{"answer": notes})
	} else if mode.Mode == "user" {
		cookie := c.Request.Header.Get("Cookie")
		// cookie, err := c.Cookie("user")
		// if err != nil {
		//     c.String(http.StatusNotFound, "Cookie not found")
		//     return
		// }
		ok, userID := lib.VerifySessionToken(cookie)
		if ok {
			notes := lib.GetNoteByKey("userid", userID)
			c.IndentedJSON(http.StatusOK, gin.H{"answer": notes})
		} else {
			c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "unauthorized"})
		}
	} else if mode.Mode == "admin" {
		cookie := c.Request.Header.Get("Cookie")
		// cookie, err := c.Cookie("user")
		// if err != nil {
		//     c.String(http.StatusNotFound, "Cookie not found")
		//     return
		// }
		ok, userID := lib.VerifySessionToken(cookie)
		if ok && lib.IsAdmin(userID) {
			notes := lib.GetAllNotes()
			c.IndentedJSON(http.StatusOK, gin.H{"answer": notes})
		} else {
			c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "unauthorized"})
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"answer": "try again"})
	}

}