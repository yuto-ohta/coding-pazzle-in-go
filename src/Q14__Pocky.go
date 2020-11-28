package main

import "fmt"

/*
	https://yukicoder.me/problems/no/63
	2020/11/28
*/

func main() {
	var pockyLength int
	var eachOneStroke int

	// sample
	pockyLength = 90
	eachOneStroke = 20

	oneTurnStroke := eachOneStroke * 2
	times := pockyLength / oneTurnStroke
	leftLength := pockyLength % oneTurnStroke

	if leftLength == 0 {
		times--
	}

	fmt.Println(eachOneStroke * times)
}
