package main

import (
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://www.golang.org", 301)
}

func main() {
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":9191", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
# Google-Code-in
