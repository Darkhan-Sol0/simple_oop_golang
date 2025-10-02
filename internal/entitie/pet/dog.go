package pet

import (
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
