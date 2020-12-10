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

func FindPossibleAdapterSets(joltages []int, currentPath []int) int {
	log.Infof("len(joltages): %d, currentPath: %v", len(joltages), currentPath)
	sort.Ints(joltages)

	if len(joltages) == 1 {
		if joltages[0] <= 3 {
			return 1
		} else {
			return 0
		}
	}

	// if len(joltages) == 2 {
	// 	if joltages[1]-joltages[0] <= 3 && joltages[0] <= 3 {
	// 		log.Infof("Found arrangement: %v", currentPath)
	// 		return 1
	// 	} else {
	// 		return 0
	// 	}
	// }

	sum := 0
	for i := len(joltages) - 2; i > len(joltages)-5; i-- {
		if i > 0 {
			log.Infof("Checking that joltages[%d] (%d) - joltages[%d] (%d) <= 3", len(joltages)-1, joltages[len(joltages)-1], i, joltages[i])
			if (joltages[len(joltages)-1] - joltages[i]) <= 3 {
				log.Infof("Calling FindPossibleAdapterSets with joltages of len: %d", len(joltages[:i+1]))
				sum += FindPossibleAdapterSets(joltages[:i], append(currentPath, joltages[len(joltages)-1]))
			}
		}
	}
	return sum
}
