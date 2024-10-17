package main

import (
	"fmt"
	"strconv"

	stringmanipulation "AdventOfCode/pkg/StringManipulation"
	"go.uber.org/zap"
)

func part1(input []string) (string, error) {
	var result string
	sum := 0
	for _, s := range input {
		coords, e := performExtraction(s)
		if e != nil {
			return "", e
		}
		sum += coords
	}
	zap.L().Info(fmt.Sprintf("Result P1: %d", sum))
	result = strconv.Itoa(sum)

	return result, nil
}

func performExtraction(line string) (int, error) {
	firstDigit, e := stringmanipulation.GetFirstNumberFromString(line)
	if e != nil {
		return -1, e
	}
	lastDigit, e := stringmanipulation.GetLastNumberFromString(line)
	if e != nil {
		return -1, e
	}
	return firstDigit*10 + lastDigit, nil
}
