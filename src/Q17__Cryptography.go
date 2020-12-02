package main

import (
	"fmt"
	"strings"
)

/*
	https://yukicoder.me/problems/no/18
	2020/12/02
*/

func main() {
	var encrypted string

	// sample
	encrypted = "ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ"

	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	decypted := ""
	for i, char := range encrypted {
		alphabetPosition := strings.Index(alphabet, string(char))
		decryptedPosition := alphabetPosition - (i + 1)

		if decryptedPosition < 0 {
			decryptedPosition = convertedMinusPosition(len(alphabet), decryptedPosition)
		}

		decryptedChar := alphabet[decryptedPosition]
		decypted += string(decryptedChar)
	}

	fmt.Println(decypted)
}

func convertedMinusPosition(listLen int, minusPosition int) int {
	if minusPosition >= 0 {
		return minusPosition
	}

	if listLen+minusPosition < 0 {
		return convertedMinusPosition(listLen, listLen+minusPosition)
	}

	return listLen + minusPosition
}
