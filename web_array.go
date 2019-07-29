package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	GetURL()
}

func GetURL() string {
	fmt.Print("What is the URL? ")
	var URL string
	fmt.Scan(&URL)
	parsedURL, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}

	type Data struct {
		Scheme   string
		Host     string
		Path     string
		Query    string
		Fragment string
	}

	info := Data{
		Scheme:   parsedURL.Scheme,
		Host:     parsedURL.Host,
		Path:     parsedURL.Path,
		Query:    parsedURL.RawQuery,
		Fragment: parsedURL.Fragment,
	}

	fmt.Println("The scheme is: ", info.Scheme)
	fmt.Println("The host is: ", info.Host)
	fmt.Println("The Path is: ", info.Path)
	fmt.Println("The Query is: ", info.Query)
	fmt.Println("The Fragment is: ", info.Fragment)
	return URL

}
