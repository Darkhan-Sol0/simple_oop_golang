# Simple OOP-composition on Golang

Минималистичный пример объектно-ориентированного программирования через композицию в Go.

## Архитектура
```
base.Person (базовый интерфейс)
      ↙           ↘
Human (фабрика) Pet (фабрика)
    ↙   ↘           ↙   ↘
Citizen Villager  Cat   Dog
```

## Структура проекта
```
internal/
├── entitie/
│ ├── base/
│ │     └── base.go # Базовый интерфейс Person
│ ├── human/
│ │     ├── human.go # Human фабрика и интерфейс
│ │     ├── citizen/
│ │     │   └── citizen.go # Citizen реализация
│ │     └── villager/
│ │         └── villager.go # Villager реализация
│ └── pet/
│       ├── pet.go # Pet фабрика и интерфейс
│       ├── cat/
│       │     └── cat.go # Cat реализация
│       └── dog/
│             └── dog.go # Dog реализация
├── service/
│       └── service.go # Универсальный сервис через интерфейс Obj
```
## Ключевые принципы

- **Композиция вместо наследования** - встраивание структур через `base.Person`
- **Интерфейсы определяют поведение** - четкое разделение контрактов
- **Разделение ответственности** - каждый пакет отвечает за свою область
- **Фабричный паттерн** - централизованное создание объектов через `NewHuman()` и `NewPet()`
- **Полиморфизм** - единый интерфейс `Obj` для всех сущностей

## Пример использования

```
go
package main

import (
    "myServ/internal/entitie/human"
    "myServ/internal/entitie/pet"
    "myServ/internal/service"
)

func main() {
    citizen := human.NewHuman("citizen", "Yakutsk", "Baska", 22)
    villager := human.NewHuman("villager", "Maya", "Uyban", 42)
    cat := pet.NewPet("cat", "Barsik", "Baska", 5)
    dog := pet.NewPet("dog", "Bobbik", "Uyban", 3)

    service.Do(citizen)  // Name: Baska, Age: 22, City: Yakutsk
    service.Do(villager) // Name: Uyban, Age: 42, Village: Maya
    service.Do(cat)      // Name: Barsik, Age: 5, Owner: Baska, Say: Mao!
    service.Do(dog)      // Name: Bobbik, Age: 3, Owner: Uyban, Say: Bark!
```

## Иерархия интерфейсов
```
base.Person (GetName, GetAge)
    ↙               ↘
Human (фасад)       Pet (фасад)
    ↙     ↘           ↙       ↘
Citizen  Villager   Cat       Dog
    ↓       ↓         ↓         ↓
(Describe) (Describe) (Describe) (Describe)

Obj (Describe) ← Реализуют все сущности
```

## Особенности реализации
    Неэкспортируемые структуры - citizen, villager, cat, dog

    Экспортируемые интерфейсы - Human, Pet, base.Person, service.Obj

    Фабричные функции - NewHuman(), NewPet(), NewCitizen(), NewVillager(), NewCat(), NewDog()

    Универсальный сервис - функция Do() работает через интерфейс Obj

    Композиция - переиспользование кода через embedding структур base.Person

    Владельцы животных - питомцы связаны с людьми через поле owner

    Полиморфизм - каждая сущность сама определяет свое строковое представление