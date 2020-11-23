package main

import (
	"fmt"
	"sort"
)

/*
	https://yukicoder.me/problems/no/275
	2020/11/23
*/

func main() {
	var numberList []int

	// sample
	numberList = []int{1, 9, 8, 4}

	sort.Ints(numberList)

	var median float32
	if len(numberList)%2 == 0 {
		median = float32(numberList[(len(numberList)/2)-1]+numberList[len(numberList)/2]) / float32(2)
	} else {
		median = float32(numberList[len(numberList)/2])
	}

	fmt.Println("median", median)
}
