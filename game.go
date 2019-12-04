package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//main function that is printing the logic
func main() {
	Number()   // executing the Number function
	Birthday() // executing the birthday function

}

//the logic function that is getting number and looping to the number
func Number() int {

	fmt.Print("Tell me any number ") // get user to give value
	var num int
	fmt.Scan(&num) //getting the value given

	valid := true //making valid a boolean
	i := 1        // setting i's value to 0

	// This loop continues while "valid" is true.
	for valid { // while valid is true

		// If i equals inputted number, set "valid" to false.
		if i == num { // once the loop gets to the users preset number
			valid = false // make vaild fals and stop the loop
		}
		fmt.Println(i)
		i++                               //increment the loop till valid is false
		time.Sleep(time.Millisecond * 50) // loop will wait 50 ms before displaying the next number

	}
	return i //setting i to be returned to the main function
}

//Birthday function
func Birthday() string {

	fmt.Print("What is your birthday? Year-MM-Day. ") //get users input
	var Bday string
	fmt.Scan(&Bday)

	var bReplacer = strings.NewReplacer("-", "")
	NStr := bReplacer.Replace(Bday) // replace '-' with nothing so there are no spaces

	e, er := strconv.Atoi(NStr) //convert the string to an int
	if er != nil {              // if there is an error
		fmt.Println("You printed your age wrong.")
	}

	currentTime := time.Now()                   //get the current time
	format := currentTime.Format("2006-01-02")  //format the time
	var replacer = strings.NewReplacer("-", "") //replace '-'with nothing so there are no spaces
	str := replacer.Replace(format)             //confirm the changes

	i, err := strconv.Atoi(str) //convert the string to an int
	if err != nil {             // if there is an error
		fmt.Println("You printed your age wrong.")
	}

	age := i - e                //subtract todays current date with the users birthday
	newAge := strconv.Itoa(age) //make age a string

	CurrentAge := newAge[:2]

	fmt.Println("You are: ", CurrentAge) // get the first 2 digits

	return Bday
}
# Go
