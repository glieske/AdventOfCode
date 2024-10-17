package services

import (
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type AdventTaskPart func(input []string) (string, error)
type AdventTask struct {
	part1               AdventTaskPart
	part2               AdventTaskPart
	part1input          []string
	part2input          []string
	part1expectedResult *string
	part2expectedResult *string
}

func NewAdventTask(
	part1 AdventTaskPart,
	part2 AdventTaskPart,
	part1input []string,
	part2input []string,
	part1expectedResult *string,
	part2expectedResult *string,
) *AdventTask {
	return &AdventTask{
		part1:               part1,
		part2:               part2,
		part1input:          part1input,
		part2input:          part2input,
		part1expectedResult: part1expectedResult,
		part2expectedResult: part2expectedResult,
	}
}

func NewSimpleAdventTask(
	part1 AdventTaskPart,
	part2 AdventTaskPart,
) *AdventTask {
	return &AdventTask{
		part1:               part1,
		part2:               part2,
		part1input:          nil,
		part2input:          nil,
		part1expectedResult: nil,
		part2expectedResult: nil,
	}
}

func (t *AdventTask) SetPart1Input(input []string) *AdventTask {
	t.part1input = input
	return t
}

func (t *AdventTask) SetPart2Input(input []string) *AdventTask {
	t.part2input = input
	return t
}

func (t *AdventTask) SetPart1ExpectedResult(expectedResult string) *AdventTask {
	t.part1expectedResult = &expectedResult
	return t
}

func (t *AdventTask) SetPart2ExpectedResult(expectedResult string) *AdventTask {
	t.part2expectedResult = &expectedResult
	return t
}

func (t *AdventTask) RunPart1() error {
	if t.part1input == nil {
		return nil
	}
	return t.runPart(t.part1, t.part1input, t.part1expectedResult)
}

func (t *AdventTask) RunPart2() error {
	if t.part2input == nil {
		return nil
	}
	return t.runPart(t.part2, t.part2input, t.part2expectedResult)
}

func (t *AdventTask) runPart(part AdventTaskPart, input []string, expectedResult *string) error {
	startTime := time.Now()
	result, e := part(input)
	elapsed := time.Since(startTime)
	zap.L().Info(fmt.Sprintf("Part took %s", elapsed))
	if expectedResult != nil {
		if result != *expectedResult {
			errTxt := fmt.Sprintf("Wrong answer, expected: %s, got: %s ", *expectedResult, result)
			zap.L().Error(errTxt)
			return errors.New(errTxt)
		} else {
			zap.L().Info(fmt.Sprintf("OK, expected: %s, got: %s", *expectedResult, result))
		}
	}
	return e
}

func (t *AdventTask) Run() error {
	err := t.RunPart1()
	if err != nil {
		return err
	}
	return t.RunPart2()
}

func (t *AdventTask) RunPart1WithInput(input []string) error {
	return t.SetPart1Input(input).RunPart1()
}

func (t *AdventTask) RunPart2WithInput(input []string) error {
	return t.SetPart2Input(input).RunPart2()
}

func (t *AdventTask) DoesPart1ExpectResult() bool {
	return t.part1expectedResult != nil
}

func (t *AdventTask) DoesPart2ExpectResult() bool {
	return t.part2expectedResult != nil
}
