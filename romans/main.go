package main

import (
	"fmt"
	"math"
)

func main() {
	var miles float64
	fmt.Scanf("%f", &miles)
	fmt.Printf("%d", int(math.RoundToEven(float64(miles*1000)*(float64(5280)/float64(4854)))))
}
