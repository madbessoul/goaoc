package day04

import (
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) string {
	grid := StringToGrid(input)
	width := len(grid[0])
	length := len(grid)

	nbAccessiblePaperRolls := 0
	for x := range width {
		for y := range length {
			if grid[y][x] == '@' {
				if CountSurroundingPaperRolls(x, y, grid) < 4 {
					nbAccessiblePaperRolls++
				}
			}
		}
	}
	return strconv.Itoa(nbAccessiblePaperRolls)
}

func SolvePuzzle2(input string) string {
	grid := StringToGrid(input)
	width := len(grid[0])
	length := len(grid)

	nbAccessiblePaperRolls := 0
	nbRounds := 0
	for {
		currentAccessiblePaperRolls := 0
		paperRollCoords := make([][2]int, 0)
		for x := range width {
			for y := range length {
				if grid[y][x] == '@' {
					if CountSurroundingPaperRolls(x, y, grid) < 4 {
						currentAccessiblePaperRolls++
						paperRollCoords = append(paperRollCoords, [2]int{x, y})
					}
				}
			}
		}
		nbAccessiblePaperRolls += currentAccessiblePaperRolls
		if currentAccessiblePaperRolls == 0 {
			break
		}
		RemovePaperRolls(&grid, paperRollCoords)
		nbRounds++
	}

	return strconv.Itoa(nbAccessiblePaperRolls)
}

func RemovePaperRolls(grid *[][]rune, paperRollCoords [][2]int) {
	for _, coord := range paperRollCoords {
		(*grid)[coord[1]][coord[0]] = '.'
	}
}

func CountSurroundingPaperRolls(x, y int, grid [][]rune) int {
	count := 0
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && grid[i][j] == '@' && !(i == y && j == x) {
				count++
			}
		}
	}
	return count
}

func StringToGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, 0, len(lines))
	for _, line := range lines {
		if line != "" {
			grid = append(grid, []rune(line))
		}
	}
	return grid
}
