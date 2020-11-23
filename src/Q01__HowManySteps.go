package main

import (
	"fmt"
)

/*
	https://yukicoder.me/problems/no/46
	2020/11/22
 */

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	allSteps := 0

	if b % a == 0 {
		allSteps = b / a
	} else {
		allSteps = b / a + 1
	}

	fmt.Println(allSteps)
}
