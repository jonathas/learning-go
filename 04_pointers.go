/**
* - struct
 */

package main

import (
	"log"
	"time"
)

// We use capital letters as convention meaning something is public
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

// Not in capital lettes, means it's private
var special string

func main() {
	user := User{
		FirstName:   "Jon",
		LastName:    "Ribeiro",
		PhoneNumber: "1 555 555-1212",
	}

	log.Println(user.FirstName, user.LastName, user.PhoneNumber)
}
