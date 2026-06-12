package main

import "fmt"

//  only define methods without implementations
// to use interface methods we need struct and also a function with struct
type Shape interface {
	getPerimeter() int
}

// then create struct to use that interface method
type Triangle struct {
	a int
	b int
	c int
}

// function to use interface implementation
func (t Triangle) getPerimeter() int {
	return t.a + t.b + t.c
}

func main() {
	// use interface method using struct and function
	var shape Shape = Triangle{a: 1, b: 2, c: 3}
	triangle_permiter := shape.getPerimeter()
	fmt.Println(triangle_permiter)
}
