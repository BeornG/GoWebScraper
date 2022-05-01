package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/gocolly/colly"
)

/*
mep
visiting - works
response - works
error - works
file writer - works
element iterator - works ish, flexbox difficult to navigate so cant separate name and price
					finding image paths suck when embedded in styles feels like i would have
					to setup a separate crawler on the product page
*/
func main() {
	crawler() // scrapes the target and generates a json with desired info

}

func crawler() {
	// Declaring c as a collector and specifying the domain we want to extract info from
	c := colly.NewCollector(
		colly.AllowedDomains("emmaus.ax", "www.emmaus.ax"),
	)
	c.SetRequestTimeout(120 * time.Second) // timeout

	Products := make([]Product, 0) // making an empty slice to hold the info we extract

	// here we start pulling info iterating over the elements and appending the info to our empty slice
	c.OnHTML("a.productlink", func(e *colly.HTMLElement) {
		e.ForEach("div.prodlistinfo", func(i int, h *colly.HTMLElement) {
			item := Product{}
			item.NameAndPrice = h.Text
			item.URL = "https://www.emmaus.ax" + e.Attr("href")

			Products = append(Products, item)
		})

	})
	// Visiting message
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	// Response message
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	// Error message
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})
	// On succesful extraction prints message and generates a json file
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Done!", r.Request.URL)
		writeJSON(Products)
	})
	c.Visit("http://www.emmaus.ax/webbshop")
}

// Writes the json file
func writeJSON(data []Product) {
	file, err := json.MarshalIndent(data, "", "     ") //formats the input
	if err != nil {
		log.Println("Error can't create json file")
		return
	}

	_ = ioutil.WriteFile("emmausproducts.json", file, 0644) // file name, 0644 sets the permissions
}

// struct of the items we want to extract
type Product struct {
	NameAndPrice string `json:"NameAndPrice"`
	URL          string `json:"URL"`
	//Image        string `json:"Image"`
}
