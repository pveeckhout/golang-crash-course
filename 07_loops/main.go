package main

import "fmt"

func main() {
	//long
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//short
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	FizzBuzz(1,100,3,5)
}

func FizzBuzz(start, end, divisor1, divisor2 int) {
	for k := start; k <= end; k++ {
		divisible1 := k%divisor1 == 0
		divisible2 := k%divisor2 == 0
		if divisible1 {
			fmt.Print("Fizz")
		}
		if divisible2 {
			fmt.Print("Buzz")
		}
		if divisible1 || divisible2 {
			fmt.Println()
		} else {
			fmt.Println(k)
		}
	}
}
