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

func LoginUser(c *gin.Context) {
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
			c.IndentedJSON(http.StatusOK, gin.H{"answer": "authenticated"})
		}
	}
}

func AuthUser(c *gin.Context) {
	cookie, err := c.Cookie("user")
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}

	ok, _ := lib.VerifySessionToken(cookie)
	if ok {
		c.IndentedJSON(http.StatusOK, gin.H{"answer": "authenticated"})
	} else {
		c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "unauthorized"})
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
		lib.DeleteSMSToken(sms.ProcessID) // delete token after successful
		cookie := lib.CreateSessionToken(userId)
		c.SetCookie("user", cookie.Cookie, 3600, "/", "localhost", false, true)

		c.IndentedJSON(http.StatusOK, gin.H{"answer": "authenticated"})
	} else {
		c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "wrong token"})
	}

}

func GetUserData(c *gin.Context) {
	var user models.User

	cookie, err := c.Cookie("user")
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}
	ok, userID := lib.VerifySessionToken(cookie)
	if ok {
		users := lib.GetUserByKey("id", userID)
		user = users[0]
		c.IndentedJSON(http.StatusOK, gin.H{"id": user.ID, "username": user.Username, "email": user.Email, "phone": user.Phone, "isAdmin": user.Admin, "apiKey": user.ApiKey})
	} else {
		c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "unauthorized"})
	}
}

func CreateNote(c *gin.Context) {

	cookie, err := c.Cookie("user")
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}
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

		cookie, err := c.Cookie("user")
		if err != nil {
			c.String(http.StatusNotFound, "Cookie not found")
			return
		}
		ok, userID := lib.VerifySessionToken(cookie)

		if ok && lib.IsAdmin(userID) {
			notes := lib.GetAllNotes()
			c.IndentedJSON(http.StatusOK, gin.H{"answer": notes})
		} else if ok {
			notes := lib.GetNoteByKey("userid", userID)
			c.IndentedJSON(http.StatusOK, gin.H{"answer": notes})
		} else {
			c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "unauthorized"})
		}
	}

}

func AddComment(c *gin.Context) {
	cookie, err := c.Cookie("user")
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}
	ok, userID := lib.VerifySessionToken(cookie)
	if ok {
		var note models.Comment
		err := c.BindJSON(&note)
		if err != nil {
			fmt.Println(err)
		}
		lib.AddCommentToPost(cookie, note.PostID, note.Content, userID)
		c.IndentedJSON(http.StatusOK, gin.H{"answer": "comment created successfully"})
	} else {
		c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "unauthorized"})
	}
}

func ChangePhone(c *gin.Context) {
	var user models.User
	cookie, err := c.Cookie("user")
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}
	ok, userID := lib.VerifySessionToken(cookie)
	if ok {
		err := c.BindJSON(&user)
		if err != nil {
			fmt.Println(err)
		}
		lib.ChangePhone(user.Phone, userID)
		c.IndentedJSON(http.StatusOK, gin.H{"answer": "phone changed successfully"})
	} else {
		c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "unauthorized"})

	}
}

func IsUserAdmin(c *gin.Context) {
	cookie, err := c.Cookie("user")
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}
	ok, userID := lib.VerifySessionToken(cookie)
	if ok {
		if lib.IsAdmin(userID) {
			c.IndentedJSON(http.StatusAccepted, gin.H{"answer": "is a admin"})
		} else {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"answer": "unauthorized"})

		}
	} else {
		c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "unauthorized"})
	}
}

func Logout(c *gin.Context) {
	cookie, err := c.Cookie("user")
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}
	lib.Logout(cookie)
	c.SetCookie("user", "", -1, "/", "localhost", false, true)
	c.String(http.StatusOK, "Cookie has been deleted")
}

func ChangeVisibility(c *gin.Context) {
	var note models.Note
	cookie, err := c.Cookie("user")
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}
	ok, userID := lib.VerifySessionToken(cookie)
	if ok {
		err := c.BindJSON(&note)
		if err != nil {
			fmt.Println(err)
		}
		if userID == note.UserID || lib.IsAdmin(userID) {
			lib.ChangeVisibility(note.ID, note.Status)
			c.IndentedJSON(http.StatusOK, gin.H{"answer": "status changed successfully"})
		} else {
			c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "unauthorized"})
		}
	} else {
		c.IndentedJSON(http.StatusTeapot, gin.H{"answer": "unauthorized"})
	}
}

func GetPublicPosts(c *gin.Context) {
	APIKey := c.Request.Header.Get("X-API-Key")
	ok := lib.ValidateAPIKey(APIKey)
	fmt.Println(APIKey)
	if ok {
		posts := lib.GetNoteByStatus("status", models.Published)
		c.IndentedJSON(http.StatusOK, gin.H{"answer": posts})
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"answer": "unauthorized"})
	}
}
