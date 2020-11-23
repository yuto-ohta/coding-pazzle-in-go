package main

import (
	"fmt"
	"strconv"
)

/*
	https://yukicoder.me/problems/no/188
	2020/11/23
*/

func main() {
	var date string

	// sample
	date = "05/23"

	month := date[0:2]
	day := date[3:5]
	strMonth := 0
	sumDayNumber := 0

	num, _ := strconv.Atoi(month)
	strMonth = num

	for _, char := range day {
		num, _ := strconv.Atoi(string(char))
		sumDayNumber += num
	}

	if strMonth == sumDayNumber {
		fmt.Printf("%v is HappyDay!\n", date)
	} else {
		fmt.Printf("%v is Day!\n", date)
	}
}
