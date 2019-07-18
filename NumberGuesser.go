package main

import (
	"fmt"
)

func main() {
	num := 5
	fmt.Print("What is your guess between numbers 1-10? ")
	var Fguess int
	fmt.Scanln(&Fguess)
	if Fguess > num {
		fmt.Println("Number is too high")
	} else if Fguess < num {
		fmt.Println("Guess is too low")
	} else {
		fmt.Println("Yep, answer was 5!")
	}

}
