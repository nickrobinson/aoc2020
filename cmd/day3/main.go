package main

import (
	"os"

	"github.com/nickrobinson/aoc2020/pkg/toboggan"
	log "github.com/sirupsen/logrus"
)

type slope struct {
	right int
	down  int
}

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	var slopeTests = []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	flightMap, _ := toboggan.LoadMap("input.txt")
	for _, s := range slopeTests {
		log.WithFields(log.Fields{"right": s.right, "down": s.down}).Infof("Tree Count: %d", toboggan.GetTrajectoryTreeCount(s.right, s.down, flightMap))
	}
}
