package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	https://yukicoder.me/problems/no/207
	2020/11/22
*/

func main() {
	var minNumber int
	var maxNumber int
	var numberList []int
	var worldNabeList []int

	// sample
	minNumber = 1
	maxNumber = 30

	for i := minNumber; i <= maxNumber; i++ {
		numberList = append(numberList, i)
	}

	for _, n := range numberList {
		if isWorldNabe(n) {
			worldNabeList = append(worldNabeList, n)
		}
	}

	fmt.Println("numberList", numberList)
	fmt.Println("worldNabeList", worldNabeList)
}

func isWorldNabe(number int) bool {
	if number%3 == 0 {
		return true
	}

	strNumber := strconv.Itoa(number)
	return strings.Contains(strNumber, "3")
}
