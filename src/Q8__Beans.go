package main

import (
	"fmt"
)

/*
	https://yukicoder.me/problems/no/143
	2020/11/22
*/

func main() {
	var beansPerOneSet int
	var howManySets int
	var family []int

	// sample
	beansPerOneSet = 2
	howManySets = 10
	family = []int{5, 5, 2}

	sumBeans := beansPerOneSet * howManySets
	necessaryBeans := 0
	leftBeans := -1

	for _, member := range family {
		necessaryBeans += member
	}

	if necessaryBeans <= sumBeans {
		leftBeans = sumBeans - necessaryBeans
	}

	fmt.Println(leftBeans)
}
