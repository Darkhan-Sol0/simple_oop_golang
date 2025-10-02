package pet

import (
	"myServ/internal/entitie/owner"
)

type cat struct {
	owner.Person
}

func NewCat(name string, age int) Pet {
	return &cat{
		Person: owner.NewPerson(name, age),
	}
}

func (c *cat) Speak() string {
	return "Mao!"
}
