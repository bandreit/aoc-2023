package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestBuild2DArrayFromLines(t *testing.T) {
	dat, err := os.ReadFile("../test_input.txt")
	check(err)

	input := string(dat)
	lines := strings.Split(input, "\n")

	twoDArr := build2DArrayFromLines(input)

	got := len(twoDArr)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = len(twoDArr[0])
	want = 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	for i, line := range lines {
		for j, char := range line {
			got := twoDArr[i][j]
			want := char

			if got != want {
				t.Errorf("got %d, wanted %d", got, want)
			}
		}
	}
}

func TestPrint2DArray(t *testing.T) {
	dat, err := os.ReadFile("../test_input.txt")
	check(err)

	input := string(dat)
	twoDArr := build2DArrayFromLines(input)

	twoDArr = addPaddingTo2DArray(twoDArr, 'x')
	for i, line := range twoDArr {
		fmt.Println(i, string(line))
	}
}
