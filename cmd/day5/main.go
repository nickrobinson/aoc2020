package main

import (
	"bufio"
	"os"
	"sort"

	"github.com/nickrobinson/aoc2020/pkg/boarding"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	fp, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(fp)

	maxSeatNum := 0
	occupiedSeats := []int{}
	for scanner.Scan() {
		lineText := scanner.Text()
		bp := boarding.BoardingPass{SpacePartition: lineText}
		seatNum := bp.GetSeatNumber()
		occupiedSeats = append(occupiedSeats, seatNum)
		if seatNum > maxSeatNum {
			maxSeatNum = seatNum
		}
	}

	sort.Ints(occupiedSeats)

	for i, s := range occupiedSeats[:len(occupiedSeats)-1] {
		if occupiedSeats[i+1] != s+1 {
			log.Infof("Empty seat found at %d", s+1)
		}
	}

	log.Infof("Max Seat Num: %d", maxSeatNum)
}
