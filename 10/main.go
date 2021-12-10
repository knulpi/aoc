package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func includes(slice []string, value string) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}
	return false
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

	//scoreLookup := map[string]int{
	//	")": 3,
	//	"]": 57,
	//	"}": 1197,
	//	">": 25137,
	//}

	charClasses := map[string]int{
		")": 1,
		"(": 1,
		"]": 2,
		"[": 2,
		"}": 3,
		"{": 3,
		">": 4,
		"<": 4,
	}

	openingChars := []string{
		"(",
		"[",
		"{",
		"<",
	}

	closingChars := []string{
		")",
		"]",
		"}",
		">",
	}

	var closingLists [][]string

	for _, line := range inputLines {
		var stack []string
		isCorrupt := false

		for _, char := range strings.Split(line, "") {
			if includes(openingChars, char) {
				stack = append(stack, char)
			}
			if includes(closingChars, char) {
				var popped string
				popped, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if charClasses[popped] != charClasses[char] {
					isCorrupt = true
					break
				}
			}
		}

		if !isCorrupt {
			fmt.Println(stack)
			var closingList []string
			for i := len(stack) - 1; i >= 0; i-- {
				char := stack[i]
				charClass := charClasses[char]
				closingChar := closingChars[charClass-1]
				closingList = append(closingList, closingChar)
			}
			closingLists = append(closingLists, closingList)
		}
	}

	scoreAddition := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	var scores []int

	for _, closingList := range closingLists {
		localScore := 0

		for _, closingChar := range closingList {
			localScore *= 5
			localScore += scoreAddition[closingChar]
		}

		scores = append(scores, localScore)
	}

	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(len(scores))

	halfI := len(scores) / 2

	fmt.Println(scores[halfI])

}
