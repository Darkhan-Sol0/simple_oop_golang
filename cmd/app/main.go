package main

import (
	"myServ/internal/entitie/human"
	"myServ/internal/entitie/pet"
	"myServ/internal/service"
)

func main() {
	h := human.NewHuman("Baska", "ykt", 22)
	cat := pet.NewPet("cat", "Barsik", 5)
	dog := pet.NewPet("dog", "Bobbik", 3)

	service.Do(h)
	service.Do(cat)
	service.Do(dog)
}
