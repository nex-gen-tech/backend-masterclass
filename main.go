package main

type Namer interface {
	GetFullName() string
	SetAge(age int) error
	GetAge() int
}

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func NewPerson(firstName, lastName string, age int) Namer {
	return &person{
		FirstName: firstName,
		LastName:  lastName,
	}
}

// GetFullName returns the full name of the person
func (p person) GetFullName() string {
	return p.FirstName + " " + p.LastName
}

// // SetAge sets the age of the person
func (p *person) SetAge(age int) error {
	p.Age = age
	return nil
}

// // GetAge -
func (p person) GetAge() int {
	return p.Age
}

func main() {
}
