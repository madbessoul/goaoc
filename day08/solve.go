package day08

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Light struct {
	x, y, z int
}

type Pair struct {
	light1, light2 Light
}

func NewPair(light1, light2 Light) Pair {
	return Pair{light1, light2}
}

func ToCircuit(pair Pair) Circuit {
	return Circuit{lights: []Light{pair.light1, pair.light2}}
}

type Circuit struct {
	lights []Light
}

func PairDistance(pair Pair) int {
	return int(math.Sqrt(math.Pow(float64(pair.light1.x-pair.light2.x), 2) +
		math.Pow(float64(pair.light1.y-pair.light2.y), 2) +
		math.Pow(float64(pair.light1.z-pair.light2.z), 2)))
}

func SolvePuzzle1(input string) string {
	lights := parseInput(input)

	// compute all pair combinations, what could go wrong
	pairs := make([]Pair, 0, len(lights)*(len(lights)-1)/2)
	for i := range len(lights) {
		for j := i + 1; j < len(lights); j++ {
			pairs = append(pairs, NewPair(lights[i], lights[j]))
		}
	}

	// sort pairs by distance
	sort.Slice(pairs, func(i, j int) bool {
		return PairDistance(pairs[i]) < PairDistance(pairs[j])
	})
	fmt.Println("Number of pairs:", len(pairs))

	// each light starts in its own circuit
	circuits := make([]Circuit, len(lights))
	for i, light := range lights {
		circuits[i] = Circuit{lights: []Light{light}}
	}

	// process only the first 1000 pairs
	connectionsToMake := 1000

	// adapt to the example
	if len(lights) <= 20 {
		connectionsToMake = 10
	}

	if len(pairs) < connectionsToMake {
		connectionsToMake = len(pairs)
	}

	for _, pair := range pairs[:connectionsToMake] {
		// find which circuits contain light1 and light2
		circuit1Idx := -1
		circuit2Idx := -1
		for i, circuit := range circuits {
			for _, light := range circuit.lights {
				if light == pair.light1 {
					circuit1Idx = i
				}
				if light == pair.light2 {
					circuit2Idx = i
				}
			}
		}

		// if both lights are already in the same circuit, skip
		if circuit1Idx == circuit2Idx {
			continue
		}

		// merge the two circuits: add all lights from circuit2 to circuit1
		circuits[circuit1Idx].lights = append(circuits[circuit1Idx].lights, circuits[circuit2Idx].lights...)
		// remove circuit2
		circuits = append(circuits[:circuit2Idx], circuits[circuit2Idx+1:]...)
	}

	fmt.Printf("Number of circuits: %d\n", len(circuits))

	// sort circuits by size desc
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i].lights) > len(circuits[j].lights)
	})

	// multiply sizes of three largest circuits
	result := len(circuits[0].lights) * len(circuits[1].lights) * len(circuits[2].lights)

	return strconv.Itoa(result)
}

func SolvePuzzle2(input string) string {
	lights := parseInput(input)

	// compute all pair combinations, what could go wrong
	pairs := make([]Pair, 0, len(lights)*(len(lights)-1)/2)
	for i := range len(lights) {
		for j := i + 1; j < len(lights); j++ {
			pairs = append(pairs, NewPair(lights[i], lights[j]))
		}
	}

	// sort pairs by distance
	sort.Slice(pairs, func(i, j int) bool {
		return PairDistance(pairs[i]) < PairDistance(pairs[j])
	})
	fmt.Println("Number of pairs:", len(pairs))

	// each light starts in its own circuit
	circuits := make([]Circuit, len(lights))
	for i, light := range lights {
		circuits[i] = Circuit{lights: []Light{light}}
	}

	var result int
	for _, pair := range pairs {
		// find which circuits contain light1 and light2
		circuit1Idx := -1
		circuit2Idx := -1
		for i, circuit := range circuits {
			for _, light := range circuit.lights {
				if light == pair.light1 {
					circuit1Idx = i
				}
				if light == pair.light2 {
					circuit2Idx = i
				}
			}
		}

		// if both lights are already in the same circuit, skip
		if circuit1Idx == circuit2Idx {
			continue
		}

		// merge the two circuits: add all lights from circuit2 to circuit1
		circuits[circuit1Idx].lights = append(circuits[circuit1Idx].lights, circuits[circuit2Idx].lights...)
		// remove circuit2
		circuits = append(circuits[:circuit2Idx], circuits[circuit2Idx+1:]...)

		if len(circuits) == 1 {
			result = pair.light1.x * pair.light2.x
			break
		}
	}

	fmt.Printf("Number of circuits: %d\n", len(circuits))

	// sort circuits by size desc
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i].lights) > len(circuits[j].lights)
	})

	return strconv.Itoa(result)
}

func parseInput(input string) []Light {
	var lights []Light
	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		light := Light{x, y, z}
		lights = append(lights, light)
	}
	return lights
}
