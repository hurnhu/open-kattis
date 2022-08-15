package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var time string
	display := make([][]string, 0)
	fmt.Scanf("%s", &time)
	splitTime := strings.Split(time, "")
	for i := 0; i < len(splitTime); i++ {
		asInt, _ := strconv.Atoi(splitTime[i])
		binRep := strconv.FormatInt(int64(asInt), 2)
		binRep = fmt.Sprintf("%0*s", 4, binRep)
		splitBin := strings.Split(binRep, "")
		forDiaplay := make([]string, 0)
		for i := 0; i < 4; i++ {
			asInt := 0
			if len(splitBin) > i {
				asInt, _ = strconv.Atoi(splitBin[i])
			}
			if asInt == 1 {
				forDiaplay = append(forDiaplay, "*")
			} else {
				forDiaplay = append(forDiaplay, ".")
			}
		}
		display = append(display, forDiaplay)
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("%s %s    %s %s\n", display[0][i], display[1][i], display[2][i], display[3][i])
	}
}
