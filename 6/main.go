package main

import (
	"fmt"
	"io"
	"log"
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

//func produce(set *[]int, wg *sync.WaitGroup) {
//	for i, fish := range *set {
//		if fish == 0 {
//			(*set)[i] = 6
//			*set = append(*set, 8)
//		} else {
//			(*set)[i]--
//		}
//	}
//	wg.Done()
//}

func main() {
	inputFile, err := os.Open("input")
	if err != nil {
		log.Fatalln(err)
	}

	fileContents, err := io.ReadAll(inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	inputLine := strings.Split(string(fileContents), "\r\n")[0]

	set, err := mapToInts(strings.Split(inputLine, ","))
	if err != nil {
		log.Fatalln(err)
	}

	days := 256

	state := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	for _, fish := range set {
		state[fish]++
	}

	for i := 0; i < days; i++ {
		zeros := state[0]

		state[0] = state[1]
		state[1] = state[2]
		state[2] = state[3]
		state[3] = state[4]
		state[4] = state[5]
		state[5] = state[6]
		state[6] = state[7] + zeros
		state[7] = state[8]
		state[8] = zeros
	}

	sum := 0
	for _, e := range state {
		sum += e
	}

	fmt.Println(sum)
}

// initial fish: double every 7 days, -initial
// 2
// 77989447176
//03620350005
