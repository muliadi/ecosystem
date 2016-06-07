package ecosystem

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Login handles user login
func Login(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims["email"] = email
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, _ := token.SignedString([]byte("ecosystem"))

	log.Println("Token: ", tokenString)

	thisUser := User{
		Email:    email,
		Password: password,
		Token:    tokenString,
	}

	c.HTML(http.StatusOK, "eco-profile-logged-in.html", thisUser)

}
