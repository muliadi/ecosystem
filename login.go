package ecosystem

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Login is called by Intercooler upon submission of login form
func Login(c *gin.Context) {

	//Retrieve the email from the posted form
	email := c.PostForm("email")

	//Create a temporary, one-off password consisting of 10 random characters
	pw := randomString(10)

	//Sets the email/temp pw pair in the cache
	EmailPWCache.Set(email, pw)

	//Set up the data map to go to the email sending function
	data := map[string]string{
		"password": pw,
	}

	//Send the email
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

//Authorise fires when the customer visits the magic link sent by email.  It is responsible for
//checking the email/temp pw combo and issuing a long-lived JWT with the email claim
func Authorise(c *gin.Context) {

	//Retrieve the email and password parameters from the URL
	email := strings.Replace(c.Query("email"), "%40", "@", 1) //Sanitise the email from the URL (change %40 to @)
	pw := c.Query("password")

	//Perform validation of the temp pw
	value, exists := EmailPWCache.Get(email)
	//First check if there is an enrty in the cache for that email
	if !exists {
		//IF there is no matching email in the cache - show an error
		c.HTML(http.StatusOK, "eco-message.html", gin.H{
			"Message": "Authorisation request timed out - please try again",
		})
	} else {
		//If the password is wrong
		if pw != value {
			//If the email/pw combo exists and the pw is write, proceed to log the user on
			c.HTML(http.StatusOK, "eco-message.html", gin.H{
				"Message": "Incorrect authorisation credentials - please try again",
			})
		} else {
			//Combo exists in cache and pw is correct - log the user in
			//First delete the email/pw combo in the cache so it can't be used again
			EmailPWCache.Remove(email)

			//Create a new token
			token := jwt.New(jwt.SigningMethodHS256)
			token.Claims["email"] = email
			token.Claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
			// Sign and get the complete encoded token as a string
			tokenString, _ := token.SignedString([]byte(Config["signingKey"]))

			//Render the login success page with the token attached to the DOM
			//The login OK page contians a script to extract the token to local storage
			c.HTML(http.StatusOK, "eco-login-ok.html", gin.H{
				"Email": email,
				"Token": tokenString,
			})

		}
	}

}

func randomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
