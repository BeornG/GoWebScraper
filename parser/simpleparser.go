package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	File, err := os.Open("emmausproducts.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("You have opened emmausproducts.json!")
	fmt.Println()
	// If the file is not empty defer the closing so we can parse it later
	if File != nil {
		defer File.Close()
	}

	// Reads the file as a byte array.
	bytes, readERR := ioutil.ReadAll(File)
	if readERR != nil {
		log.Fatal(readERR)
	}
	// Make an array to store info
	var Items []Product

	jsonERR := json.Unmarshal(bytes, &Items) // Slap the bytes into our array
	if jsonERR != nil {
		log.Fatal(jsonERR)
	}
	// Iterate over the array parsing the info
	for i := 0; i < len(Items); i++ {
		fmt.Print("Name and price: ")
		fmt.Printf("%v\n", Items[i].NameAndPrice)
		fmt.Print("URL: ")
		fmt.Printf("%v\n", Items[i].URL)
		fmt.Println()
	}

}

type Product struct {
	NameAndPrice string `json:"NameAndPrice"`
	URL          string `json:"URL"`
}
