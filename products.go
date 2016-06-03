package ecosystem

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"path"
	"os"
)

type Product struct {
	Sku string
	Weight float64
	Salestaxcode string
	Price float64
	Stock int
	Slug string
	Image string
	Tags []string
	Template string
	Pagetitle string
	Title string
	Description string
	Custom map[string]string
}

var ProductsMap = make(map[string]Product)

func GetProducts(api string) {
	res, _ := http.Get(api)
	products, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(products, &ProductsMap)
}

func ReadProducts(file string) {
	products, err := ioutil.ReadFile(file)
	if err == nil {
		json.Unmarshal(products, &ProductsMap)
	}
	//What do do as a last resort if products cannot be read? TODO
}

func ShowProduct(c *gin.Context){

	//If the product exists AND has a specified template
	thisProduct, ok := ProductsMap[c.Param("product")]
	if  ok == true && len(thisProduct.Template) > 0 {
		//Does the specified template exist?
		templateFile := "drilldown-" + thisProduct.Template + ".html"
		templatePath := path.Join("templates", templateFile)
		if _, err := os.Stat(templatePath); os.IsNotExist(err) {
			//If it doesn't, revert to default
			c.HTML(http.StatusOK, "drilldown-default.html", thisProduct)
		} else {
			//If it does, use it
			c.HTML(http.StatusOK, templateFile, thisProduct)
		}
	} else {
		//For the case where either the product isn't found or doesn't have a specified template
		//Shoud return an error if not found TODO
		c.HTML(http.StatusOK, "drilldown-default.html", thisProduct)
	}

}