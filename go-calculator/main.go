package main

import "fmt"

func main() {
	fmt.Println("NICK'S GO CALCULATOR 1.0")
	fmt.Println("========================")

	var operation string
	var number1, number2 int

	fmt.Println("What operation would you like to do?")
	fmt.Println("(add, subtract, multiply, divide)")
	fmt.Scanf("%s", &operation)

	fmt.Println("What is the 1st number?")
	fmt.Scanf("%v", &number1)

	fmt.Println("What is the 2nd number?")
	fmt.Scanf("%v", &number2)

	var result int

	switch operation {
	case "multiply":
		result = number1 * number2
	case "add":
		result = number1 + number2
	case "divide":
		result = number1 / number2
	case "subtract":
		result = number1 - number2
	}

	fmt.Println("========RESULT===========")
	fmt.Printf("%d %s %d = %d", number1, operation, number2, result)
}
