package day04_test

import (
	"goaoc/day04"
	"goaoc/utils"
	"testing"
)

func TestPuzzleOnExample1(t *testing.T) {
	input := utils.ReadExample()
	actual := day04.SolvePuzzle1(input)
	expected := "13"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput()
	actual := day04.SolvePuzzle1(input)
	t.Log(actual)
}

func TestPuzzleOnExample2(t *testing.T) {
	input := utils.ReadExample()
	actual := day04.SolvePuzzle2(input)
	expected := "43"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput()
	actual := day04.SolvePuzzle2(input)
	t.Log(actual)
}
