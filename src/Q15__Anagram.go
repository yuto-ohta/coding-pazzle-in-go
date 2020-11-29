package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
	https://yukicoder.me/problems/no/69
	2020/11/29
*/

func main() {
	var stringA string
	var stringB string

	// sample
	stringA = "apple"
	stringB = "elppa"

	stringASlice := []string{}
	stringBSlice := []string{}

	for _, char := range stringA {
		stringASlice = append(stringASlice, string(char))
	}
	for _, char := range stringB {
		stringBSlice = append(stringBSlice, string(char))
	}

	sort.Strings(stringASlice)
	sort.Strings(stringBSlice)

	answer := false
	if reflect.DeepEqual(stringASlice, stringBSlice) {
		answer = true
	}

	fmt.Println(stringASlice)
	fmt.Println(stringBSlice)
	fmt.Println(answer)
}
