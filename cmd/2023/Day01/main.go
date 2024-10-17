package main

import (
	_ "embed"
	"os"

	"AdventOfCode/internal/factories"
	"go.uber.org/zap"
)

//go:embed input/input-p1.txt
var inputPart1 []byte

//go:embed input/input-p2.txt
var inputPart2 []byte

func init() {
	logger := zap.Must(zap.NewProduction())
	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(zap.NewDevelopment())
	}
	zap.ReplaceGlobals(logger)
}

func main() {
	defer zap.L().Sync()
	zap.L().Info("Task starting...")
	task := factories.CreateAdventTask(inputPart1, inputPart2, part1, part2)

	if e := task.Run(); e != nil {
		zap.L().Error("Task failed!", zap.Error(e))
	}

	zap.L().Info("Task finished.")
}
