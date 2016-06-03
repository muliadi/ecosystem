package ecosystem

import (
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type Collection struct {
	Tag         string
	NotEmbedded bool
	Products    map[string]Product
}

func (c Collection) FilterProducts() []Product {
	//Set up a slice to hold the filtered list of products
	var filtered []Product
	//Iterate over all the products
	for _, product := range c.Products {
		//Iterate over the categories specified in the product
		//If a match is found, add the product to the returned slice
		for _, prodcat := range product.Tags {
			if prodcat == c.Tag {
				filtered = append(filtered, product)
				break
			}
		}
	}
	return filtered
}

func (c Collection) NewCollection(newTag string) Collection {
	return Collection{Products: ProductsMap, Tag: newTag, NotEmbedded: false}
}

func ShowCategoryTemplate(c *gin.Context) {
	//Filter the product list
	//filteredProducts := filterbyCategory(c.Param("cat"))
	//Does the specified template exist?

	thisCollection := Collection{Products: ProductsMap, Tag: c.Param("cat"), NotEmbedded: true}

	templateFile := "collection-" + c.Param("cat") + ".html"
	templatePath := path.Join("templates", "collection", templateFile)
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		//If it doesn't, revert to default
		c.HTML(http.StatusOK, "collection-default.html", thisCollection)
	} else {
		//If it does, use it
		c.HTML(http.StatusOK, templateFile, thisCollection)
	}

}

//// Work out the current category, filter the products, apply the correct template and render
//func ShowCategory(c *gin.Context) {
//	//Filter the product list
//	filteredProducts := filterbyCategory(c.Param("cat"))
//	//If the category exists AND has a specified template
//	thisCat, ok := CategoryMap[c.Param("cat")]
//	if  ok == true && len(thisCat.Template) > 0 {
//		//Does the specified template exist?
//		templateFile := "category-list-" + thisCat.Template + ".html"
//		templatePath := path.Join("templates", templateFile)
//		if _, err := os.Stat(templatePath); os.IsNotExist(err) {
//			//If it doesn't, revert to default
//			c.HTML(http.StatusOK, "category-list-default.html", gin.H{
//				"products": filteredProducts,
//				"category": thisCat,
//			})
//		} else {
//			//If it does, use it
//			c.HTML(http.StatusOK, templateFile, gin.H{
//				"products": filteredProducts,
//				"category": thisCat,
//			})
//		}
//	} else {
//		//For the case where either the category isn't found or doesn't have a specified template
//		//Shoud return an error if not found TODO
//		c.HTML(http.StatusOK, "category-list-default.html", gin.H{
//			"products": filteredProducts,
//			"category": thisCat,
//		})
//	}
//
//}
////Function to filter down the list of products in preparation for sending to a category view
//func filterbyCategory(filter string) []Product {
//	//Set up a slice to hold the filtered list of products
//	var filtered []Product
//	//Iterate over all the products
//	for _, product := range ProductsMap {
//		//Iterate over the categories specified in the product
//		//If a match is found, add the product to the returned slice
//		for _, prodcat := range product.Tags {
//			if prodcat == filter {
//				filtered = append(filtered, product)
//				break
//			}
//		}
//	}
//	return filtered
//
//}

//func GetCategories(api string) {
//	res, _ := http.Get(api)
//	categories, _ := ioutil.ReadAll(res.Body)
//	json.Unmarshal(categories, &CategoryMap)
//}
//
//func ReadCategories(file string) {
//	categories, err := ioutil.ReadFile(file)
//	if err == nil {
//		json.Unmarshal(categories, &CategoryMap)
//	}
//}

// type Category struct {
//	Title string
//	Slug string
//	Description string
//	Image string
//	Template string
//	Embedded []Category
//}
//
//var CategoryMap = make(map[string]Category)
