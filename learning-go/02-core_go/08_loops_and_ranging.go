/**
* - Loops and ranging over data
 */

package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	mySlice := []string{"dog", "cat", "horse", "fish", "banana"}

	for _, x := range mySlice {
		log.Println(x)
	}

	myMap := make(map[string]string)
	myMap["dog"] = "dog"
	myMap["fish"] = "fish"
	myMap["hat"] = "hat"

	// as you can see, maps are randomized
	for i, x := range myMap {
		log.Println(i, x)
	}

	var myUserSlice []User
	u1 := User{
		FirstName: "Jon",
	}
	u2 := User{
		FirstName: "Trevor",
	}

	myUserSlice = append(myUserSlice, u1)
	myUserSlice = append(myUserSlice, u2)

	for i, x := range myUserSlice {
		log.Println(i, x.FirstName)
	}
}
