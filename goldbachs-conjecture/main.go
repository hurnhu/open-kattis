package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int
	numToGetSums := make([]int, 0)
	var numToAppend int
	fmt.Scanf("%d\n", &input)
	for i := 0; i < input; i++ {
		fmt.Scanf("%d\n", &numToAppend)
		numToGetSums = append(numToGetSums, numToAppend)
	}
	primes := getPrimeUpTo(getMaxInput(numToGetSums))
	for i := 0; i < len(numToGetSums); i++ {
		permutations := make([]string, 0)
		highest := primes[len(primes)-1]
		for j, num := range primes {
			if highest < num {
				continue
			}
			for _, numTwo := range primes[j:] {
				if num+numTwo == numToGetSums[i] {
					permutations = append(permutations, strconv.Itoa(num)+"+"+strconv.Itoa(numTwo))
					highest = numTwo
				}
			}
		}
		fmt.Printf("%d has %d representation(s)\n", numToGetSums[i], len(permutations))
		for l := 0; l < len(permutations); l++ {
			fmt.Printf("%s\n", permutations[l])
		}
		fmt.Println()
	}

}

func getMaxInput(inputs []int) int {
	max := 0
	for i := 0; i < len(inputs); i++ {
		if max < inputs[i] {
			max = inputs[i]
		}
	}
	return max
}

func getPrimeUpTo(num int) []int {
	toReturn := make([]int, 0)
	for i := 2; i < num; i++ {
		if isPrime(i) {
			toReturn = append(toReturn, i)
		}
	}
	return toReturn
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
