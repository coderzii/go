package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		unescapedUrl       string `json:"unescapedUrl"`
		url                string `json:"url"`
		visibleUrl         string `json:"visibleUrl"`
		cacheUrl           string `json:"cacheUrl"`
		title              string `json:"title"`
		titleNoFormatting  string `json:"titleNoFormatting"`
		content            string `json:"content"`
	}

	gResponse struct {
		ResponseData struct {
			Results []gResult `json:results`
		} `json:"responseData`
	}
)

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	// uri := "./data.json"

	resp, err := http.Get(uri)

	if err != nil {
		log.Println("Error : %s", err)

		return
	}

	var gr gResponse

	err = json.NewDecoder(resp.Body).Decode(&gr)

	if err != nil {
		log.Println("Error : %s", err)

		return
	}

	fmt.Println(gr)
}
