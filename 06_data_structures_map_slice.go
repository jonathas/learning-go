/**
* - Other data structures maps and slices
 */

package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// maps
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	myMap["dog"] = "Spike"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]int)
	myMap2["First"] = 1
	myMap2["Second"] = 2

	log.Println(myMap2["First"], myMap2["Second"])

	myCustomMap := make(map[string]User)
	me := User{
		FirstName: "Jon",
		LastName:  "Ribeiro",
	}

	myCustomMap["me"] = me
	log.Println(myCustomMap["me"].FirstName)

	// slices
	var mySlice []string
	mySlice = append(mySlice, "Ribeiro")
	mySlice = append(mySlice, "Jon")
	log.Println(mySlice)

	sort.Strings(mySlice)
	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)
	log.Println(numbers[0:2])
	log.Println(numbers[6:9])

	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
