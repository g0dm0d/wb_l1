package main

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	name string
	age  int
}

func NewHuman(name string, age int) *Human {
	return &Human{
		name: name,
		age:  age,
	}
}

func (h *Human) NewName(name string) {
	h.name = name
}

func (h *Human) isHuman() bool {
	return true
}

type Action struct {
	*Human
}

func NewAction(human *Human) *Action {
	return &Action{
		Human: human,
	}
}

func (a *Action) NewAge(age int) {
	a.age = age
}

func (a *Action) isHuman() bool {
	return false
}

func task1() {
	person := NewHuman("Oleg", 10)

	action := NewAction(person)

	// Когда струтура встроенна мы можем обращаться к методам и полям на прямую
	action.NewName("Neoleg")
	action.NewAge(20)

	// Если методы функции имеют одинаковое название отрабатывает слой вложенности который выше
	action.isHuman() // false

	action.Human.isHuman() // true
}
