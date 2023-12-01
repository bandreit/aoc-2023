package main

import (
	"testing"
)

func TestAdd(t *testing.T) {

	type test struct {
		data   string
		answer int
	}

	tests := []test{
		{"xtwone3four", 24},
		{"dusoneeualagsaeightwo", 12},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"sevenine", 79},
		{"eighthree", 83},
	}

	for _, input := range tests {
		line := replaceLettersWithDigits(input.data)
		got := calibrateLine(line)
		want := input.answer

		if got != want {
			t.Errorf("got %d, wanted %d for %s", got, want, input.data)
		}
	}
}
