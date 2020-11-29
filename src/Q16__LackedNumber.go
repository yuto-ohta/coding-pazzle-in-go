package main

import (
	"fmt"
)

/*
	https://yukicoder.me/problems/no/481
	2020/11/29
*/

func main() {
	var numbers []int
	var target []int

	// sample
	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target = []int{1, 2, 3, 4, 6, 7, 9, 10}

	for _, num := range numbers {
		if !containsInt(target, num) {
			fmt.Println(num)
		}
	}
}

func containsInt(list []int, number int) bool {
	for _, num := range list {
		if num == number {
			return true
		}
	}

	return false
}
