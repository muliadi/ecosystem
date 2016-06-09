package ecosystem

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//User is
type User struct {
	Email string
	Cart  string
}

//UserProfilePage displays the skeleton for the profile page which is basically an Intercooler content block
//which requests profile content along with the token from local storage
func UserProfilePage(c *gin.Context) {
	c.HTML(http.StatusOK, "eco-profile-page.html", gin.H{})
}

//UserProfileContent provides the content for the Intercooler block on the profile page.
//Content depends on whether user is logged in or out, which in turn depends on the presence/validity
//of the token from localstorage sent along depending on whether the user is logged in
func UserProfileContent(c *gin.Context) {
	//Get the token from Intercooler
	tokenString := c.Query("token")
	//Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config["signingKey"]), nil
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
