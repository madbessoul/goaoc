package day03

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	joltage := 0
	for scanner.Scan() {
		line := scanner.Text()
		bank, _ := stringToDigits(line)
		largestBank := 0
		// Brute force goes brrr
		for i := range len(bank) {
			for j := i + 1; j < len(bank); j++ {
				capacity := int(bank[i])*10 + int(bank[j])
				if largestBank < capacity {
					largestBank = capacity
				}
			}
		}
		// fmt.Println(largestBank)
		joltage += largestBank
	}
	return strconv.Itoa(joltage)
}

func SolvePuzzle2(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	joltage := 0
	for scanner.Scan() {
		bank := scanner.Text()
		joltage += int(maxJoltage(bank, 12))
	}
	return strconv.Itoa(joltage)
}

func maxJoltage(bank string, k int) int64 {
	n := len(bank)
	result := int64(0)
	resultStr := ""
	start := 0

	for i := range k {
		// If n = 15, search for leftmost digit but keep space for the rest
		// search max only in [0..3]

		end := n - (k - i)

		bestPos := start
		bestDigit := bank[start]

		for pos := start + 1; pos <= end; pos++ {
			if bank[pos] > bestDigit {
				bestDigit = bank[pos]
				bestPos = pos
			}
		}

		// Black magic shit with ascii, thanks Claude
		result = result*10 + int64(bestDigit-'0')
		resultStr += string(bestDigit)
		start = bestPos + 1
	}

	return result
}

func stringToDigits(s string) ([]int, error) {
	digits := make([]int, len(s))
	for i, r := range s {
		if r < '0' || r > '9' {
			return nil, fmt.Errorf("invalid digit at position %d: %c", i, r)
		}
		digits[i] = int(r - '0')
	}
	return digits, nil
}
