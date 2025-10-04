package dog

import (
	"fmt"
	"myServ/internal/entitie/base"
)

type dog struct {
	base.Person
	owner string
}

func NewDog(name, owner string, age int) *dog {
	return &dog{
		Person: base.NewPerson(name, age),
		owner:  owner,
	}
}

func (d *dog) Speak() string {
	return "Bark!"
}

func (d *dog) GetOwner() string {
	return d.owner
}

func (d *dog) Describe() string {
	return fmt.Sprintf("Name: %s, Age: %d, Owner: %s, Say: %s", d.GetName(), d.GetAge(), d.GetOwner(), d.Speak())
}
