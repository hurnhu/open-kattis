package main

import (
	"fmt"
	"strings"
)

func main() {
	var loopCount int
	fmt.Scanf("%d\n", &loopCount)
	for i := 0; i < loopCount; i++ {
		var num string
		fmt.Scanf("%s\n", &num)
		fmt.Printf("%d\n", len(strings.Split(num, "")))
	}
}
