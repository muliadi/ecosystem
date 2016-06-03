package ecosystem

import (
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func ShowTemplate(c *gin.Context) {

	templateFile := "test.html"
	templatePath := path.Join("templates", "layout", templateFile)
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		//If it doesn't, revert to default
		c.HTML(http.StatusOK, "collection-default.html", gin.H{
			"title": "Main website",
		})
	} else {
		//If it does, use it
		c.HTML(http.StatusOK, templateFile, gin.H{
			"title": "TEST",
		})
	}

}
