package day02_test

import (
	"goaoc/day02"
	"goaoc/utils"
	"testing"
)

func TestPuzzleOnExample1(t *testing.T) {
	input := utils.ReadExample()
	actual := day02.SolvePuzzle1(input)
	expected := "1227775554"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput()
	actual := day02.SolvePuzzle1(input)
	t.Log(actual)
}

func TestPuzzleOnExample2(t *testing.T) {
	input := utils.ReadExample()
	actual := day02.SolvePuzzle2(input)
	expected := "4174379265"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput()
	actual := day02.SolvePuzzle2(input)
	t.Log(actual)
}
