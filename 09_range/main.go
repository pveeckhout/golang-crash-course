package main

import (
	"fmt"
)

func main() {
	ids := []int{8,99,125,4,2,14}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("sum: ", sum)

	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Alice"] = "alice@gmail.com"
	emails["Karen"] = "karen@gmail.com"
	emails["Pieter"] = "pieter@gmail.com"

	for key, value := range emails {
		fmt.Printf("key: %s, val: %s\n", key, value)
	}
}