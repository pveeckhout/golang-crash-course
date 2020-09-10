package main

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{
		x:      0,
		y:      0,
		radius: 4,
	}

	rectangle := Rectangle{
		width:  4,
		height: 9,
	}

	fmt.Println("circle shape area: ", getArea(circle))
	fmt.Println("rectangle shape area: ", getArea(rectangle))
}

func getArea(shape Shape) float64 {
	return shape.area()
}

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func (circle Circle) area() float64 {
	return math.Pow(circle.radius, 2) * math.Pi
}