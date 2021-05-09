package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
	https://yukicoder.me/problems/no/227
	2021/05/08
*/

func main() {
	//五枚のカード（1~13の数が割り当てられる）
	var fiveCards []int

	// 標準入力から5つの数字を受け取る
	// ex) "1 2 3 4 5"
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strFiveCards := strings.Fields(scanner.Text())
	if len(strFiveCards) != 5 {
		_, _ = fmt.Fprintln(os.Stderr, "error: 5つの数字を指定する")
	}
	for _, strNum := range strFiveCards {
		num, _ := strconv.Atoi(strNum)
		fiveCards = append(fiveCards, num)
	}
	sort.Ints(fiveCards)

	//1~13のそれぞれの数が
	//	カードの中にいくつ存在するかを数えて、配列に格納する
	var howManyEachNumber [13]int
	counter := 0
	for i := 1; i <= 13; i++ {
		for _, num := range fiveCards {
			if i == num {
				counter++
			}
		}
		howManyEachNumber[i-1] = counter
		counter = 0
	}

	//配列が
	//		3, 2を含んでいればFULL HOUSE
	if containSpecificNum(howManyEachNumber[:], 3) && containSpecificNum(howManyEachNumber[:], 2) {
		fmt.Println("FULL HOUSE")
		return
		//	3を含んでいれば
	} else if containSpecificNum(howManyEachNumber[:], 3) {
		fmt.Println("THREE CARD")
		return
	} else if containSpecificNum(howManyEachNumber[:], 2) {
		firstTwoIndex := getSpecificNumIndex(howManyEachNumber[:], 2)
		//	2, 2を含んでいれば
		if containSpecificNum(howManyEachNumber[firstTwoIndex+1:], 2) {
			fmt.Println("TWO PAIR")
			return
		}
		//	2を含んでいれば
		fmt.Println("ONE PAIR")
	} else {
		//	それ以外は NO HAND
		fmt.Println("NO HAND")
	}
}

func containSpecificNum(target []int, number int) bool {
	if getSpecificNumIndex(target, number) != -1 {
		return true
	}
	return false
}

func getSpecificNumIndex(target []int, number int) int {
	for i, el := range target {
		if el == number {
			return i
		}
	}
	return -1
}
