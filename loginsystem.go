package main

import (
	"fmt"
)

func main() {
	//calling the function, UserCredentials
	UserCredentials()
}

//Logic function, letting computer know that we will be returning a boolean
func UserCredentials() bool {

	//getting the username
	fmt.Println("What is your Username? ")
	var UserName string
	fmt.Scan(&UserName)

	//getting the password
	fmt.Println("What is your password? ")
	var PassWord string
	fmt.Scan(&PassWord)

	//definig UserInfo making username and password strings
	type UserInfo struct {
		username string
		password string
	}
	//taking the structure and apllying it with the users inputs, in this case, username and password.
	UserClass := UserInfo{
		username: UserName,
		password: PassWord,
	}

	//setting applicitable as a bollean, and false.
	//getting the length of the given username and password
	applicitable := false
	nameLength := len(UserName)
	passLength := len(PassWord)

	//if length is under 10 for username and password, applicitable is true. Otherwise, applicitable is false.
	if nameLength < 10 && passLength < 10 {
		applicitable = true
	} else {
		applicitable = false
	}
	//if applicitable is true, they are welcomed. Otherwise, permission denied.
	if applicitable == true {
		fmt.Println("Welcome, ", UserClass.username, UserClass.password)
	} else {
		fmt.Println("Cannot register you.")
	}

	//returning the value of applicitable
	return applicitable
}
Go
