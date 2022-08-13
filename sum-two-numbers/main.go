package main

import "fmt"

func main() {
	var firstNum, secondNum int
	fmt.Scanf("%d %d", &firstNum, &secondNum)
	if firstNum < secondNum {
		fmt.Printf("%d %d", firstNum, secondNum)
		return
	}
	fmt.Printf("%d %d", secondNum, firstNum)
}
