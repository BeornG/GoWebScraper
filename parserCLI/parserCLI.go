package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*
mep
write a help and some kind of initial ascii what commands you can use
rewrite csv converter to use ioutils instead of os so i dont have to have a separate chmod function
more commenting
find out if its possible to convert info into plaintext and other formats
another get command for specific items
*/

func main() {
	fmt.Println("ᶘ ᵒᴥᵒᶅ workwork")
	// Get commands
	getCMD := flag.NewFlagSet("get", flag.ExitOnError)
	getALL := getCMD.Bool("all", false, "Get all the products")

	// Convert commands
	convertCMD := flag.NewFlagSet("convert", flag.ExitOnError)
	convertCSV := convertCMD.Bool("csv", false, "Convert to .csv")

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		handleGET(getCMD, getALL)
	case "convert":
		handleCONVERT(convertCMD, convertCSV)
	default:
		fmt.Println("Not a command")
	}

}

func handleGET(getCMD *flag.FlagSet, all *bool) {
	getCMD.Parse(os.Args[2:])

	switch {
	case *all == false:
		fmt.Print("Need to specify -all")
		getCMD.PrintDefaults()
		os.Exit(1)
	case *all == true:
		jsonParseAllToCommandline()
	default:
		return
	}
}
func handleCONVERT(convertCMD *flag.FlagSet, csv *bool) {
	convertCMD.Parse(os.Args[2:])
	switch {
	case *csv == false:
		fmt.Print("Need to specify -csv")
		convertCMD.PrintDefaults()
		os.Exit(1)
	case *csv == true:
		csvConvert()
	default:
		return
	}

}

func csvConvert() {
	file, err := ioutil.ReadFile("emmausproducts.json")

	if err != nil {
		log.Fatal(err)
	}
	var data []Product
	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		log.Fatal(err)
	}

	csvFile, err := os.Create("emmausproducts.csv")

	if err != nil {
		log.Fatal(err)
	}
	err = os.Chmod("emmausproducts.csv", 0700)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	// iterate and append rows
	for _, strings := range data {
		var row []string
		row = append(row, strings.NameAndPrice)
		row = append(row, strings.URL)
		writer.Write(row)
	}
	writer.Flush()
}

func jsonParseAllToCommandline() {
	File, err := os.Open("emmausproducts.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("You have opened the file!")
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
