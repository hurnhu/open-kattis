package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var output string
	var aFound bool
	fmt.Scanf("%s", &input)
	split := strings.Split(input, "")
	for i := 0; i < len(split); i++ {
		if aFound || strings.ToLower(split[i]) == "a" {
			output += split[i]
			aFound = true
		}
	}
	fmt.Print(output)
}
