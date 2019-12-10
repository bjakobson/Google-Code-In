package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func booksCheckedOut(w http.ResponseWriter, r *http.Request) {
	type booksOut struct {
		Booksout string
		Userout  string
	}

	var putIn = booksOut{Booksout: "MockingBird, House on Mango Street, Persopolis", Userout: "\n user2323, CGhs23443, jkdk345"}

	fmt.Fprintln(w, putIn)
}

func booksCheckedIn(w http.ResponseWriter, r *http.Request) {
	type booksIn struct {
		Books string
		User  string
	}

	var putOut = booksIn{Books: "Books: Lightning Theif, How to Train your dragon, Diary of a Wimpy kid.", User: "\n Users: Guest_6556, Wix223, Iwin46"}

	fmt.Fprintln(w, putOut)
}

func booksAvailable(w http.ResponseWriter, r *http.Request) {
	type booksAV struct {
		BooksAV string
	}

	var putAV = booksAV{BooksAV: "Harry Potter, \n Island of The Blue Dolphins, \n The Outsiders."}

	fmt.Fprintln(w, putAV)
}

func message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Library by Brandon Jakobson. At the end of the URL, type /av for availability, /in for books checked in, or /out for books checked out!")
}

func Execute() {
	// execute books checked out
	booksout := mux.NewRouter().StrictSlash(true)
	booksout.HandleFunc("/out", booksCheckedOut)
	booksout.HandleFunc("/in", booksCheckedIn)
	booksout.HandleFunc("/av", booksAvailable)
	booksout.HandleFunc("/", message)
	log.Fatal(http.ListenAndServe(":8080", booksout))

}

func main() {
	Execute()

}
# Google-Code-in
