package main

import "fmt"

// map in the go language
func map_learn() {

	// basically this is the full version of define map
	var mp map[string]int = map[string]int{"age": 26}
	fmt.Println(mp)

	//  also we can define map like below
	my_map := map[string]int{"age": 26}
	fmt.Println(my_map)

	// type can be changed like with below (we can enter type like array)
	my_array_map := map[string][]string{
		"teams": {"Financial", "IT", "HR"},
	}
	fmt.Println(my_array_map)

	// to adding key value in to the map
	my_array_map["metadata"] = []string{
		"UTF-8", "ASCII",
	}
	fmt.Println(my_array_map)

	// append a value for existing slice
	my_array_map["teams"] = append(my_array_map["teams"], "DevOps")
	fmt.Println(my_array_map)

	// delete map key and value
	delete(my_array_map, "metadata")
	fmt.Println(my_array_map)

	// check for key and value exists in the map
	value_of_metadata, ok := my_array_map["metadata"]
	fmt.Println(value_of_metadata, ok)

	value_of_teams, ok := my_array_map["teams"]
	fmt.Println(value_of_teams, ok)
}

//  functions in go language
func add_numbers(number1 int, number2 int) int {
	sum := number1 + number2
	return sum
}

// function with many return type
func many_return_type(number1 int, number2 int) (int, string) {
	sum := number1 + number2
	name := "Kavishka"
	return sum, name
}

// get function as a parameter
func call_func(inside_func func(number int) int) int {
	return inside_func(10)
}

func double(number int) int {
	return number * 2
}

// struct in go language
type Person struct {
	name        string
	age         int
	designation string
}

func main() {
	map_learn()

	value := add_numbers(5, 5)
	fmt.Println(value)

	value2, str := many_return_type(4, 7)
	fmt.Println(value2, str)

	value3 := call_func(double)
	fmt.Println(value3)

	person1 := Person{name: "Kavishka", age: 26, designation: "DevOps & Platform Engineer"}
	fmt.Println(person1)

	// define struct with Capital Name mandotory
	type Animal struct {
		name string
		age  uint
	}

	animal1 := Animal{name: "Dog", age: 3}
	fmt.Println(animal1)

	// access values by key name in strut
	fmt.Println(animal1.name)
}
