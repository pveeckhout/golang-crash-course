package main

import (
	"fmt"
	"strconv"
)

func main() {
	person1 := Person{
		firstName: "Pieter",
		lastName:  "Van Eeckkhout",
		cityName:  "Wanzele",
		gender:    true,
		age:       32,
	}

	fmt.Println(person1)
	fmt.Println(person1.age)

	fmt.Println(person1.greet())
	fmt.Println(person1)
	fmt.Println(person1.age)

	fmt.Println(person1.changer())
	fmt.Println(person1)
	fmt.Println(person1.age)
}

func (person Person) greet() string {
	// this should not update the person object outside of the func
	person.age++
	return "Hello my name is " + person.firstName + " " + person.lastName + " aged " + strconv.Itoa(person.age)
}

func (person *Person) changer() string {
	// this will update the person object outside of the func
	person.age++
	return "Hello my age is " + strconv.Itoa(person.age)
}

type Person struct {
	firstName, lastName, cityName string
	gender                        bool
	age                           int
}
