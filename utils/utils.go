package utils

import "fmt"

func SendMoney() {
	println("sending money")
}

func Collections() {
	var array1 [5]int
	var slice1 []int
	var map1 map[string]string

	array1[0] = 99
	// slice1[0] = 88
	// map1["foo"] = "bar"

	fmt.Println(array1)
	fmt.Println(slice1)
	fmt.Println(map1)
}
