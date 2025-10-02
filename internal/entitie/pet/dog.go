package pet

import (
	"fmt"
	"myServ/internal/entitie/owner"
)

type dog struct {
	owner.Person
}

func NewDog(name string, age int) Pet {
	return &dog{
		Person: owner.NewPerson(name, age),
	}
}

func (d *dog) Speak() string {
	return "Bark!"
}

func (d *dog) Describe() string {
	return fmt.Sprintf("Name: %s, Age: %d, Say: %s", d.GetName(), d.GetAge(), d.Speak())
}
