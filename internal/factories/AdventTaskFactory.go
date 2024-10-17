package factories

import (
	"AdventOfCode/internal/services"
	typeconvert "AdventOfCode/pkg/TypeConvert"
)

func CreateAdventTask(
	inputPart1 []byte,
	inputPart2 []byte,
	part1 services.AdventTaskPart,
	part2 services.AdventTaskPart,
) *services.AdventTask {
	input1 := typeconvert.ArrayOfBytesToArrayOfStrings(inputPart1)
	input2 := typeconvert.ArrayOfBytesToArrayOfStrings(inputPart2)

	task := services.NewSimpleAdventTask(part1, part2)
	task.SetPart1Input(input1)
	task.SetPart2Input(input2)

	return task
}
