package main

import (
	"strconv"
	"strings"
)

// takes the string as an input, and returns boolean flag true if the given string complies with the format,
// or false if the string does not comply
// complexity Medium, time O(N)
func TestValidity(str string) bool {
	words := strings.Split(str, "-")
	numCount := 0
	wordCount := 0
	for _, word := range words {
		if word == "" {
			return false
		}
		if _, err := strconv.Atoi(word); err == nil {
			numCount++
		} else {
			wordCount++
		}
		if numCount-wordCount == 1 || numCount == wordCount {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {

}
