package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func booksCheckedIn(w http.ResponseWriter, r *http.Request) {
	type booksIn struct {
		Name string
		Book string
	}

	booksin := &booksIn{
		Name: "Frank, Bob",
		Book: " Fault in Our Stars, The Hobbit",
	}
	b, _ := json.Marshal(booksin)
	s := string(b)
	open := strings.Replace(s, "{", "", -1)
	close := strings.Replace(open, "}", "", -1)
	fmt.Fprintln(w, string(close))
}

func booksCheckedOut(w http.ResponseWriter, r *http.Request) {
	type booksOut struct {
		Name string
		Book string
	}

	booksOUT := &booksOut{
		Name: "Betty Joe, Adam Sandler",
		Book: " To Kill a Mocking Bird, Who's Who?",
	}
	bo, _ := json.Marshal(booksOUT)
	s := string(bo)
	open := strings.Replace(s, "{", "", -1)
	close := strings.Replace(open, "}", "", -1)
	fmt.Fprintln(w, string(close))
}

func booksAvailable(w http.ResponseWriter, r *http.Request) {
	type available struct {
		Author string
		Book   string
	}

	booksav := &available{
		Author: "J.K Rowling, Sandra Cisneros",
		Book:   " Harry Potter, House on Mango Street",
	}
	av, _ := json.Marshal(booksav)
	s := string(av)
	open := strings.Replace(s, "{", "", -1)
	close := strings.Replace(open, "}", "", -1)
	fmt.Fprintln(w, string(close))
}

func popularBooks(w http.ResponseWriter, r *http.Request) {
	type available struct {
		Author string
		Book   string
	}

	pop := &available{
		Author: "Suzanne Collins, Stephenie Meyer ",
		Book:   " The Hunger Games, Twighlight",
	}
	bp, _ := json.Marshal(pop)
	s := string(bp)
	open := strings.Replace(s, "{", "", -1)
	close := strings.Replace(open, "}", "", -1)
	fmt.Fprintln(w, string(close))

}

func message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Library by Brandon Jakobson. At the end of the URL, type /av for availability, /in for books checked in, /out for books checked out, or /pop for popular books!")
}

func Execute() {
	booksout := mux.NewRouter().StrictSlash(true)
	booksout.HandleFunc("/out", booksCheckedOut)
	booksout.HandleFunc("/in", booksCheckedIn)
	booksout.HandleFunc("/av", booksAvailable)
	booksout.HandleFunc("/pop", popularBooks)
	booksout.HandleFunc("/", message)
	log.Fatal(http.ListenAndServe(":8080", booksout))
}

func main() {
	Execute()
}
