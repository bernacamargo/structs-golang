package main

import "fmt"

type contactInfo struct {
	email string
	cep   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	bernardo := person{
		firstName: "Bernardo",
		lastName:  "Camargo",
		contactInfo: contactInfo{
			email: "personalemail@gmail.com",
			cep:   18342444,
		},
	}
	bernardo.updateName("Ber")
	bernardo.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) print() {
	fmt.Printf("%+v", p)
}
