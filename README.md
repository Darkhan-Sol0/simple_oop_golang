# Simple OOP-composition on Golang

Минималистичный пример объектно-ориентированного программирования через композицию в Go.

## Архитектура
```
Person (базовый интерфейс) 
  ↙    ↘ 
Human   Pet 
      ↙ ↘ 
    Cat Dog
```

## Структура проекта
```
internal/<br>
├── entitie/<br>
│   ├── owner/<br>
│   │   └── owner.go # Базовый интерфейс Person<br>
│   ├── human/<br>
│   │   └── human.go # Human структура и интерфейс<br>
│   ├── pet/<br>
│   │   ├── pet.go # Pet интерфейс и фабрика<br>
│   │   ├── cat.go # Cat реализация<br>
│   │   └── dog.go # Dog реализация<br>
├── service/<br>
│   └── service.go # Бизнес-логика <br>
```
## Ключевые принципы

- **Композиция вместо наследования** - встраивание структур через `owner.Person`
- **Интерфейсы определяют поведение** - четкое разделение контрактов
- **Разделение ответственности** - каждый пакет отвечает за свою область
- **Фабричный паттерн** - централизованное создание объектов через `NewPet()`

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
    h := human.NewHuman("Baska", "ykt", 22)
    cat := pet.NewPet("cat", "Barsik", 5)
    dog := pet.NewPet("dog", "Bobbik", 3)

    service.Do(h)   // Human: Baska from ykt, 22 years old
    service.Do(cat) // Pet: Barsik (Mao!), 5 years old
    service.Do(dog) // Pet: Bobbik (Bark!), 3 years old
}
```

## Иерархия интерфейсов
```
Person (GetName, GetAge)
    ↙           ↘
Human (+GetCity)  Pet (+Speak)
```

## Особенности реализации

    Неэкспортируемые структуры - person, human, cat, dog

    Экспортируемые интерфейсы - Human, Pet, Person

    Фабричные функции - NewHuman(), NewPet(), NewCat(), NewDog()

    Type-switch - полиморфная обработка в сервисном слое

    Композиция - переиспользование кода через embedding структур