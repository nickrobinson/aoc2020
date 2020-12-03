package toboggan

import (
	"fmt"
	"testing"
)

func TestLoadingMap(t *testing.T) {
	flightMap, err := LoadMap("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v", flightMap[0])

	if len(flightMap) != 11 {
		t.Errorf("Expected length of 11, Got %d", len(flightMap))
	}

	if len(flightMap[1]) != 11 {
		t.Errorf("Expected length of 11, Got %d", len(flightMap[0]))
	}
}

func TestTreeCounting(t *testing.T) {
	flightMap, _ := LoadMap("input.txt")
	treeCount := GetTrajectoryTreeCount(3, 1, flightMap)
	if treeCount != 7 {
		t.Errorf("Expected to see 7 trees, saw %d", treeCount)
	}
}
