package main

import "fmt"

func main() {
	fmt.Println(greeting("Pieter"))
	fmt.Println(getSum(8,14))
}

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) (int, string) {
	return num1 + num2, ""
}
