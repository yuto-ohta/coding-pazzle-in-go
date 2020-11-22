package main

import "fmt"

/*
	https://yukicoder.me/problems/no/56
	2020/11/22
*/

func main() {
	var rawValue int
	var consumptionTax int

	// sample
	rawValue = 100
	consumptionTax = 8

	fmt.Println(rawValue + rawValue*consumptionTax/100)
}
