package main

import (
	"fmt"
)

type EmployeeInfo struct {
	Firstname, LastName string
}

func main() {

	fmt.Print("What is your first name? ")
	var Fname string
	fmt.Scanln(&Fname)

	fmt.Print("What is your last name? ")
	var Lname string
	fmt.Scanln(&Lname)

	e := EmployeeInfo{
		Firstname: Fname,
		LastName:  Lname,
	}

	fmt.Println(e.Firstname, e.LastName)
}
# Go
