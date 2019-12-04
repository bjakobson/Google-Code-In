package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 3
	sum := x + y
	fmt.Println(sum)

	if sum > 5 {
		fmt.Println("Too high")
	} else if sum < 3 {
		fmt.Println("Too low")
	} else {
		fmt.Println("Yep, 4")
	}

}
