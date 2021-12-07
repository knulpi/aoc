package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func getNextSum(numbers []int) (int, error) {
	if len(numbers) < 3 {
		return 0, fmt.Errorf("")
	}

	return numbers[0] + numbers[1] + numbers[2], nil
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

	var inputNums []int

	for _, line := range inputLines {
		if line == "" {
			continue
		}
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln(err)
		}
		inputNums = append(inputNums, lineInt)
	}

	count := 0
	movingSlc := inputNums
	lastSum, err := getNextSum(movingSlc)
	if err != nil {
		log.Fatalln(err)
	}

	for range inputLines {
		newSum, err := getNextSum(movingSlc)
		if err != nil {
			break
		}
		if newSum > lastSum {
			count++
		}
		lastSum = newSum
		movingSlc = movingSlc[1:]
		fmt.Printf("X%vX\n", newSum)
	}

	fmt.Println(count)
}
