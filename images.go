package ecosystem

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"net/http"
	"path"
	"strconv"
	"github.com/gin-gonic/gin"
)


//Deal with images without width specified and othe formats TODO
func ServeImage(c *gin.Context) {

	//Get the params from the URL
	width := c.DefaultQuery("width", Config["DefaultWidth"])

	//Create the composite target filename
	targetImageFileName := path.Join("public/img", width + "w", c.Param("image"))

	//If the file doesn't exist we make it and store it
	if _, err := os.Stat(targetImageFileName); os.IsNotExist(err) {

		remoteImagePath := Config["Imgsrc"] + c.Param("image")
		remoteImage, err := http.Get(remoteImagePath)
		// Try to decode jpeg into image.Image
		img, err := jpeg.Decode(remoteImage.Body)
		if err != nil {
			log.Println("Couldn't decode")
		} else {
			// resize
			w64, err := strconv.ParseUint(width, 10, 64)
			w := uint(w64)
			m := resize.Resize(w, 0, img, resize.Lanczos3)

			// save
			os.Mkdir(path.Join("public/img", width + "w"), 0777)
			out, err := os.Create(targetImageFileName)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			// write new image to file
			jpeg.Encode(out, m, nil)
		}
	}

	//Now try to serve it
	c.File(targetImageFileName)

}