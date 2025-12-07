package utils

import (
	"strings"
)

func MakeMatrixFromText(text string) [][]string {
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

func MakeCharacterMatrixFromText(text string) [][]rune {
	var matrix [][]rune
	var lines = strings.Split(text, "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		matrix = append(matrix, []rune(fields[0]))
	}
	return matrix
}
