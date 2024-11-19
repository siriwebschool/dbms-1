package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var searchKey string

func init() {
	fmt.Println("Parsing the search flag")

	flag.StringVar(&searchKey, "search", "", "provide the primary key of the record")
	flag.Parse()

	if searchKey == "" {
		log.Fatal("Provide the search flag")
	}
}

func main() {
	fmt.Println("reading the index file")
	data, err := os.ReadFile("index/index.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("creating map for the index")
	var indexes map[string]string
	err = json.Unmarshal(data, &indexes) // O(n) n - number of indexes
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON: %v", err)
	}

	fileName := fmt.Sprint("page/", indexes[searchKey], ".txt") // `page/${index[searchKey].txt}`

	fmt.Println("Reading the appropriate page file")
	fmt.Println(fileName)
	page, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("there is no such key")
	}

	fmt.Println("creating slice containing the records in the page")
	pageSlice := strings.SplitN(string(page), "\n", 10) // O(n + m) m - number of records in the page

	fmt.Println("looping through the slice to find the correct record")
	for i := 0; i < len(pageSlice); i++ {
		id, _, _ := strings.Cut(pageSlice[i], ",")
		if id == searchKey {
			fmt.Println(pageSlice[i])
		}
	}

}
