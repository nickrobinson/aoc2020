package toboggan

import (
	"bufio"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

const TreeRune = "#"

func LoadMap(filepath string) ([][]string, error) {
	var flightMap [][]string
	fp, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Got error while trying to open %s: %v", filepath, err)
		return nil, err
	}

	scanner := bufio.NewScanner(fp)

	i := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		points := strings.Split(lineText, "")
		flightMap = append(flightMap, points)
		i++
	}

	return flightMap, nil
}

func GetTrajectoryTreeCount(stepsRight int, stepsDown int, flightMap [][]string) int {
	treeCount := 0
	log.Debugf("Getting tree count for map with len %d, %d and trajectory down: %d right: %d", len(flightMap), len(flightMap[0]), stepsDown, stepsRight)
	i := stepsDown
	for j := stepsRight; ; j += stepsRight {
		log.WithFields(logrus.Fields{"i": i, "j": j}).Debug("Checking Position")
		if flightMap[i][j%len(flightMap[0])] == TreeRune {
			treeCount++
		}
		i += stepsDown
		if i >= len(flightMap) {
			break
		}
	}

	return treeCount
}
