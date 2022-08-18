package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	hand := make(map[string]int)
	inputSplit := strings.Split(input, " ")
	max := 0
	for i := 0; i < len(inputSplit); i++ {
		rank := strings.Split(inputSplit[i], "")
		hand[rank[0]] = hand[rank[0]] + 1
		if max < hand[rank[0]] {
			max = hand[rank[0]]
		}
	}
	fmt.Print(max)
}
