package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/nickrobinson/aoc2020/pkg/infosec"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func ConvertLineToPasswordData(line string) (infosec.PasswordCountPolicy, string) {
	splitLine := strings.Split(line, " ")
	countInfo := strings.Split(splitLine[0], "-")
	character := splitLine[1]
	character = strings.TrimSuffix(character, ":")
	password := splitLine[2]
	minCount, _ := strconv.Atoi(countInfo[0])
	maxCount, _ := strconv.Atoi(countInfo[1])
	return infosec.PasswordCountPolicy{MinAppearanceCount: minCount, MaxAppearanceCount: maxCount, Character: character}, password
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validPasswordCount := 0
	for scanner.Scan() {
		passwordLine := scanner.Text()
		policy, password := ConvertLineToPasswordData(passwordLine)
		if policy.ValidatePassword(password) {
			validPasswordCount++
		}
	}

	log.Infof("Found %d valid password (count)", validPasswordCount)
}
