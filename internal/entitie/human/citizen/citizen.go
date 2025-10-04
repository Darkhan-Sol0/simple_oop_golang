package citizen

import (
	"fmt"
	"myServ/internal/entitie/base"
)

type citizen struct {
	base.Person
	City string
}

func NewCitizen(name, city string, age int) *citizen {
	return &citizen{
		Person: base.NewPerson(name, age),
		City:   city,
	}
}

func (h *citizen) GetCity() string {
	return h.City
}

func (h *citizen) Describe() string {
	return fmt.Sprintf("Name: %s, Age: %d, City: %s", h.GetName(), h.GetAge(), h.GetCity())
}
