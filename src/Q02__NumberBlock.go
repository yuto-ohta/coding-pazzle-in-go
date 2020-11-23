package main

import (
	"fmt"
	"sort"
)

/*
	https://yukicoder.me/problems/no/5
	2020/11/22
 */

func main() {
	var boxWidth int
	var stoneList []int

	// sample
	boxWidth = 30
	stoneList = []int{3,6,3,6,4,2,4,6,8,6,5,3,5,7,1,4,6}

	sort.Ints(stoneList)

	sumStoneWidth := 0
	counter := 0
	for _, stone := range stoneList {
		if sumStoneWidth + stone > boxWidth {
			break
		}
		sumStoneWidth += stone
		counter ++
	}

	fmt.Println("sumStoneWidth", sumStoneWidth)
	fmt.Println("counter", counter)
	fmt.Println("boxWidth", boxWidth)
	fmt.Println("stoneList", stoneList)
}
