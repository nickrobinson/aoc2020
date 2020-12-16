package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type MemoryGame struct {
	StartingNumbers   []int
	SpokenNumbers     map[int][]int
	SpokenWordCounter int
	LastSpokenNumber  int
}

func NewMemoryGame(startingNumbers []int) MemoryGame {
	mg := MemoryGame{StartingNumbers: startingNumbers}
	mg.SpokenNumbers = make(map[int][]int)
	for i, n := range startingNumbers {
		mg.SpokenNumbers[n] = []int{i}
	}
	mg.SpokenWordCounter = len(startingNumbers)
	mg.LastSpokenNumber = startingNumbers[len(startingNumbers)-1]
	return mg
}

func (mg *MemoryGame) Step() int {
	spokenTimes, _ := mg.SpokenNumbers[mg.LastSpokenNumber]
	log.Infof("Number of times %d spoken, %d", mg.LastSpokenNumber, len(spokenTimes))
	if len(spokenTimes) == 1 && spokenTimes[0] == mg.SpokenWordCounter-1 {
		log.Infof("First time %d was spoken, %v, counter: %d", mg.LastSpokenNumber, spokenTimes, mg.SpokenWordCounter)
		mg.LastSpokenNumber = 0
	} else {
		log.Infof("Last number spoken was %d", mg.LastSpokenNumber)
		mg.LastSpokenNumber = spokenTimes[len(spokenTimes)-1] - spokenTimes[len(spokenTimes)-2]
	}
	log.Infof("Number spoken: %d", mg.LastSpokenNumber)
	mg.SpokenNumbers[mg.LastSpokenNumber] = append(mg.SpokenNumbers[mg.LastSpokenNumber], mg.SpokenWordCounter)
	mg.SpokenWordCounter++
	return mg.LastSpokenNumber
}

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	mg := NewMemoryGame([]int{2, 0, 6, 12, 1, 3})
	for i := 0; i < 30000000-len(mg.StartingNumbers); i++ {
		mg.Step()
	}
	log.Warnf("Last number spoken: %d", mg.LastSpokenNumber)
}
