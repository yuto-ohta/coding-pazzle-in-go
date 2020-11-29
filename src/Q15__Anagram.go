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

	listStringA := []string{}
	listStringB := []string{}

	for _, char := range stringA {
		listStringA = append(listStringA, string(char))
	}
	for _, char := range stringB {
		listStringB = append(listStringB, string(char))
	}

	sort.Strings(listStringA)
	sort.Strings(listStringB)

	answer := false
	if reflect.DeepEqual(listStringA, listStringB) {
		answer = true
	}

	fmt.Println(listStringA)
	fmt.Println(listStringB)
	fmt.Println(answer)
}
