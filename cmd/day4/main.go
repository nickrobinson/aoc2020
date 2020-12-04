package main

import (
	"os"

	"github.com/nickrobinson/aoc2020/pkg/passport"
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
	passports, _ := passport.LoadPassports("input.txt")
	validPassports := 0
	for _, p := range passports {
		passportValid := passport.IsValidPassport(&p)
		log.WithField("isValid", passportValid).Infof("Passport: %+v", p)
		if passportValid {
			validPassports++
		}
	}
	log.Infof("Valid Passport Count: %d", validPassports)
}
