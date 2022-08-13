package main

import "fmt"

func main() {
	var numParts int
	var runningTotal int
	var temp int
	fmt.Scanln(&numParts)
	for i := 0; i < numParts; i++ {
		fmt.Scanf("%d\n", &temp)
		runningTotal += temp
	}
	runningTotal -= numParts - 1
	fmt.Println(runningTotal)
}
