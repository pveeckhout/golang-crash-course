package main

import "fmt"

var isCool = false

func main() {
	var name = "Pieter"
	var age int64 = 32
	isCool := true
	size := 1.375
	multi1, multi2 := "multi1", "multi2"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(size)
	fmt.Println(isCool)
	fmt.Println(multi1, multi2)
	fmt.Printf("type is %T", age)

}
