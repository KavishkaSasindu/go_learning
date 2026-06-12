package main

import (
	"errors"
	"fmt"
)

func divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Can not divide by zero")
	} else {
		return num1 / num2, nil
	}
}

func main() {
	value, err := divide(1, 1)
	if err != nil {
		fmt.Println("There is an error")
	} else {
		fmt.Println(value)
	}

}
