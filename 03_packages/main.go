package main

import (
	"fmt"
	"github.com/pveeckhout/golang-crash-course/03_packages/strutil"
	"math"
)

func main() {
	fmt.Println(math.Max(0.8, 1.8))
	fmt.Println(math.Sqrt(25))
	fmt.Println(math.Sqrt(25))
	fmt.Println(strutil.Reverse("reverse reverse"))
}
