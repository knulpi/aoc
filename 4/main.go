package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Sequence struct {
	base    []string
	crossed int
}

func (s *Sequence) Cross(nr string) {
	if
}

type Puzzle struct {
	rows [][]string
}

func (p *Puzzle) CheckNr(nr string) {

}

func (p *Puzzle) CheckComplete() bool {

}

func MapToInts(input []string) ([]int, error) {
	var output []int
	for _, s := range input {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		output = append(output, i)
	}
	return output, nil
}

func main() {
	inputFile, err := os.Open("input")
	if err != nil {
		log.Fatalln(err)
	}

	fileContents, err := io.ReadAll(inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	inputLines := strings.Split(string(fileContents), "\r\n")

	gameInput := strings.Split(inputLines[0], ",")

	inputLines = inputLines[2:]

	var puzzles [][][]int

	var currentPuzzle [][]int

	for _, line := range inputLines {
		if line == "" {
			var newPuzzle [][]int

			puzzles = append(puzzles, currentPuzzle)
			currentPuzzle = newPuzzle

			continue
		}

		strings.ReplaceAll(line, "  ", " ")
		numbers, err := MapToInts(strings.Split(line, " "))
		if err != nil {
			log.Fatalln(err)
		}

		currentPuzzle = append(currentPuzzle, numbers)
	}

}
