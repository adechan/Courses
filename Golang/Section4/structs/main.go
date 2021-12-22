package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipCode int
}

// custom type
type person struct{
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	// Method1
	// ade := person{"Ade", "Chan"}

	// Method2
	// ade := person{}
	// var ade person
	// ade.firstName = "Ade"
	// ade.lastName = "Chan"

	// Method3
	// ade := person{
	// 	firstName: "Ade",
	// 	lastName: "Chan",
	// }

	// fmt.Println(ade)

	// ===============================

	jim := person{
		firstName: "Jim",
		lastName: "Meow",
		contact: contactInfo{
			email: "jim.meow@mail.com",
			zipCode: 1234,
		},
	}

	// jim: a reference to the struct in memory (the actual value of the struct)
	// &jim: memory address = pointer
	// jimPointer := &jim // give me the MEMORY ADDRESS of the value this variable is pointing at
	// jimPointer = some memory address
	// jimPointer.updateName("Jimmy")

	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// it makes a copy out of p and it changes the first name for the copied value
// not the original p
// because it is PASS BY VALUE
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// PASS BY REFERENCE
// take this memory address and give me the value that exists there

// a pointer that points to a person
// * is not an operator
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

