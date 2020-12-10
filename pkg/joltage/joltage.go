package joltage

import (
	"sort"

	log "github.com/sirupsen/logrus"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Given an array of joltages find the sum of differences
func FindJoltageDifference(joltages []int) map[int]int {
	differences := make(map[int]int)
	sort.Ints(joltages)
	currentJoltage := 0
	for _, joltage := range joltages {
		log.Infof("joltage: %d", joltage)
		difference := abs(joltage - currentJoltage)
		if difference > 3 {
			log.Errorf("Joltage difference too large for jump from %d to %d", currentJoltage, joltage)
		} else {
			differences[difference]++
			currentJoltage = joltage
		}
	}
	// Account for built-in adapter (always 3 difference)
	differences[3]++
	return differences
}

func areCompatible(low, high int) bool {
	return low+1 == high || low+2 == high || low+3 == high
}

func FindPossibleAdapterSets(fromIndex int, joltages []int, visited map[int]int) int {
	log.Infof("len(joltages): %d, visited: %v, fromIndex: %d", len(joltages), visited, fromIndex)
	if fromIndex >= len(joltages)-3 {
		return 1
	}

	num := joltages[fromIndex]
	if res, ok := visited[num]; ok {
		return res
	}

	var count int
	for i := fromIndex + 1; i < fromIndex+4; i++ {
		n := joltages[i]
		log.Info(num, n)
		if areCompatible(num, n) {
			count += FindPossibleAdapterSets(i, joltages, visited)
		}
	}

	visited[num] = count // store the result
	return count
}
