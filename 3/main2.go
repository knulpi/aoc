package main2

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

var storage []int

var storageControl = map[string]func(int){
	"0": func(digit int) {
		for len(storage) < digit+1 {
			storage = append(storage, 0)
		}

		storage[digit]--
	},
	"1": func(digit int) {
		for len(storage) < digit+1 {
			storage = append(storage, 0)
		}

		storage[digit]++
	},
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

	for _, line := range inputLines {
		if line == "" {
			continue
		}

		digits := strings.Split(line, "")

		for digit, value := range digits {
			storageControl[value](digit)
		}
	}

	var gamma, epsilon int

	for digit, value := range storage {
		addVal := int(math.Pow(2, float64(len(storage)-digit-1)))
		if value > 0 {
			gamma += addVal
		} else if value < 0 {
			epsilon += addVal
		} else {
			log.Fatalln("value was null")
		}
	}

	fmt.Println(gamma * epsilon)
}
