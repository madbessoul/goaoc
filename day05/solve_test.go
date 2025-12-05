package day05_test

import (
	"goaoc/day05"
	"goaoc/utils"
	"testing"
)

func TestPuzzleOnExample1(t *testing.T) {
	input := utils.ReadExample()
	actual := day05.SolvePuzzle1(input)
	expected := "3"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput()
	actual := day05.SolvePuzzle1(input)
	t.Log(actual)
}

func TestPuzzleOnExample2(t *testing.T) {
	input := utils.ReadExample()
	actual := day05.SolvePuzzle2(input)
	expected := "14"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput()
	actual := day05.SolvePuzzle2(input)
	t.Log(actual)
}
