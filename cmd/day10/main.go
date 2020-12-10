package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/nickrobinson/aoc2020/pkg/joltage"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	fp, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(fp)
	joltages := []int{0}

	for scanner.Scan() {
		lineText := scanner.Text()
		joltageVal, _ := strconv.Atoi(lineText)
		joltages = append(joltages, joltageVal)
		differences := joltage.FindJoltageDifference(joltages)
		log.Infof("Differences: %v", differences)
	}

	visited := make(map[int]int)
	sets := joltage.FindPossibleAdapterSets(0, joltages, visited)
	log.Infof("Possible adapter combinations: %d", sets)
}
