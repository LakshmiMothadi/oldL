package main

import (
	"fmt"
)

func SubString(s string) int {
	firstElement := 0
	lastElement := make(map[byte]int)
	length := 0

	for i, j := range []byte(s) {
		if index, data := lastElement[j]; data && index >= firstElement {
			firstElement = index + 1
		}
		basicCondition := i - firstElement + 1
		if basicCondition > length {
			length = basicCondition

		}
		lastElement[j] = i
	}
	return length
}
func main() {
	var s string
	fmt.Println("enter string:")
	fmt.Scan(&s)
	fmt.Println(SubString(s))
}

