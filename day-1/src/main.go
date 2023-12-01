package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numberRegex = regexp.MustCompile("[0-9]")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFirstAndLast(arr []string) (string, string) {
	return arr[0], arr[len(arr)-1]
}

func sum(int1 int, int2 int) int {
	return int1 + int2
}

func calibrateLine(line string) int {
	allDigitsAndNumbers := numberRegex.FindAllString(line, -1)

	if len(allDigitsAndNumbers) == 0 {
		return 0
	}

	first, last := getFirstAndLast(allDigitsAndNumbers)
	lineCalibrationString := first + last
	lineCalibration, err := strconv.Atoi(lineCalibrationString)
	check(err)

	return lineCalibration
}

func replaceLettersWithDigits(line string) string {
	line = strings.ReplaceAll(line, "one", "o1e")
	line = strings.ReplaceAll(line, "two", "t2o")
	line = strings.ReplaceAll(line, "three", "t3e")
	line = strings.ReplaceAll(line, "four", "f4r")
	line = strings.ReplaceAll(line, "five", "f5e")
	line = strings.ReplaceAll(line, "six", "s6x")
	line = strings.ReplaceAll(line, "seven", "s7n")
	line = strings.ReplaceAll(line, "eight", "e8t")
	line = strings.ReplaceAll(line, "nine", "n9e")

	return line
}

func calibrateAllLines(calibrationLines []string, parseLetters bool) int {
	sumOfCalibrationLines := 0

	for _, line := range calibrationLines {
		if parseLetters {
			line = replaceLettersWithDigits(line)
		}

		lineCalibration := calibrateLine(line)
		sumOfCalibrationLines += lineCalibration
	}

	return sumOfCalibrationLines
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	calibrationLines := strings.Split(string(dat), "\n")

	calibrationOnlyDigits := calibrateAllLines(calibrationLines, false)
	calibrationDigitsAndLetters := calibrateAllLines(calibrationLines, true)

	fmt.Println(calibrationOnlyDigits)
	fmt.Println(calibrationDigitsAndLetters)
}
