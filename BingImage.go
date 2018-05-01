package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Listen on port and reply with bing image url
	r := gin.Default()
	r.GET("/bingImgURL/", func(c *gin.Context) {
		c.String(200, getImageURL())
	})

	port := os.Args[1:][0]
	r.Run(":" + port) // listen on localhost:port
}

func getImageURL() string {
	bingHomeURL := "https://www.bing.com"
	bingImageReqURL := "https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=en-US"

	response, err := http.Get(bingImageReqURL)
	if err != nil {
		fmt.Println("Could not get the Bing image URL:", err)
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		var result map[string]interface{}
		json.Unmarshal(body, &result)
		images := result["images"].([]interface{})
		url := images[0].(map[string]interface{})["url"].(string)
		return bingHomeURL + url
	}

	return ""
}
