package main

import (
	"fmt"
)

func main() {
	var angles = make(map[string]int)
	angles["Triangle"] = 3
	angles["Square"] = 4
	angles["Hexagon"] = 6
	for i := 0; i < 3; i++ {
		fmt.Println(angles)
		delete(angles, "Triangle")
		fmt.Println(angles)
		delete(angles, "Hexagon")
		fmt.Println(angles)
		break
	}
}
# Go
