package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Label struct {
	label       string
	focalLength int
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	input := strings.ReplaceAll(string(dat), "\n", "")
	pieces := strings.Split(input, ",")
	// sumOfNewPieces := part1(pieces)
	// fmt.Println(sumOfNewPieces)

	boxes := make(map[int][]Label)
	fmt.Println(boxes)

	re := regexp.MustCompile("[a-zA-Z]+")
	for _, label := range pieces {
		lettersInLabel := re.FindString(label)
		box := hash(lettersInLabel)
		operationInfo := label[len(label)-2:]

		if operationInfo[len(operationInfo)-1] == '-' {
			handleMinus(box, label, boxes)
		} else {
			handleEquals(operationInfo, boxes, box, label)
		}
	}

	s := 0
	for box, labels := range boxes {
		fmt.Printf("Box %d:\n", box)
		// 		One plus the box number of the lens in question.
		// The slot number of the lens within the box: 1 for the first lens, 2 for the second lens, and so on.
		// The focal length of the lens.
		for i, label := range labels {
			fmt.Printf("\t%s\n", label.label)
			s += (1 + box) * (i + 1) * label.focalLength
		}
	}

	fmt.Println(s)
}

func handleEquals(operationInfo string, boxes map[int][]Label, box int, label string) {
	focalLength := string(operationInfo[len(operationInfo)-1])
	intFocalLength, err := strconv.Atoi(focalLength)
	if err != nil {
		panic(err)
	}

	newLabel := Label{
		label:       label[:len(label)-2],
		focalLength: intFocalLength,
	}
	// if the box exists
	if labels, ok := boxes[box]; ok {
		// replace the label if it exists
		isLensInTheBox := slices.ContainsFunc(labels, func(label Label) bool {
			return label.label == newLabel.label
		})

		if isLensInTheBox {
			for i, label := range labels {
				if label.label == newLabel.label {
					labels[i].focalLength = newLabel.focalLength
				}
			}
		} else {
			// append the label if it doesn't exist
			boxes[box] = append(boxes[box], newLabel)
		}
	} else {
		// create the box if it doesn't exist
		boxes[box] = []Label{newLabel}
	}
}

func handleMinus(box int, label string, boxes map[int][]Label) {
	if labels, ok := boxes[box]; ok {
		newLabel := Label{
			label:       label[:len(label)-1],
			focalLength: -1,
		}

		if slices.ContainsFunc(labels, func(label Label) bool {
			return label.label == newLabel.label
		}) {
			for i, label := range labels {
				if label.label == newLabel.label {
					labels = slices.Delete(labels, i, i+1)
				}
			}

			boxes[box] = labels
		}
	}
}

func part1(pieces []string) int {
	sumOfNewPieces := 0

	for _, piece := range pieces {
		newPiece := 0

		newPiece = hash(piece)

		sumOfNewPieces += newPiece
	}
	return sumOfNewPieces
}

func hash(piece string) int {
	newPiece := 0

	for _, r := range piece {
		newPiece += int(r)
		newPiece *= 17
		newPiece %= 256
	}
	return newPiece
}
