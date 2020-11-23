package main

import "fmt"

/*
	https://yukicoder.me/problems/no/480
	2020/11/22
 */

func main() {
	var N int

	// sample
	N = 100

	max := 0
	for i := 1; i <= N; i++ {
		max += i
	}

	fmt.Println(max)
}
