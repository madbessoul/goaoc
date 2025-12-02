package day02

import (
	"fmt"
	"strconv"
	"strings"

	// no backref support in go
	"github.com/gijsbers/go-pcre"
)

func SolvePuzzle1(input string) string {
	input = strings.TrimSpace(input)
	ranges := strings.Split(input, ",")
	var sum int
	regex := pcre.MustCompile(`^(\d+)\1$`, 0)

	for _, interval := range ranges {

		// Goddamn trailing whitespaces
		interval = strings.TrimSpace(interval)
		// fmt.Println(interval)
		parts := strings.Split(interval, "-")
		low, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		high, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		fmt.Printf("Processing range: %d to %d\n", low, high)
		for i := low; i <= high; i++ {
			s := strconv.Itoa(i)
			matcher := regex.MatcherString(s, 0)
			match := matcher.Matches()
			if match {
				sum += i
			}
		}
	}
	return strconv.Itoa(sum)
}

func SolvePuzzle2(input string) string {
	input = strings.TrimSpace(input)
	ranges := strings.Split(input, ",")
	var sum int
	regex := pcre.MustCompile(`^(\d+)\1+$`, 0)

	for _, interval := range ranges {

		// Goddamn trailing whitespaces
		interval = strings.TrimSpace(interval)
		// fmt.Println(interval)
		parts := strings.Split(interval, "-")
		low, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		high, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		fmt.Printf("Processing range: %d to %d\n", low, high)
		for i := low; i <= high; i++ {
			s := strconv.Itoa(i)
			matcher := regex.MatcherString(s, 0)
			match := matcher.Matches()
			if match {
				sum += i
			}
		}
	}
	return strconv.Itoa(sum)
}
