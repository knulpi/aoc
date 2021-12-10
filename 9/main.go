package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
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

func swarm(
	ds [][]int, hadSwarm *[][]bool, basinSize *int, mut *sync.Mutex, wg *sync.WaitGroup,
	r, c int,
) {
	defer wg.Done()

	if ds[r][c] == 9 {
		return
	}

	mut.Lock()
	// return if swarming had been done on this point
	if (*hadSwarm)[r][c] {
		mut.Unlock()
		return
	} else {
		(*hadSwarm)[r][c] = true
		mut.Unlock()
	}

	mut.Lock()
	defer mut.Unlock()

	*basinSize += 1

	if r > 0 {
		wg.Add(1)
		go swarm(ds, hadSwarm, basinSize, mut, wg, r-1, c)
	}
	if r < len(ds)-1 {
		wg.Add(1)
		go swarm(ds, hadSwarm, basinSize, mut, wg, r+1, c)
	}
	if c > 0 {
		wg.Add(1)
		go swarm(ds, hadSwarm, basinSize, mut, wg, r, c-1)
	}
	if c < len(ds[r])-1 {
		wg.Add(1)
		go swarm(ds, hadSwarm, basinSize, mut, wg, r, c+1)
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

	var ds [][]int

	//var lowValues [][]int

	for _, line := range inputLines {
		local, err := mapToInts(strings.Split(line, ""))
		if err != nil {
			log.Fatalln(err)
		}
		ds = append(ds, local)
	}

	var hadSwarm [][]bool
	var basins []int

	for _, cs := range ds {
		hadSwarm = append(hadSwarm, make([]bool, len(cs)))
	}

	for ri, cs := range ds {
		for ci, v := range cs {
			if v != 9 && !hadSwarm[ri][ci] {
				var wg sync.WaitGroup
				basinSize := 0
				wg.Add(1)
				swarm(ds, &hadSwarm, &basinSize, &sync.Mutex{}, &wg, ri, ci)
				wg.Wait()
				basins = append(basins, basinSize)
			}
		}
	}

	//for ri, cs := range ds {
	//	for ci, v := range cs {
	//		usedOperator := "|"
	//		if hadSwarm[ri][ci] && v == 9 {
	//			usedOperator = "!"
	//		} else if hadSwarm[ri][ci] {
	//			usedOperator = "."
	//		}
	//
	//		fmt.Printf("%v", usedOperator)
	//	}
	//	fmt.Print("\n")
	//}

	sort.Ints(basins)
	//
	//fmt.Println(basins)
	//
	fmt.Println(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
	//fmt.Println(hadSwarm)

	//for ri, cs := range ds {
	//	for ci, v := range cs {
	//if ri > 0 {
	//	if ds[ri-1][ci] <= v {
	//		continue
	//	}
	//}
	//if ri < len(ds)-1 {
	//	if ds[ri+1][ci] <= v {
	//		continue
	//	}
	//}
	//if ci > 0 {
	//	if ds[ri][ci-1] <= v {
	//		continue
	//	}
	//}
	//if ci < len(cs)-1 {
	//	if ds[ri][ci+1] <= v {
	//		continue
	//	}
	//}
	//
	//ins := []int{ri, ci, v}
	//
	//lowValues = append(lowValues, ins)
	//	}
	//}
	//
	//sum := 0
	//
	//for _, value := range lowValues {
	//	sum += value[2] + 1
	//}
	//
	//fmt.Println(sum)
	//
	//fmt.Println(lowValues)
}
