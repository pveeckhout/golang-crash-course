package main

import "fmt"

func main() {
	emails := make(map[string]string)
	twatter := map[string]string{"Bob":"bobinator","karen":"IwannaTalkToYourManager","pieter":"pveeckhout"}

	emails["Bob"] = "bob@gmail.com"
	emails["Alice"] = "alice@gmail.com"
	emails["Karen"] = "karen@gmail.com"
	emails["Pieter"] = "pieter@gmail.com"

	fmt.Println(emails)
	fmt.Println(twatter)
	fmt.Println(len(emails))
	fmt.Println(emails["Karen"])

	delete(emails, "Karen")

	fmt.Println(emails)
	fmt.Println(len(emails))
}
