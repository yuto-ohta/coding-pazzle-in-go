package main

import "fmt"

/*
	https://yukicoder.me/problems/no/32
	2020/11/23
*/

func main() {
	var coin100 int
	var coin25 int
	var coin1 int

	// sample
	coin100 = 0
	coin25 = 0
	coin1 = 0

	sumMoney := 100*coin100 + 25*coin25 + coin1
	afterExchangedCoin100 := (sumMoney % 1000) / 100
	afterExchangedCoin25 := (sumMoney % 1000) % 100 / 25
	afterExchangedCoin1 := (sumMoney % 1000) % 100 % 25

	sumAfterExchangedCoin := afterExchangedCoin100 + afterExchangedCoin25 + afterExchangedCoin1

	fmt.Println(sumAfterExchangedCoin)
}
