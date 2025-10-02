package pet

import (
	"fmt"
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

func (c *cat) Describe() string {
	return fmt.Sprintf("Name: %s, Age: %d, Say: %s", c.GetName(), c.GetAge(), c.Speak())
}
