package main

import "fmt"

/*
	https://yukicoder.me/problems/no/388
	2020/11/22
*/

func main() {
	var oneFloorStairs int
	var nowStairs int
	var nowFloor int

	// sample
	oneFloorStairs = 35
	nowStairs = 30

	nowFloor = nowStairs/oneFloorStairs + 1
	fmt.Println(nowFloor)
}
