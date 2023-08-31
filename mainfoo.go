package main

import (
	"fmt"
	"nick/learning/go/utils"
)

func main() {
	print("Hello")
	printFloat()

	utils.SendMoney()
	utils.Collections()

	var age = 22
	birthday(&age)
	fmt.Println(age) //=> 22, not 23
}

func printNum() {
	num := 123
	println(num)
}

func printFloat() {
	num := 123.0
	println(num)
}

func birthday(age *int) {
	*age += 1
}
