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

entitie/ <br>
├── interfaces.go # Интерфейсы <br>
├── entities.go # Структуры и конструкторы <br>
└── api.go # Методы объектов <br>

service/ <br>
└── service.go # Бизнес-логика <br>

## Ключевые принципы

- **Композиция вместо наследования** - встраивание структур
- **Интерфейсы определяют поведение** - а не структуры
- **Разделение ответственности** - четкие границы пакетов

## Пример использования

```
cat := entitie.NewPet("cat", "Barsik", "Baska", 5)
human := entitie.NewHuman("Baska", "ykt", 22)

service.Do(cat)   // 🐾 Pet: Barsik, Owner: Baska, Age: 5, Says: Mao!
service.Do(human) // 👨 Human: Baska, From: ykt, Age: 22, Says: HUMAN!

Иерархия интерфейсов

    personI (GetName, GetAge)
        ↓ 
    living (+Speak)
        ↙           ↘
    Human (+From)  Pet (+GetOwner)

Особенности реализации

    Неэкспортируемые структуры (person, human, cat, dog)

    Экспортируемые интерфейсы (Human, Pet, Living)

    Фабричные функции для создания объектов

    Type-switch для обработки разных типов в сервисе
```