package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T, %T\n", a, b)

	fmt.Println(a * (*b))

	*b = 18
	fmt.Println(a, *b)
}
