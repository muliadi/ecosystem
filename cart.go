`package ecosystem

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

//CartLine represents a single line in the cart
type CartLine struct {
	ID          int
	LastUpdated int64
	Cart        string
	Sku         string
	Price       float64
	Qty         int
	Img         string
	Link        string
	Title       string
}

//LineTotal multiplies price by quantity for the CartLine
func (line CartLine) LineTotal() float64 {
	return float64(line.Qty) * line.Price
}

//Cart reprsents the overall cart
type Cart struct {
	ID          string
	LastUpdated int64
	Lines       []CartLine
}

//CartTotal sums each LineTotal to reach a cart subtotal
func (cart Cart) CartTotal() float64 {
	var subtotal float64
	for _, line := range cart.Lines {
		subtotal += line.LineTotal()
	}
	return subtotal
}

//fetchCart retrieves cartlines from the database, creates and returns a Cart
func fetchCart(cartid string) (cart Cart, err error) {
	//If the cart doesn't exist, an empty cart will be returned, NOT an error
	//and error will only be returned if there is a problem accessing the database
	_, err = Dbmap.Select(&cart.Lines, "select * from cartlines where Cart=? order by ID desc", cartid)
	return
}

//ShowCart displays the content of the cart
func ShowCart(c *gin.Context) {
	cart, err := fetchCart(c.Query("cart"))
	if err != nil {
		//If the database is out of reach for some reason, display an error message
		c.String(http.StatusServiceUnavailable, "The shopping cart service is currently unavailable.  Please try again later, or contact us")
	} else {
		//Otherwise display the cart even if it is empty
		c.HTML(http.StatusOK, "eco-cart.html", cart)
	}
}

//AddToCart responds to a form input by adding the product to the cart
func AddToCart(c *gin.Context) {
	crt := c.PostForm("cart")
	//Reference the product to be added
	newProduct := ProductsMap[c.PostForm("product")]
	newQty, _ := strconv.Atoi(c.PostForm("qty"))
	//Create the new cart line
	newCartLine := &CartLine{
		Cart:        crt,
		LastUpdated: time.Now().Unix(),
		Sku:         newProduct.Sku,
		Price:       newProduct.Price,
		Qty:         newQty,
		Img:         newProduct.Image,
		Link:        newProduct.Slug,
		Title:       newProduct.Pagetitle,
	}
	//Insert the cartline and return either OK or error
	err := Dbmap.Insert(newCartLine)
	if err != nil {
		c.String(http.StatusServiceUnavailable, "")
	} else {
		c.String(http.StatusOK, "")
	}
}

//UpdateQty resets the quantity according to user input.  If the quantity is zero, deletes the cartlines
//Delete functionality uses UpdateQty via a PUT request as Gin appears to NOT make query parameters available on DELETE requests, meaning that
//we are unable to access the 'cart' identifier with DELETE requests
func UpdateQty(c *gin.Context) {
	lineid, _ := strconv.Atoi(c.Param("lineid"))
	cartid := c.PostForm("cart")
	newQty, _ := strconv.Atoi(c.PostForm("qty"))
	var cartLine CartLine

	err := Dbmap.SelectOne(&cartLine, "select * from cartlines where ID = :line and Cart = :cart", map[string]interface{}{
		"cart": cartid,
		"line": lineid,
	})

	if err != nil { //If the line cannot be found, return an error straight away
		c.String(http.StatusServiceUnavailable, "")
	} else { //else proceed with the update
		if newQty == 0 { //If the quantity is being set to 0 then delete the line
			_, err := Dbmap.Delete(&cartLine)
			if err != nil {
				c.String(http.StatusServiceUnavailable, "")
			} else {
				c.String(http.StatusOK, "")
			}
		} else { //else adjust the quantity and update the database
			cartLine.Qty = newQty
			_, err := Dbmap.Update(&cartLine)
			if err != nil {
				c.String(http.StatusServiceUnavailable, "")
			} else {
				c.String(http.StatusOK, "")
			}
		}
	}

}

//NewCart returns a UUID to be used as the cart identifer (e.g. store in Local Storage on client)
func NewCart(c *gin.Context) {
	u1 := uuid.NewV1()
	c.String(http.StatusOK, u1.String())
}
