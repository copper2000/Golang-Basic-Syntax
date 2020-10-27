package main

import "fmt"

func main() {
	var number int
	number = 10
	fmt.Println(number)

	var number1 int = 11
	fmt.Println(number1)

	// Type Inference using var only
	var email = "cooper@gmail.com"
	fmt.Println(email)

	// Declare multiple variables
	// 1. Declare multiple variables with the same data type

	var a, b int
	a = 1
	b = 2
	fmt.Println(a, b)

	var a1, b1 = 10, 11
	fmt.Println(a1, b1)

	var a2, b2 = 12, 13
	fmt.Println(a2, b2)

	// 2. Declare many variables but different data types

	var (
		name string
		age int
		address string
	)

	name = "Batman"
	address = "Gotham"
	age = 40

	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(age)

	var (
		name1 string = "Iron man"
		age1 int = 45
		address1 string = "California"
	)

	fmt.Println(name1)
	fmt.Println(address1)
	fmt.Println(age1)

	// Type Inference using var only for multiple declare

	var name2, age2, address2 = "Spider man", 16, "Queen"

	fmt.Println(name2)
	fmt.Println(address2)
	fmt.Println(age2)

	// shorthand declaration
	language := "Golang"
	fmt.Println(language)
}
