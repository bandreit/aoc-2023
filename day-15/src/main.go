package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	input := strings.ReplaceAll(string(dat), "\n", "")
	pieces := strings.Split(input, ",")
	sumOfNewPieces := 0

	for _, piece := range pieces {
		newPiece := 0

		for _, r := range piece {
			newPiece += int(r)
			newPiece *= 17
			newPiece %= 256
		}

		// fmt.Println(newPiece)
		sumOfNewPieces += newPiece
	}

	fmt.Println(sumOfNewPieces)
}
