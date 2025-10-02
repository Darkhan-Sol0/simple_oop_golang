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
│   └── service.go # Универсальный сервис через интерфейс Obj <br>
```
## Ключевые принципы

- **Композиция вместо наследования** - встраивание структур через `owner.Person`
- **Интерфейсы определяют поведение** - четкое разделение контрактов
- **Разделение ответственности** - каждый пакет отвечает за свою область
- **Фабричный паттерн** - централизованное создание объектов через `NewPet()`
- **Полиморфизм** - единый интерфейс для разных типов объектов

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

    service.Do(h)   // Name: Baska, Age: 22, City: ykt
    service.Do(cat) // Name: Barsik, Age: 5, Say: Mao!
    service.Do(dog) // Name: Bobbik, Age: 3, Say: Bark!
}
```

## Иерархия интерфейсов
```
Person (GetName, GetAge)
    ↙           ↘
Human (+GetCity, +Discribe)  Pet (+Speak, +Discribe)
                                    ↙           ↘
                               Cat (+Discribe)  Dog (+Discribe)

Obj (Discribe) ← Реализуют все сущности
```

## Особенности реализации

    Неэкспортируемые структуры - person, human, cat, dog

    Экспортируемые интерфейсы - Human, Pet, Person, Obj

    Фабричные функции - NewHuman(), NewPet(), NewCat(), NewDog()

    Универсальный сервис - функция Do() работает через интерфейс Obj

    Композиция - переиспользование кода через embedding структур

    Полиморфизм - каждая сущность сама определяет свое строковое представление