package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

// returns four things:
// -the shortest word
// -the longest word
// -the average word length
// -the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
// complexity Medium, time O(N)
func StoryStats(str string) (string, string, float32, []string) {
	words := strings.Split(str, "-")
	var shortestWord string
	var longestWord string
	numWords := 0
	totalLenWords := 0
	var listWord []string
	for _, word := range words {
		if _, err := strconv.Atoi(word); err == nil {
			continue
		} else {
			if shortestWord == "" || len(word) < len(shortestWord) {
				shortestWord = word
			}
			if longestWord == "" || len(word) > len(longestWord) {
				longestWord = word
			}
			numWords++
			totalLenWords += len(word)
		}
	}
	avgWordLength := float32(totalLenWords) / float32(numWords)
	avgWordLengthRoundedUp := int(avgWordLength) + 1
	avgWordLengthRoundedDown := int(avgWordLength)
	for _, word := range words {
		if _, err := strconv.Atoi(word); err == nil {
			continue
		} else {
			if len(word) == avgWordLengthRoundedDown || len(word) == avgWordLengthRoundedUp {

				listWord = append(listWord, word)
			}
		}
	}
	return shortestWord, longestWord, avgWordLength, listWord
}

// Write a generate function, that takes boolean flag and generates random correct strings if the parameter is true
// and random incorrect strings if the flag is false.
// complexity Medium, time O(N)
func generate(isCorrect bool) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	numItems := r1.Intn(10) + 2
	result := ""
	for i := 0; i < numItems; i++ {
		result += fmt.Sprintf("%d-%s-", r1.Intn(1000), RandStringBytes(1 + r1.Intn(10)))
	}
	if isCorrect {
		return strings.Trim(result, "-")
	} else {
		return result + "----"
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	fmt.Println(generate(true))
	fmt.Println(generate(false))
}
