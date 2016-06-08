package ecosystem

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//User is
type User struct {
	Email    string
	Password string
	Token    string
}

//UserProfile displays the skeleton for the profile page
func UserProfilePage(c *gin.Context) {
	c.HTML(http.StatusOK, "eco-profile-page.html", gin.H{})
}

//UserProfileContent fills in the profile page depending on whether the user is logged on or not and all sub-cases
func UserProfileContent(c *gin.Context) {
	//Get the token from Intercooler
	tokenString := c.Query("token")
	//Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("ecosystem"), nil
	})
	//Validate
	if err == nil && token.Valid && token.Claims["email"].(string) != "" {
		thisUser := User{Email: token.Claims["email"].(string)}
		c.HTML(http.StatusOK, "eco-profile-logged-in.html", thisUser)
	} else {
		c.HTML(http.StatusOK, "eco-profile-logged-out.html", gin.H{
			"Message": "You're not currently logged into British Bins on this computer.  If you'd like to log in, just enter your email below",
		})
	}

}
