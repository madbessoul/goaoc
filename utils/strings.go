package utils

import (
	"strings"
)

func makeMatrixFromText(text string) [][]string {
	var matrix [][]string
	var lines = strings.Split(text, "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		matrix = append(matrix, fields)
	}
	return matrix
}
