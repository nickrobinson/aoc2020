package encryption

import (
	"errors"
	"sort"

	log "github.com/sirupsen/logrus"
)

func FindInvalidNumber(preambleLen int, values []int) (int, error) {
	for currentPosition := preambleLen; currentPosition < len(values)-preambleLen; currentPosition++ {
		log.Infof("Candidates: [%d-%d]", currentPosition-preambleLen, currentPosition)
		var candidates = make([]int, preambleLen)
		copy(candidates, values[currentPosition-preambleLen:currentPosition])
		sort.Ints(candidates)
		log.Infof("Checking %d against candidates: %v", values[currentPosition], candidates)
		i := 0
		j := len(candidates) - 1
		for i != j {
			sum := candidates[i] + candidates[j]
			if sum < values[currentPosition] {
				i++
			} else if sum > values[currentPosition] {
				j--
			} else {
				log.Infof("Found a match for %d with [%d, %d]", values[currentPosition], candidates[i], candidates[j])
				break
			}
		}
		if i == j {
			return values[currentPosition], nil
		}
	}

	return 0, errors.New("Could not find invalid number")
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func FindContiguousMatch(target int, values []int) (int, error) {
	for i := 0; i < len(values)-1; i++ {
		j := i + 2
		for ; j < len(values); j++ {
			log.Infof("Checking values[%d:%d]", i, j)
			sum := sum(values[i:j])
			log.Infof("sum: %d", sum)
			if sum == target {
				log.Infof("Values %v", values[i:j])
				return values[i] + values[j], nil
			} else if sum > target {
				break
			}
		}
	}
	return 0, nil
}
