package main

import (
	_ "embed"
	"testing"

	"AdventOfCode/internal/services"
	bytesconvert "AdventOfCode/pkg/TypeConvert"
)

//go:embed input/test-input-p1.txt
var inputPart1Test []byte

//go:embed input/test-input-p1-result.txt
var inputPart1TestResult string

//go:embed input/test-input-p2.txt
var inputPart2Test []byte

//go:embed input/test-input-p2-result.txt
var inputPart2TestResult string

//go:embed input/input-p1-result.txt
var inputPart1Result string

//go:embed input/input-p2-result.txt
var inputPart2Result string

func Test_part1(t *testing.T) {
	t.Run(
		"Part 1", func(t *testing.T) {
			input1 := bytesconvert.ArrayOfBytesToArrayOfStrings(inputPart1Test)
			input2 := bytesconvert.ArrayOfBytesToArrayOfStrings(inputPart2Test)
			task := services.NewSimpleAdventTask(part1, part2)
			task.SetPart1Input(input1).SetPart1ExpectedResult(inputPart1TestResult)
			task.SetPart2Input(input2).SetPart2ExpectedResult(inputPart2TestResult)
			err := task.RunPart1()
			if err != nil {
				t.Errorf("Error: %v", err)
			}
		},
	)
}
func Test_part1_full(t *testing.T) {
	t.Run(
		"Part 1 - full", func(t *testing.T) {
			input1 := bytesconvert.ArrayOfBytesToArrayOfStrings(inputPart1)
			input2 := bytesconvert.ArrayOfBytesToArrayOfStrings(inputPart2)
			task := services.NewSimpleAdventTask(part1, part2)
			task.SetPart1Input(input1).SetPart1ExpectedResult(inputPart1Result)
			task.SetPart2Input(input2).SetPart2ExpectedResult(inputPart2Result)
			err := task.RunPart1()
			if err != nil {
				t.Errorf("Error: %v", err)
			}
		},
	)
}

func Test_part2(t *testing.T) {
	t.Run(
		"Part 2", func(t *testing.T) {
			input1 := bytesconvert.ArrayOfBytesToArrayOfStrings(inputPart1Test)
			input2 := bytesconvert.ArrayOfBytesToArrayOfStrings(inputPart2Test)
			task := services.NewSimpleAdventTask(part1, part2)
			task.SetPart1Input(input1).SetPart1ExpectedResult(inputPart1TestResult)
			task.SetPart2Input(input2).SetPart2ExpectedResult(inputPart2TestResult)
			err := task.RunPart2()
			if err != nil {
				t.Errorf("Error: %v", err)
			}
		},
	)
}

func Test_part2_full(t *testing.T) {
	t.Run(
		"Part 2 - full", func(t *testing.T) {
			input1 := bytesconvert.ArrayOfBytesToArrayOfStrings(inputPart1)
			input2 := bytesconvert.ArrayOfBytesToArrayOfStrings(inputPart2)
			task := services.NewSimpleAdventTask(part1, part2)
			task.SetPart1Input(input1).SetPart1ExpectedResult(inputPart1Result)
			task.SetPart2Input(input2).SetPart2ExpectedResult(inputPart2Result)
			err := task.RunPart2()
			if err != nil {
				t.Errorf("Error: %v", err)
			}
		},
	)
}
