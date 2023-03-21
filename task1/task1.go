package task1

import (
	"fmt"
)

// Определяем структуру Human с несколькими полями и методами

type Human struct {
	Name string
	Age  int
	Sex  Sex
}

func NewHuman(name string, age int, sex Sex) *Human {
	human := Human{name, age, sex}
	return &human
}

func (h *Human) Say(smth string) {
	fmt.Println(smth)
}

func (h *Human) Walk() {
	fmt.Println("Walking...")
}

func (h *Human) Sit() {
	fmt.Println("Sitting...")
}

// Реализуем композицию структур Human и Action. Action содержит Human как анонимное поле.
// Таким образом происходит автоматически встраивание методов Human в Action.

type Action struct {
	Human
}

func NewAction(human Human) *Action {
	return &Action{human}
}

// Теперь объект Action имеет поля и методы объекта Human. При желании их можно переопределить.
// Например:

func (a *Action) Say(smth string) {
	fmt.Println(a.Name, "saying:", smth)
}

func Run() {
	human := NewHuman("Egor", 21, Male)
	action := NewAction(*human)
	action.Say("Hello!")
	action.Walk()
	action.Sit()
}
