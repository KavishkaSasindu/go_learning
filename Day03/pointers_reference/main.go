package main

import (
	"fmt"
)

// changing the struct by pointer
type Book struct {
	name string
	id   int
}

func (b *Book) setName(name string) string {
	b.name = name
	return b.name
}

func main() {
	x := 0
	y := &x
	fmt.Println(x, y) // it gives y value with the address of x stored in RAM
	fmt.Println(*y)   // it gives the value of that address of x stored in RAM

	// also we can changed the value from that
	*y = 100 // so it changed the value of address of x stored value
	fmt.Println(x)

	// example with struct
	// it do not change the copy of the struct, it changed the original struct value
	book := Book{name: "Sherlock Holmes", id: 1}
	fmt.Println(book.setName("The five friends"))
	fmt.Println(book)

}
