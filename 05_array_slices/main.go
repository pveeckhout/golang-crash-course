package main

import "fmt"

func main() {
	//arrays
	var fruitArray [2]string

	fruitArray[0] = "apple"
	fruitArray[1] = "cherry"
	fmt.Println(fruitArray)

	beerArray := [2]string{"westmalle", "chimay"}
	fmt.Println(beerArray)

	//slices
	var fruitSlice []string
	fruitSlice = append(fruitSlice,"apple")
	fruitSlice = append(fruitSlice, "cherry")
	fruitSlice = append(fruitSlice, "banana")
	fmt.Println(fruitSlice)

	beerSlice := []string{"westmalle", "chimay","Rochefort"}
	fmt.Println(beerSlice)

	fmt.Println(len(beerSlice))
	fmt.Println(beerSlice[1:])
	fmt.Println(fruitSlice[0:2])
}
