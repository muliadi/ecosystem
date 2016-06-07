package ecosystem

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//User is
type User struct {
	Email    string
	Password string
	Token    string
}

//UserProfile shows the user's profile page
func UserProfilePage(c *gin.Context) {
	thisUser := User{Email: "jon@pincas.co.uk", Password: "1234"}
	c.HTML(http.StatusOK, "eco-profile-page.html", thisUser)
}

//UserProfile shows the user's profile page
func UserProfileContent(c *gin.Context) {
	var token string //Initialising deals with token=empty case
	token = c.Query("token")
	thisUser := User{Email: "jon@pincas.co.uk", Password: "1234", Token: token}
	if token != "" {
		c.HTML(http.StatusOK, "eco-profile-logged-in.html", thisUser)
	} else {
		c.HTML(http.StatusOK, "eco-profile-logged-out.html", thisUser)
	}

}
