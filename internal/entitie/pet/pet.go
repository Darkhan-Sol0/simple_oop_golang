package pet

import (
	"myServ/internal/entitie/pet/cat"
	"myServ/internal/entitie/pet/dog"
)

type Pet interface {
	Describe() string
}

func NewPet(petType, name, owner string, age int) Pet {
	switch petType {
	case "cat":
		return cat.NewCat(name, owner, age)
	case "dog":
		return dog.NewDog(name, owner, age)
	default:
		return nil
	}
}
