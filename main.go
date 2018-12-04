package main

import (
	"errors"
	"fmt"
)

//Function to add two integers
func Add(value1 int, value2 int) int {

	return value1 + value2
}

//Function to subtract two integers
func Subtract(value1 int, value2 int) int {
	return value1 - value2
}

//Function to multiply two integers
func Multiply(value1 int, value2 int) int {
	return value1 * value2
}

//Function to divide two integers
func Divide(value1 int, value2 int) (int, error) {

	//Division by zero : Returns max integer value and an error
	if value2 == 0 {
		return (1<<16 - 1), errors.New("Division by zero exception")
	}

	return value1 / value2, nil
}

func main() {
	fmt.Println("Testing a simple application")
}
