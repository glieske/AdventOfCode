package main

import (
	"fmt"
	"strconv"
	"strings"

	stringmanipulation "AdventOfCode/pkg/StringManipulation"
	typeconvert "AdventOfCode/pkg/TypeConvert"
	"go.uber.org/zap"
)

func part2(input []string) (string, error) {
	var result string

	sum := 0
	for _, s := range input {
		coords := performExtractionWithNumericNames(s)
		sum += coords
	}
	zap.L().Info(fmt.Sprintf("Result P2: %d", sum))
	result = strconv.Itoa(sum)

	return result, nil
}

func performExtractionWithNumericNames(s string) int {
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		leftDigit, correct := getFirstNumber(s[i:])
		if !correct {
			continue
		}
		for y := len(rs) - 1; y >= 0; y-- {
			rightDigit, rCorrect := getLastNumber(s[:y+1])
			if !rCorrect {
				continue
			}
			result := leftDigit*10 + rightDigit
			zap.L().Debug(fmt.Sprintf("Got %s, extracted: %d", s, result))
			return result
		}
	}
	return -1
}

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func getFirstNumber(str string) (int, bool) {
	if stringmanipulation.IsStringRuneANumber(rune(str[0])) {
		return typeconvert.StringRuneToInt(rune(str[0])), true
	}
	for txt, repl := range numbers {
		if strings.HasPrefix(str, txt) {
			return repl, true
		}
	}
	return -1, false
}

func getLastNumber(str string) (int, bool) {
	x := len(str) - 1
	if stringmanipulation.IsStringRuneANumber(rune(str[x])) {
		return typeconvert.StringRuneToInt(rune(str[x])), true
	}
	for txt, repl := range numbers {
		if strings.HasSuffix(str, txt) {
			return repl, true
		}
	}
	return -1, false
}
