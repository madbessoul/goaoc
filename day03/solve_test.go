package day03_test

import (
	"goaoc/day03"
	"goaoc/utils"
	"testing"
)

func TestPuzzleOnExample1(t *testing.T) {
	input := utils.ReadExample()
	actual := day03.SolvePuzzle1(input)
	expected := "357"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput()
	actual := day03.SolvePuzzle1(input)
	t.Log(actual)
}

func TestPuzzleOnExample2(t *testing.T) {
	input := utils.ReadExample()
	actual := day03.SolvePuzzle2(input)
	expected := "3121910778619"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput()
	actual := day03.SolvePuzzle2(input)
	t.Log(actual)
}
