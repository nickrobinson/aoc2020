package main

import (
	"bufio"
	"os"

	"github.com/nickrobinson/aoc2020/pkg/customs"
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

	groupAnswers := ""
	groupMemberCount := 0
	totalAnswerCount := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		if lineText == "" {
			groupAnswerCount := customs.GetCustomsCount(groupAnswers, groupMemberCount)
			totalAnswerCount += groupAnswerCount
			log.Infof("Group Answer Count: %d, Current answer count: %d", groupAnswerCount, totalAnswerCount)
			groupAnswers = ""
			groupMemberCount = 0
			continue
		}
		groupAnswers += lineText
		groupMemberCount++
	}
}
