package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var horizontal, depth, aim int

var cmds = map[string]func(int){
	"forward": func(amount int) {
		horizontal += amount
		depth += aim * amount
	},
	"down": func(amount int) {
		aim += amount
	},
	"up": func(amount int) {
		aim -= amount
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

		parts := strings.Split(line, " ")

		cmd := parts[0]
		amount, err := strconv.Atoi(parts[1])

		if err != nil {
			log.Fatalln(err)
		}

		cmds[cmd](amount)
	}

	fmt.Println(depth * horizontal)
}
