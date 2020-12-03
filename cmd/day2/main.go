package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

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

func ConvertLineToPasswordData(line string) (infosec.PasswordPoilcy, string) {
	re := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)
	matches := re.FindStringSubmatch(line)
	firstCondition, _ := strconv.Atoi(matches[1])
	secondCondition, _ := strconv.Atoi(matches[2])
	return infosec.PasswordPoisitionPoilcy{FirstCheckedPosition: firstCondition, SecondCheckedPosition: secondCondition, Character: matches[3]}, matches[4]
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
