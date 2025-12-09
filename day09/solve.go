package day09

import (
	"goaoc/utils"
	"strconv"
	"strings"
)

type Loc struct {
	X int
	Y int
}

func SolvePuzzle1(input string) string {
	coords := parseInput(input)
	maxArea := 1

	for _, loc1 := range coords {
		for _, loc2 := range coords[1:] {
			w := utils.Abs(loc2.X-loc1.X) + 1
			h := utils.Abs(loc2.Y-loc1.Y) + 1
			if w*h > maxArea {
				maxArea = w * h
			}
		}
	}
	return strconv.Itoa(maxArea)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isOnSegment(p, a, b Loc) bool {
	if a.X == b.X {
		// Vertical segment
		return p.X == a.X && p.Y >= min(a.Y, b.Y) && p.Y <= max(a.Y, b.Y)
	} else {
		// Horizontal segment
		return p.Y == a.Y && p.X >= min(a.X, b.X) && p.X <= max(a.X, b.X)
	}
}

func isPointInPolygon(p Loc, polygon []Loc) bool {
	inside := false
	n := len(polygon)

	for i := range n {
		j := (i + 1) % n

		// Check if point is on the edge
		if isOnSegment(p, polygon[i], polygon[j]) {
			return true
		}

		// i have no idea what I'm doing
		if ((polygon[i].Y > p.Y) != (polygon[j].Y > p.Y)) &&
			(p.X < (polygon[j].X-polygon[i].X)*(p.Y-polygon[i].Y)/(polygon[j].Y-polygon[i].Y)+polygon[i].X) {
			inside = !inside
		}
	}
	return inside
}

func isRectInPolygon(loc1, loc2 Loc, polygon []Loc) bool {
	minX := min(loc1.X, loc2.X)
	maxX := max(loc1.X, loc2.X)
	minY := min(loc1.Y, loc2.Y)
	maxY := max(loc1.Y, loc2.Y)

	corners := []Loc{
		{X: minX, Y: minY},
		{X: maxX, Y: minY},
		{X: minX, Y: maxY},
		{X: maxX, Y: maxY},
	}

	// all corners must be in polygon
	for _, corner := range corners {
		if !isPointInPolygon(corner, polygon) {
			return false
		}
	}
	return true
}

func SolvePuzzle2(input string) string {
	coords := parseInput(input)
	maxArea := 0
	for i := range len(coords) {
		for j := i + 1; j < len(coords); j++ {
			loc1, loc2 := coords[i], coords[j]

			// cSheck if rectangle is fully contained in polygon
			if isRectInPolygon(loc1, loc2, coords) {
				w := utils.Abs(loc2.X-loc1.X) + 1
				h := utils.Abs(loc2.Y-loc1.Y) + 1
				area := w * h

				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return strconv.Itoa(maxArea)
}

func parseInput(input string) []Loc {
	coords := make([]Loc, 0)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")

		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

		coords = append(coords, Loc{X: x, Y: y})
	}

	return coords
}
