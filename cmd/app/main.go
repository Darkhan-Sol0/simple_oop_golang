package main

import (
	"myServ/internal/entitie"
	"myServ/internal/service"
)

func main() {
	cat := entitie.NewPet("cat", "Barsik", "Baska", 5)
	dog := entitie.NewPet("dog", "Bobik", "Baska", 3)
	human := entitie.NewHuman("Baska", "ykt", 22)

	service.Do(cat)
	service.Do(dog)
	service.Do(human)
}
