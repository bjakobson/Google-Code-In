package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter first word: ")
	var Firstword string
	fmt.Scanln(&Firstword)
	fmt.Print("Enter second word: ")
	var Secondword string
	fmt.Scanln(&Secondword)

	fmt.Println(Firstword + " " + Secondword)
}
