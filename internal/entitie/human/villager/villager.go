package villager

import (
	"fmt"
	"myServ/internal/entitie/base"
)

type villager struct {
	base.Person
	City string
}

func NewVillager(name, city string, age int) *villager {
	return &villager{
		Person: base.NewPerson(name, age),
		City:   city,
	}
}

func (h *villager) GetCity() string {
	return h.City
}

func (h *villager) Describe() string {
	return fmt.Sprintf("Name: %s, Age: %d, Villige: %s", h.GetName(), h.GetAge(), h.GetCity())
}
