package cat

import (
	"fmt"
	"myServ/internal/entitie/base"
)

type cat struct {
	base.Person
	owner string
}

func NewCat(name, owner string, age int) *cat {
	return &cat{
		Person: base.NewPerson(name, age),
		owner:  owner,
	}
}

func (c *cat) Speak() string {
	return "Mao!"
}

func (c *cat) GetOwner() string {
	return c.owner
}

func (c *cat) Describe() string {
	return fmt.Sprintf("Name: %s, Age: %d, Owner: %s, Say: %s", c.GetName(), c.GetAge(), c.GetOwner(), c.Speak())
}
