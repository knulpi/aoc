package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func map_(slice []string, mapFn func(string) string) []string {
	mapped := make([]string, len(slice))
	for i, elem := range slice {
		mapped[i] = mapFn(elem)
	}
	return mapped
}

func includes(slice []string, value string) bool {
	for _, s := range slice {
		if value == s {
			return true
		}
	}
	return false
}

func includesInt(slice []int, value int) bool {
	for _, s := range slice {
		if value == s {
			return true
		}
	}
	return false
}

func union(a []string, b []string) []string {
	u := a

	for _, s := range b {
		if !includes(u, s) {
			u = append(u, s)
		}
	}

	return u
}

func intersection(a []string, b []string) []string {
	var i []string

	for _, s := range a {
		if includes(b, s) {
			i = append(i, s)
		}
	}

	return i
}

func subtract(a []string, b []string) []string {
	var diff []string

	for _, s := range a {
		if !includes(b, s) {
			diff = append(diff, s)
		}
	}

	return diff
}

func symDiff(a []string, b []string) []string {
	diffA := subtract(a, b)
	diffB := subtract(b, a)
	return append(diffA, diffB...)
}

type Digit struct {
	settings map[string]bool
}

func NewDigit(code string) *Digit {
	d := Digit{}
	d.settings = map[string]bool{}
	for _, char := range strings.Split(code, "") {
		d.settings[char] = true
	}
	return &d
}

func (d *Digit) length() int {
	return len(d.settings)
}

func (d *Digit) toStringSlice() []string {
	var r []string
	for s := range d.settings {
		r = append(r, s)
	}
	return r
}

func (d *Digit) toString() string {
	r := ""
	for s := range d.settings {
		r += s
	}
	return r
}

type DigitKeeper struct {
	digits        []*Digit
	knownMappings map[int]*Digit
}

func NewDigitKeeper() *DigitKeeper {
	dk := DigitKeeper{}
	dk.digits = []*Digit{}
	dk.knownMappings = map[int]*Digit{}
	return &dk
}

func (dk *DigitKeeper) Add(digit *Digit) {
	dk.digits = append(dk.digits, digit)
}

func (dk *DigitKeeper) GetMapping(number int) (*Digit, bool) {
	val, ok := dk.knownMappings[number]
	return val, ok
}

func (dk *DigitKeeper) SetMapping(number int, digit *Digit) {
	dk.knownMappings[number] = digit
}

func (dk *DigitKeeper) GetWithLength(length int) []*Digit {
	var got []*Digit

	for _, digit := range dk.digits {
		if digit.length() == length {
			got = append(got, digit)
		}
	}

	return got
}

func GetDigitFromCodeWithMapping(code []string, mapping map[string]string) int {
	inverseMapping := map[string]string{}
	for org, fake := range mapping {
		inverseMapping[fake] = org
	}

	//1
	if len(code) == 2 {
		return 1
	}

	//7
	if len(code) == 3 {
		return 7
	}

	//4
	if len(code) == 4 {
		return 4
	}

	//8
	if len(code) == 7 {
		return 8
	}

	//0, 6, 9
	if len(code) == 6 {
		if !includes(code, mapping["d"]) {
			return 0
		}
		if includes(code, mapping["e"]) {
			return 6
		}
		if includes(code, mapping["c"]) {
			return 9
		}
		panic("")
	}

	//2, 3, 5
	if len(code) == 5 {
		if includes(code, mapping["e"]) {
			return 2
		}
		if includes(code, mapping["b"]) {
			return 5
		}
		return 3
	}

	panic("")
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

	total := 0

	for _, line := range inputLines {
		parts := strings.Split(line, " | ")
		inputValues := strings.Split(parts[0], " ")
		outputValues := strings.Split(parts[1], " ")

		dk := NewDigitKeeper()

		for _, value := range inputValues {
			digit := NewDigit(value)
			dk.Add(digit)
		}

		letterMappings := map[string]string{}

		// get A
		d1 := dk.GetWithLength(2)[0]
		d7 := dk.GetWithLength(3)[0]

		lA := subtract(d7.toStringSlice(), d1.toStringSlice())

		letterMappings["a"] = lA[0]

		// get E
		d4 := dk.GetWithLength(4)[0]
		d069 := dk.GetWithLength(6)
		//var d0 *Digit
		var d6 *Digit
		var d9 *Digit

		for _, digit := range d069 {
			if len(union(digit.toStringSlice(), d4.toStringSlice())) == 6 {
				d9 = digit
			} else if len(union(d1.toStringSlice(), digit.toStringSlice())) == 6 {
				//d0 = digit
			} else if len(union(d1.toStringSlice(), digit.toStringSlice())) != 6 {
				d6 = digit
			} else {
				panic("wrong branch")
			}
		}

		//fmt.Println(d0)

		sub4a := subtract(d6.toStringSlice(), d4.toStringSlice())
		sub4b := subtract(d9.toStringSlice(), d4.toStringSlice())

		lE := symDiff(sub4a, sub4b)
		letterMappings["e"] = lE[0]

		// get c
		d235 := dk.GetWithLength(5)

		var d2 *Digit
		//var d3 *Digit
		//var d5 *Digit

		var lC []string
		var lF []string

		for _, digit := range d235 {
			if len(intersection(lE, digit.toStringSlice())) == 1 {
				d2 = digit
				lC = intersection(digit.toStringSlice(), d1.toStringSlice())
			} else if len(intersection(digit.toStringSlice(), d1.toStringSlice())) == 1 {
				//d5 = digit
				lF = intersection(digit.toStringSlice(), d1.toStringSlice())
			} else {
				//d3 = digit
			}
		}

		//fmt.Println(d3, d5)

		letterMappings["c"] = lC[0]
		letterMappings["f"] = lF[0]

		// get b
		lB := subtract(subtract(d4.toStringSlice(), d1.toStringSlice()), d2.toStringSlice())
		letterMappings["b"] = lB[0]

		// get d
		lD := intersection(subtract(d4.toStringSlice(), d1.toStringSlice()), d2.toStringSlice())

		letterMappings["d"] = lD[0]

		// get g
		lG := subtract(
			[]string{"a", "b", "c", "d", "e", "f", "g"}, []string{
				lA[0], lB[0], lC[0],
				lD[0],
				lE[0],
				lF[0],
			},
		)

		letterMappings["g"] = lG[0]

		sum := 0

		for i, value := range outputValues {
			sum += int(math.Pow(10, float64(len(outputValues)-i-1))) *
				GetDigitFromCodeWithMapping(strings.Split(value, ""), letterMappings)
		}

		total += sum
	}

	fmt.Println(total)
}
