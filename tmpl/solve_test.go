package dayXX_test

import (
	"goaoc/dayXX"
	"goaoc/utils"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput()
	actual := dayXX.SolvePuzzle1(input)
	t.Log(actual)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput()
	actual := dayXX.SolvePuzzle2(input)
	t.Log(actual)
}
