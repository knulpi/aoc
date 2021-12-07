package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func getMajorityBits(input []string, digit int) (string, error) {
	counter := 0
	for _, s := range input {
		d := strings.Split(s, "")[digit]

		if d == "1" {
			counter++
		} else if d == "0" {
			counter--
		} else {
			log.Fatalln("invalid value")
		}
	}

	if counter == 0 {
		return "", fmt.Errorf("counter was nil")
	} else if counter > 0 {
		return "1", nil
	} else {
		return "0", nil
	}
}

func getMinorityBits(input []string, digit int) (string, error) {
	maj, err := getMajorityBits(input, digit)
	if err != nil {
		return "", err
	} else {
		if maj == "1" {
			return "0", nil
		} else {
			return "1", nil
		}
	}
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

	worker := inputLines

	digit := 0
	for len(worker) != 1 {
		var newWorker []string

		maj, err := getMajorityBits(worker, digit)
		if err != nil {
			maj = "1"
		}

		for _, line := range worker {
			digits := strings.Split(line, "")
			if digits[digit] == maj {
				newWorker = append(newWorker, line)
			}
		}

		worker = newWorker

		digit++
	}

	first := worker[0]

	worker = inputLines

	digit = 0
	for len(worker) != 1 {
		var newWorker []string

		min, err := getMinorityBits(worker, digit)
		if err != nil {
			min = "0"
		}

		for _, line := range worker {
			digits := strings.Split(line, "")
			if digits[digit] == min {
				newWorker = append(newWorker, line)
			}
		}

		worker = newWorker

		digit++
	}

	second := worker[0]

	s1 := 0
	for digit, value := range strings.Split(first, "") {
		if value == "1" {
			s1 += int(math.Pow(2, float64(len(first)-digit-1)))
		}
	}

	s2 := 0
	for digit, value := range strings.Split(second, "") {
		if value == "1" {
			s2 += int(math.Pow(2, float64(len(second)-digit-1)))
		}
	}

	fmt.Println(s1, s2, s1*s2)

}
