package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	ContactInfo
}

func main() {
	p1 := Person{firstName: "Alice", lastName: "Smith", ContactInfo: ContactInfo{email: "alice@example.com", zipCode: 12345}}
	p2 := Person{firstName: "Bob", lastName: "Johnson", ContactInfo: ContactInfo{email: "bob@example.com", zipCode: 67890}}

	p1.updateName("Alice", "Johnson")

	p1.print()
	p2.print()
}

func (p *Person) updateName(firstName, lastName string) {
	p.firstName = firstName
	p.lastName = lastName
}

func (p Person) print() {
	fmt.Printf("Name: %s %s\n", p.firstName, p.lastName)
	fmt.Printf("Email: %s\n", p.ContactInfo.email)
	fmt.Printf("Zip Code: %d\n", p.ContactInfo.zipCode)
}
