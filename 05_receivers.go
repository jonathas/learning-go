/**
* - Receivers structs with functions
 */

package main

import "log"

type myStruct struct {
	FirstName string
}

// This is a receiver in the function
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Jon"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstName()) // The receiver "adds" the function to the struct
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
