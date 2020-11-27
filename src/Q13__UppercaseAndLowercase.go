package main

import (
	"fmt"
	"strings"
)

/*
	https://yukicoder.me/problems/no/163
	2020/11/27
*/

func main() {
	var sentence string

	// sample
	sentence = "This is a PEN"

	converted := ""
	for _, char := range sentence {
		isUpper, err := isUpperCase(string(char))

		if err != nil {
			fmt.Println(err)
			continue
		}

		if isUpper {
			converted += strings.ToLower(string(char))
			continue
		}

		converted += strings.ToUpper(string(char))
	}

	fmt.Println(converted)
}

func isUpperCase(alphabet string) (bool, error) {
	if len(alphabet) != 1 {
		return false, fmt.Errorf("argument has to be 1 word alphabet: %v", alphabet)
	}
	if alphabet != strings.ToUpper(alphabet) {
		return false, nil
	}

	return true, nil
}
