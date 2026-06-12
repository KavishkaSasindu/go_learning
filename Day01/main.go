package main

import "fmt"

// printing statements
func print_values() {
	fmt.Println("Hello! World")
}

// data types and variables
func data_types() {
	// ----------- data types ------------
	// 1. uint32, uint64 = unassigned integer (can not store negative values)
	// 2. int8, int16, int32, int64 = assigned integer (can store positive values and negative values)
	// 3. string = a sequence of characters
	// 4. bool = boolean value (true or false)
	// 5. float32, float64 = floating-point numbers (numbers with decimal points)
	// 6. byte = a single byte (int8) of data, often used to represent characters or small integers
	// 7. rune = a single Unicode code point (int32), used to represent characters in Go

	const number int = 32 // based on here it is constant value and it cannot be changed
	fmt.Println(number)

	var name string = "Kavishka Sasindu"
	fmt.Println(name)

	// multiple variable declaration
	var (
		first_name  string = "Kavishka"
		middle_name string = "Sasindu"
		last_name   string = "Mudithananda"
	)
	fmt.Println(first_name, middle_name, last_name)
}

// implicit and explicit variable declaration
func variables() {
	// explicit variable declaration
	var name string = "Kavishka"
	fmt.Println(name)

	// implicit variable declaration
	name2 := "Sasindu"
	fmt.Println(name2)

	// -------------when you use implicit and explicit variable -----------------
	//  if you define variable and assign value you can use implicit variable declaration (variable_name := value)
	//  otherwise you can use explicit variable declaration (var variable_name data_type = value) then after you can assign value to the varibale
}

// type casting
func type_casting() {
	number := 3.2
	integer_number := int(number)
	fmt.Println(integer_number)
}

//  console output
func console_output() {
	//  ----- println----------(basically it prints with a new line)
	fmt.Println("Hello! World")

	// ----------printf ------------- (very usefull when print with arguments)

	// 1 value  = %v
	age := 25
	fmt.Printf(" My age is %v", age)

	// 2 type = %T
	fmt.Printf("Age type is %T \n", age)

	// 3 also it has string= %s like wise binary = %b

}

// conditions and conditionals
func conditions_conditionals() {
	number := 10

	// here we do not use paranthesis
	if number < 10 {
		fmt.Printf("The number %v is less than 10 \n", number)
	} else if number > 10 {
		fmt.Printf("The number %v is greater than 10 \n", number)
	} else {
		fmt.Printf("The number is %v equal to 10 \n", number)
	}

	number2 := 2

	switch number2 {
	case 1:
		if number2 > 2 {
			fmt.Printf("The number is greater than 2 \n")
		}
	case 2:
		if number2 < 2 {
			fmt.Printf("The number is less than 2 \n")
		}
	default:
		fmt.Println("The number is equal to 2")
	}

	// switch  {
	// case number2 > 2:
	// 	fmt.Printf("The number is greater than 2 \n")
	// case number2 < 2:
	// 	fmt.Printf("The number is less than 2 \n")
	// default:
	// 	fmt.Println("The number is equal to 2")
	// }

}

// loops
func loops() {

	// ----- for loop -----
	for i := 1; i < 10; i++ {
		fmt.Printf("i is %v \n", i)
	}

	//  now to write while loop basically it just like below
	idx := 10

	for idx < 15 { // also we can use  infinte loop with for ::  as an example -->  " for true {} "
		fmt.Printf("%v \n", idx)
		idx++
	}
}

//  about strings
func about_string() {
	/* when we use string it automatically assign as uint8 which mean 1byte
	 */

	str := "hello world"
	fmt.Printf("%v \n", str[0]) // it prints integer value
	fmt.Printf("%T \n", str[0]) // this print uint8

	for _, char := range str { // if you do not need like "for i, char := rage str " here i do not wnat "i" so i can simply ignore it by writing "_"
		fmt.Printf("%c \n", char)
	}
}

// Arrays
func array_learn() {
	// basically we have some ways to define an array in go

	// 1. var array_name [size]data_type
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v \n", numbers)

	// 2. array_name := [size]data_type{value1, value2, value3, ...}
	numbers2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v \n", numbers2)
	numbers2[0] = 10 // we can change the value of the array by using index
	fmt.Printf("%v \n", numbers2)

	// 3. array_name := [...]data_type{value1, value2, value3, ...} // the size of the array is determined by the number of values provided
	numbers3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%v \n", numbers3)

	for i, number := range numbers3 {
		fmt.Printf("Index: %v, Number: %v \n", i, number)
	}
}

//  slices in go
func slice_go() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:]
	fmt.Println(sl)
	// so basically learning slices we can manipulate the array what we define
	sl1 := arr[:2]
	fmt.Println(sl1)
}

func main() {

	print_values()
	data_types()
	variables()
	type_casting()
	console_output()
	conditions_conditionals()
	loops()
	about_string()
	array_learn()
	slice_go()
}
