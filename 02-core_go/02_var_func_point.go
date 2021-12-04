/**
* - Variables & functions
* - Pointers
 */

package main

import (
	"log"
)

var s = "seven"

func main() {
	var i int
	var whatToSay string

	whatToSay, _ = saySomething("Hello")

	log.Println(whatToSay)
	log.Println(i)

	var myString = "Green"
	log.Println("myString is set to", myString)

	// A pointer is a reference to a variable in memory
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)

	log.Println("s is", s)
	var s2 = "six"
	log.Println("s2 is", s2)
}

func saySomething(s string) (string, string) {
	return s, "World"
}

func changeUsingPointer(s *string) { // This actually makes it nicer than how this same thing is handled in JavaScript, for example
	log.Println("s is set to", s) // The memory address will be printed out
	newValue := "Red"
	*s = newValue
}
