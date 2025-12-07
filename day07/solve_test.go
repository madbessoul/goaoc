package day07_test

import (
	"goaoc/day07"
	"goaoc/utils"
	"testing"
)

func TestPuzzleOnExample1(t *testing.T) {
	input := utils.ReadExample()
	actual := day07.SolvePuzzle1(input)
	expected := "21"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput()
	actual := day07.SolvePuzzle1(input)
	t.Log(actual)
}

func TestPuzzleOnExample2(t *testing.T) {
	input := utils.ReadExample()
	actual := day07.SolvePuzzle2(input)
	expected := "40"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput()
	actual := day07.SolvePuzzle2(input)
	t.Log(actual)
}
