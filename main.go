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

// takes the string, and returns the average number from all the numbers
// complexity Easy, time O(N)
func AverageNumber(str string) float32 {
	words := strings.Split(str, "-")
	numCount := 0
	sum := 0
	for _, word := range words {
		if val, err := strconv.Atoi(word); err == nil {
			numCount++
			sum += val
		}
	}
	return float32(sum) / float32(numCount)
}

// takes the string, and returns a text that is composed from all the text words separated by spaces,
// e.g. story called for the string 1-hello-2-world should return text: "hello world"
// complexity easy, time O(N)
func WholeStory(str string) string {
	words := strings.Split(str, "-")
	story := ""
	for _, word := range words {
		if _, err := strconv.Atoi(word); err == nil {
			continue
		} else {
			story += word + " "
		}
	}
	return strings.TrimSpace(story)
}
