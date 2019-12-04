package main

import (
	"fmt"
)

func main() {
	fmt.Print("What is your username? ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("What is your email? ")
	var email string
	fmt.Scanln(&email)

	var Array = []string{}
	Array = append(Array, username, email)
	fmt.Println(Array)
}
# go
# Go
