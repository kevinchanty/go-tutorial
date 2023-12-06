package main

import (
	_ "embed"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var digit = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var digitStrs = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type Match struct {
	value    string
	position int
}

func main() {
	fmt.Println("Part 1:")
	fmt.Println(part1())
	fmt.Println("Part 2:")
	fmt.Println(part2())
}

func part1() int {
	var sum int

	for _, line := range strings.Split(input, "\n") {
		var digitStr string

		for _, char := range line {
			if slices.Contains(digit, char) {
				if digitStr == "" {
					digitStr = fmt.Sprintf("%v%v", string(char), string(char))
				} else {
					digitStr = fmt.Sprintf("%v%v", digitStr[0:1], string(char))
				}
			}
		}

		parsedInt, err := strconv.Atoi(digitStr)
		if err != nil {
			panic(err)
		}
		sum += parsedInt
	}
	return sum
}

func part2() int {
	var sum int

	for _, line := range strings.Split(input, "\n") {

		println(line)

		var matchArr []Match

		for _, digitStr := range digitStrs {
			fistIndex := strings.Index(line, digitStr)
			lastIndex := strings.LastIndex(line, digitStr)

			if fistIndex != -1 {
				matchArr = append(matchArr, Match{
					value:    digitMap[digitStr],
					position: fistIndex,
				})
			}

			if lastIndex != fistIndex {
				matchArr = append(matchArr, Match{
					value:    digitMap[digitStr],
					position: lastIndex,
				})
			}
		}

		for index, char := range line {
			if slices.Contains(digit, char) {
				matchArr = append(matchArr, Match{
					value:    string(char),
					position: index,
				})
			}
		}

		if len(matchArr) == 1 {
			res, _ := strconv.Atoi(fmt.Sprintf("%v%v", matchArr[0].value, matchArr[0].value))
			println(res)

			sum += res
		} else {
			sort.Slice(matchArr, func(i, j int) bool { return matchArr[i].position < matchArr[j].position })
			res, _ := strconv.Atoi(fmt.Sprintf("%v%v", matchArr[0].value, matchArr[len(matchArr)-1].value))
			println(res)
			sum += res
		}
		println(sum)
	}

	return sum
}
