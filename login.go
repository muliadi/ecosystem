package ecosystem

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TempPWs struct {
  Email string
  Passwrod string
  Expires time.Unix()
}

//PreLogin collects the users email and sends them the login link
func PreLogin(c *gin.Context) {

	email := c.PostForm("email")

	data := map[string]string{
		"password": "1234",
	}

	err := sendEmail(
		//Email server settings
		Config["smtpServer"],
		Config["smtpPort"],
		Config["smtpUserName"],
		Config["smtpPW"],
		[]string{email},            //Recipient
		"Your Authentication Link", //Subject
		data, //Data to include in the email
		"eco-email-auth.html") //Email template to use

	if err != nil {
		c.HTML(http.StatusOK, "eco-embedded-message.html", gin.H{
			"Message": "There was a problem sending you a message.",
		})
	} else {
		c.HTML(http.StatusOK, "eco-embedded-message.html", gin.H{
			"Message": "We have sent you an email",
		})
	}

}

//Login handles user login
func Login(c *gin.Context) {

	email := c.Param("email")
	password := c.Param("password")
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims["email"] = email
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, _ := token.SignedString([]byte("ecosystem"))

	thisUser := User{
		Email:    email,
		Password: password,
		Token:    tokenString,
	}

	//Render the login page with the token attached to the DOM
	//The login page contians a script to extract the token to local storage
	c.HTML(http.StatusOK, "eco-profile-logged-in.html", thisUser)

}
