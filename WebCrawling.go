// parse_urls.go
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// Parse a complex URL
	complexUrl := "https://www.example.com/path/to/?query=123&this=that#fragment"
	parsedUrl, err := url.Parse(complexUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Print out URL pieces
	fmt.Println("Scheme: " + parsedUrl.Scheme)
	fmt.Println("Host: " + parsedUrl.Host)
	fmt.Println("Path: " + parsedUrl.Path)
	fmt.Println("Query string: " + parsedUrl.RawQuery)
	fmt.Println("Fragment: " + parsedUrl.Fragment)

}
