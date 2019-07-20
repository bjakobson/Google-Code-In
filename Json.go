package main

import (
	"fmt"
)

func main() {

	fmt.Print("Who is the manufacter? ")
	var Maker string
	fmt.Scanln(&Maker)

	fmt.Print("What is the Model? ")
	var Model string
	fmt.Scanln(&Model)

	fmt.Print("What is the model number? ")
	var mNumber string
	fmt.Scanln(&mNumber)

	var new = func() {
		type CameraClass struct {
			Manufacter string
			Model      string
			productID  string
		}
		CameraInfo := CameraClass{
			Manufacter: Maker,
			Model:      Model,
			productID:  mNumber,
		}
		fmt.Println(CameraClass(CameraInfo))

	}
	new()
}

//func main() {
//type CameraClass struct {
//Manufacter string
//Model      string
//productID  string
//}
//fmt.Println(Info)

//fmt.Print("Who is the manufacter? ")
//var Maker string
//fmt.Scanln(&Maker),

//fmt.Print("What is the Model? ")
//var Model string
//fmt.Scanln(&Model)

//fmt.Print("What is the model number? ")
//var mNumber string
//fmt.Scanln(&mNumber)

//CameraInfo := CameraClass{
//Manufacter: Maker,
//Model:      Model,
//productID:  mNumber,
//}
//fmt.Println(CameraClass(CameraInfo))
