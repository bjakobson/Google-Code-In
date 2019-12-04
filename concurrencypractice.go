package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// the main function that prints channel and assigns vaiables
func main() {
	c := make(chan string)           // assigning c to a channel
	s := strconv.Itoa(rand.Intn(30)) //converting int to str
	go Loop(s, c)                    // getting the function "Loop" and setting loop

	for msg := range c { // receiving message and making sure it is open and will stop when closed
		fmt.Println(msg)
	}

}

// the "logic" function that created the loop and channel
func Loop(thing string, c chan string) { //formatting the function to make thing a string and channel a string
	for i := 1; i <= 10; i++ { //creating a for loop that incremants and stops at 10
		c <- thing                        // sending message from thing to c
		time.Sleep(time.Millisecond * 50) // applying loop every 50 ms

		//ReturnValue := UserTyped
		//if ReturnValue == "S" {
		//close(c)

	}
	close(c) // stopping the channel after loop is finished
}
