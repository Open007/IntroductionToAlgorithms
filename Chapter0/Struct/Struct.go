package main

import (
	"fmt"
)

type books struct {
	name string
}

func (book books) modify_1(name string) {
	book.name = name
}

func (book *books) modify_2(name string) {
	book.name = name
}

func modify_3(book books, name string) {
	book.name = name
}

func modify_4(book *books, name string) {
	book.name = name
}

type books_inter_1 interface {
	modify_1(name string)
	modify_2(name string)
}

func modify_5(book books_inter_1, name string) {
	book.modify_1(name)
}

func modify_6(book books_inter_1, name string) {
	book.modify_2(name)
}

func main() {
	a := books{"A"}
	fmt.Print("Origin Name: " + a.name + "\n")

	a.modify_1("Modify_1")
	fmt.Print("Modify_1 Name: " + a.name + "\n")

	a.modify_2("Modify_2")
	fmt.Print("Modify_2 Name: " + a.name + "\n")

	b := books{"B"}
	fmt.Print("Origin Name: " + b.name + "\n")

	modify_3(b, "Modify_3")
	fmt.Print("Modify_3 Name: " + b.name + "\n")

	modify_4(&b, "Modify_4")
	fmt.Print("Modify_4 Name: " + b.name + "\n")

	c := books{"C"}
	fmt.Print("Origin Name: " + c.name + "\n")

	modify_5(&c, "Modify_5")
	fmt.Print("Modify_5 Name: " + c.name + "\n")

	modify_6(&c, "Modify_6")
	fmt.Print("Modify_6 Name: " + c.name + "\n")

}
