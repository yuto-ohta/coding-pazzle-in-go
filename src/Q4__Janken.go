package main

import "fmt"

/*
	https://yukicoder.me/problems/no/264
	2020/11/22
*/

func main() {
	var myHand Janken
	var enemysHand Janken

	// sample
	myHand = Guu
	enemysHand = Paa

	fmt.Println(myHand.isWin(enemysHand))
}

type Janken int

const (
	Guu Janken = iota
	Tyoki
	Paa
)

func (j Janken) isWin(enemysHand Janken) string {
	var result string

	if j == Guu {
		switch enemysHand {
		case Guu:
			result = "あいこ"
		case Tyoki:
			result = "勝ち"
		case Paa:
			result = "負け"
		}
	} else if j == Tyoki {
		switch enemysHand {
		case Guu:
			result = "負け"
		case Tyoki:
			result = "あいこ"
		case Paa:
			result = "勝ち"
		}
	} else {
		switch enemysHand {
		case Guu:
			result = "勝ち"
		case Tyoki:
			result = "負け"
		case Paa:
			result = "あいこ"
		}
	}
	return result
}
