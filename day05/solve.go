package day05

import (
	"sort"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) string {
	ranges, numbers := parseInput(input)

	count := 0
	for _, n := range numbers {
		for _, r := range ranges {
			if n >= r[0] && n <= r[1] {
				// fmt.Printf("number %d is in range %d-%d\n", n, r[0], r[1])
				count++
				break
			}
		}
	}
	return strconv.Itoa(count)
}

func SolvePuzzle2(input string) string {
	ranges, _ := parseInput(input)

	// sort by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	// fmt.Println("sorted ranges:", ranges)
	var mergedRanges [][2]int64
	for i := range len(ranges) {
		// fist range or no overlap, add new range
		if i == 0 || ranges[i][0] > mergedRanges[len(mergedRanges)-1][1]+1 {
			mergedRanges = append(mergedRanges, ranges[i])
		} else {
			// extend the last range
			mergedRanges[len(mergedRanges)-1][1] = max(mergedRanges[len(mergedRanges)-1][1], ranges[i][1])
		}
	}
	// fmt.Println("merge ranges:", mergedRanges)

	freshIdCount := 0
	for _, r := range mergedRanges {
		freshIdCount += int(r[1] - r[0] + 1)
	}
	return strconv.Itoa(freshIdCount)
}

func parseInput(input string) ([][2]int64, []int64) {
	// ranges in the form [start-end] until blank line is found
	lines := strings.Split(input, "\n")
	var ranges [][2]int64
	for _, line := range lines {
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, [2]int64{int64(start), int64(end)})
	}

	// numbers after blank line, loop in reverse, skip EOF emtpy line
	numbers := []int64{}
	for i := len(lines) - 2; i >= 0; i-- {
		if lines[i] == "" {
			break
		}
		n, _ := strconv.Atoi(lines[i])
		numbers = append(numbers, int64(n))
	}
	return ranges, numbers
}
