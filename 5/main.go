package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func map_(slice []int, mapFn func(int) int) []int {
	mapped := make([]int, len(slice))
	for i, elem := range slice {
		mapped[i] = mapFn(elem)
	}
	return mapped
}

func mapToInts(slice []string) ([]int, error) {
	mapped := make([]int, len(slice))
	for i, elem := range slice {
		conv, err := strconv.Atoi(elem)
		if err != nil {
			return nil, err
		}
		mapped[i] = conv
	}
	return mapped, nil
}

func reduce(slice []int, reducer func(prev int, cur int) int, initial int) int {
	prev := initial
	for _, elem := range slice {
		prev = reducer(prev, elem)
	}
	return prev
}

func filter(slice []int, filterFn func(int) bool) []int {
	var mapped []int
	for _, elem := range slice {
		if filterFn(elem) {
			mapped = append(mapped, elem)
		}
	}
	return mapped
}

func insert(grid [][]int, x int, y int) {
	grid[x][y]++
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

	grid := make([][]int, 1000)

	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, line := range inputLines {
		coords := strings.Split(line, " -> ")

		c1, err := mapToInts(strings.Split(coords[0], ","))
		if err != nil {
			log.Fatalln(err)
		}
		c2, err := mapToInts(strings.Split(coords[1], ","))
		if err != nil {
			log.Fatalln(err)
		}

		x := []int{c1[0], c2[0]}
		y := []int{c1[1], c2[1]}

		sort.Ints()

		if line == "" {
			continue
		}
	}
}
