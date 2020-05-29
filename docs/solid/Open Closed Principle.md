### Open/Closed Principle (принцип открытости/закрытости)

> Система должна быть открыта для расширения и закрыта для модификации.

Не будем уходить далеко и рассмотрим пример с животными, структура `Animal`:

```go
type Animal struct {
	name string
}
```

Мы хотим перебрать список животных, каждое из которых представлено объектом класса Animal, и узнать о том, какие звуки они издают. 
Представим, что мы решаем эту задачу с помощью функции `AnimalSounds`:

```go
func AnimalSounds() {
	animals := []Animal{
		Animal{name: "lion"},
		Animal{name: "mouse"},
	}

	for _, animal := range animals {
		if animal.name == "lion" {
			fmt.Println("roar")
		} else if animal.name == "mouse" {
			fmt.Println("squeak")
		}
	}
}
```

Самая главная проблема такой архитектуры заключается в том, что функция определяет то, какой звук издаёт то или иное животное, анализируя конкретные объекты. 
Функция `AnimalSounds` не соответствует принципу открытости-закрытости, так как, например, при появлении новых видов животных, нам, для того, чтобы с её помощью можно было бы узнавать звуки, издаваемые ими, придётся её изменить.


Добавим в массив новый элемент:

```go
func AnimalSounds() {
	animals := []Animal{
		Animal{name: "lion"},
		Animal{name: "mouse"},

        // Новый
		Animal{name: "snake"},
	}
    
    ...
}
```

После этого нам придётся поменять код функции `AnimalSounds`:

```go
func AnimalSounds() {
	...

	for _, animal := range animals {
		if animal.name == "lion" {
			fmt.Println("roar")
		} else if animal.name == "mouse" {
			fmt.Println("squeak")
		} else if animal.name == "snake" {
			fmt.Println("hiss")
		}
	}
}
```

Как видите, при добавлении в массив нового животного придётся дополнять код функции. 
Пример это очень простой, но если подобная архитектура используется в реальном проекте, функцию придётся постоянно расширять, добавляя в неё новые выражения `if`.

Как привести функцию `AnimalSounds` в соответствие с принципом открытости-закрытости? Например — так:

```go
type Animal interface {
	MakeSound() string
}

type Lion struct {}
func (lion *Lion) MakeSound() string {
	return "roar"
}

type Squirrel struct {}
func (squirrel *Squirrel) MakeSound() string {
	return "squeak"
}

type Snake struct {}
func (snake *Snake) MakeSound() string {
	return "hiss"
}

func AnimalSounds() {
	animals := []Animal{
		&Lion{},
		&Squirrel{},
		&Snake{},
	}

	for _, animal := range animals {
		fmt.Println(animal.MakeSound())
	}
}
```

Можно заметить, что у кадой структуры реализующий интерфейс `Animal` теперь есть метод `MakeSound`. 
При таком подходе нужно, чтобы структуры, предназначенные для описания конкретных животных, реализовывали бы интерфейс `Animal`.

В результате у каждой стуктуры, описывающего животного, будет собственный метод `MakeSound`, а при переборе массива с животными в функции `AnimalSounds` достаточно будет вызвать этот метод для каждого элемента массива.

Если теперь добавить в массив объект, описывающий новое животное, функцию `AnimalSounds` менять не придётся. 
Мы привели её в соответствие с принципом открытости-закрытости.

Код: [Принцип Открытости/Закрытости](./code/solid/open-closed/open-closed.go)