package day01_test

import (
	"goaoc/day01"
	"goaoc/utils"
	"testing"
)

func TestPuzzleOnExample1(t *testing.T) {
	input := utils.ReadExample()
	actual := day01.SolvePuzzle1(input)
	expected := "3"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput()
	actual := day01.SolvePuzzle1(input)
	t.Log(actual)
}

func TestPuzzleOnExample2(t *testing.T) {
	input := utils.ReadExample()
	actual := day01.SolvePuzzle2(input)
	expected := "6"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput()
	actual := day01.SolvePuzzle2(input)
	t.Log(actual)
}
