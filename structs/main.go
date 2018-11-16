package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//Person struct
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@testing.com",
			zipCode: 23401,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}

func (pp *person) updateName(newName string) {
	(*pp).firstName = newName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
