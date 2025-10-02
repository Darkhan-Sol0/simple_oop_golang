package pet

import "myServ/internal/entitie/owner"

type Pet interface {
	owner.Person
	Describe() string
	Speak() string
}

func NewPet(petType, name string, age int) Pet {
	switch petType {
	case "cat":
		return NewCat(name, age)
	case "dog":
		return NewDog(name, age)
	default:
		return nil
	}
}
