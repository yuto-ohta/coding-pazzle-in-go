package main

import (
	"fmt"
	"strconv"
)

/*
	https://yukicoder.me/problems/no/289
	2020/11/23
*/

func main() {
	var mayBeContainsNumber string

	// sample
	mayBeContainsNumber = "1test34これはなんですのん3333"

	sumNumber := 0
	for _, char := range mayBeContainsNumber {
		num, _ := strconv.Atoi(string(char))
		sumNumber += num
	}

	fmt.Println(sumNumber)
}
