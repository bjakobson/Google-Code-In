package main

import (
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://www.golang.org", 301) //link and web type.
}

func main() {
	http.HandleFunc("/", redirect)           // calling the redirect function
	err := http.ListenAndServe(":1104", nil) //localhost server. ex: localhost:1104
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
