package organization

import (
	"fmt"
)

//Person struct
type Person struct {
	FirstName string
	LastName  string
}

//NewPerson - creates a new person
func NewPerson(fn, ln string) Person {
	return Person{
		FirstName: fn,
		LastName:  ln,
	}
}

//FullName gives the full name of the person
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}
