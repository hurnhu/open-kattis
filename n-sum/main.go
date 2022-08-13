package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numOfInputs int
	var nSum int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numOfInputs, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	inputArr := strings.Split(scanner.Text(), " ")
	for i := 0; i < numOfInputs; i++ {
		num, _ := strconv.Atoi(inputArr[i])
		nSum += num
	}
	fmt.Println(nSum)
}
