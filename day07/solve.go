package day07

import (
	"goaoc/utils"
	"strconv"
)

func SolvePuzzle1(input string) string {
	grid := utils.MakeCharacterMatrixFromText(input)

	// active[x] = true if there's a beam at column x
	active := make([]bool, len(grid[0]))
	for x, c := range grid[0] {
		if c == 'S' {
			active[x] = true
			break
		}
	}

	splits := 0
	for y := 1; y < len(grid); y++ {
		next := make([]bool, len(active))
		for x := range grid[y] {
			if !active[x] {
				continue
			}
			if grid[y][x] == '^' {
				next[x-1] = true
				next[x+1] = true
				splits++
			} else {
				next[x] = true
			}
		}
		active = next
	}

	return strconv.Itoa(splits)
}

func SolvePuzzle2(input string) string {
	grid := utils.MakeCharacterMatrixFromText(input)

	// paths[x] = number of timelines reaching column x
	paths := make([]int, len(grid[0]))
	for x, c := range grid[0] {
		if c == 'S' {
			paths[x] = 1
			break
		}
	}

	for y := 1; y < len(grid); y++ {
		next := make([]int, len(paths))
		for x := range grid[y] {
			if paths[x] == 0 {
				continue
			}
			if grid[y][x] == '^' {
				next[x-1] += paths[x]
				next[x+1] += paths[x]

				// extend
			} else {
				next[x] += paths[x]
			}
		}
		// fmt.Println(next)
		paths = next
	}

	total := 0
	for _, p := range paths {
		total += p
	}
	return strconv.Itoa(total)
}
