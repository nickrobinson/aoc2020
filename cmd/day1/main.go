package main

import (
	"bufio"
	"io"
	"os"
	"strconv"

	"github.com/nickrobinson/aoc2020/pkg/accounting"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {
	fp, _ := os.Open("input.txt")
	ints, _ := ReadInts(fp)
	match, err := accounting.FindThreeMatchingExpenses(ints, 2020)
	if err != nil {
		log.Errorf("Got error: %v", err)
	}
	log.WithFields(log.Fields{"matches": match}).Info("Found matches!")
}
