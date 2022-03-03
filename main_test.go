package main_test

import (
	"testing"
	"task_upwork"
)

func TestTestValidity(t *testing.T) {
	table := []struct {
		input   string
		isValid bool
	}{
		{"23-ab-48-caba-56-haha", true},
		{"23-ab-48--56-haha", false},
		{"23-ab-48-ram-caba-56-haha", false},
	}

	for _, testCase := range table {
		isValid := main.TestValidity(testCase.input)
		if isValid != testCase.isValid {
			t.Errorf("testValidity was incorrect for input %s, got: %v, want: %v.", testCase.input, isValid, testCase.isValid)
		}
	}

}

func TestAverageNumber(t *testing.T) {
	table := []struct {
		input string
		avg   float32
	}{
		{"23-ab-48-caba-55-haha", 42.0},
		{"12-ab-34--56-haha", 34.0},
		{"23-ab-48-ram-caba--haha", 35.5},
	}

	for _, testCase := range table {
		avg := main.AverageNumber(testCase.input)
		if avg != testCase.avg {
			t.Errorf("AverageNumber was incorrect for input %s, got: %v, want: %v.", testCase.input, avg, testCase.avg)
		}
	}

}
