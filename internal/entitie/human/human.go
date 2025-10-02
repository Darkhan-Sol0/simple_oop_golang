package human

import "myServ/internal/entitie/owner"

type human struct {
	owner.Person
	City string
}

type Human interface {
	owner.Person
	GetCity() string
}

func NewHuman(name, city string, age int) Human {
	return &human{
		Person: owner.NewPerson(name, age),
		City:   city,
	}
}

func (h *human) GetCity() string {
	return h.City
}
