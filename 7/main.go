package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func max(slice []int) int {
	return reduce(
		slice, func(prev int, cur int) int {
			if cur > prev {
				return cur
			}
			return prev
		}, 0,
	)
}

func gaussianSum(n int) int {
	return n * (n + 1) / 2
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

	inputLine := string(fileContents)

	positions, err := mapToInts(strings.Split(inputLine, ","))
	if err != nil {
		log.Fatalln(err)
	}

	fuelCosts := map[int]int{}

	for i := 0; i < max(positions); i++ {
		for _, position := range positions {
			fuelCosts[i] += gaussianSum(int(math.Abs(float64(position - i))))
		}
	}

	min := fuelCosts[0]
	minI := -1

	for index, totalFuel := range fuelCosts {
		if totalFuel < min {
			minI = index
			min = totalFuel
		}
	}

	fmt.Println(minI, min)
}
