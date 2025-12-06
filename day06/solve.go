package day06

import (
	"fmt"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) string {
	grid, lastLine := parseInput(input)
	total := 0
	for i := range len(lastLine) {
		op := lastLine[i]

		switch op {
		case "+":
			colRes := 0
			for j := range grid {
				val, _ := strconv.Atoi(grid[j][i])
				colRes += val
			}
			total += colRes
		case "*":
			colRes := 1
			for j := range grid {
				val, _ := strconv.Atoi(grid[j][i])
				colRes *= val
			}
			total += colRes
		}
	}

	return strconv.Itoa(total)
}

func SolvePuzzle2(input string) string {

	grid := parseToCharacterGrid(input)
	fmt.Println(grid)
	totalLines := len(grid)
	globalNumberList := make([][]int, 0)
	newNumberList := make([]int, 0)
	for i := range grid[0] {

		oneCharCol := make([]string, 0)
		for j := range totalLines {
			oneCharCol = append(oneCharCol, string(grid[j][i]))
		}

		if strings.Join(oneCharCol, "") != strings.Repeat(" ", totalLines) {
			for k := range oneCharCol {
				if oneCharCol[k] == " " {
					oneCharCol[k] = ""
				}
			}

			n, _ := strconv.Atoi(strings.Join(oneCharCol, ""))
			fmt.Println(n)
			newNumberList = append(newNumberList, n)

		} else {
			// reset new number list
			globalNumberList = append(globalNumberList, newNumberList)
			newNumberList = make([]int, 0)
		}
	}
	// last col
	if len(newNumberList) > 0 {
		globalNumberList = append(globalNumberList, newNumberList)
	}

	_, lastLine := parseInput(input)
	result := 0
	for i := range len(lastLine) {
		op := lastLine[i]

		switch op {
		case "+":
			colRes := 0
			for j := range globalNumberList[i] {
				colRes += globalNumberList[i][j]
			}
			result += colRes
		case "*":
			colRes := 1
			for j := range globalNumberList[i] {
				colRes *= globalNumberList[i][j]
			}
			result += colRes
		}
	}
	return strconv.Itoa(result)
}

func parseToCharacterGrid(input string) [][]string {
	lines := make([]string, 0)
	maxLen := 0

	for line := range strings.SplitSeq(input, "\n") {
		// skip last line (operators) and empty lines
		if line == "" || strings.Contains(line, "+") || strings.Contains(line, "*") {
			continue
		}
		lines = append(lines, line)
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// build grid and normalize
	grid := make([][]string, 0)
	for _, line := range lines {
		row := make([]string, 0)
		for _, char := range []rune(line) {
			row = append(row, string(char))
		}
		// pad
		for len(row) < maxLen {
			row = append(row, " ")
		}
		grid = append(grid, row)
	}
	return grid
}

func parseInput(input string) ([][]string, []string) {
	grid := make([][]string, 0)
	for line := range strings.SplitSeq(input, "\n") {
		// skip last line
		if line == "" || strings.Contains(line, "+") || strings.Contains(line, "*") {
			continue
		}
		row := make([]string, 0)
		for num := range strings.FieldsSeq(line) {
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	// get last line of input, seq of operations,
	lastLine := strings.Fields(strings.Split(input, "\n")[len(strings.Split(input, "\n"))-2])

	return grid, lastLine
}
