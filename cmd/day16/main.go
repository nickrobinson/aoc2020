package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type IdRange struct {
	LowNum  int
	HighNum int
}

type TicketValidator struct {
	ValidRanges []IdRange
}

type Ticket struct {
	Numbers []int
}

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func LoadValidRanges(filename string) TicketValidator {
	validator := TicketValidator{}
	re := regexp.MustCompile(`.*: (\d+)-(\d+) or (\d+)-(\d+)`)
	fp, _ := os.Open(filename)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lineText := scanner.Text()
		matches := re.FindStringSubmatch(lineText)
		for i := 1; i < len(matches); i += 2 {
			lowVal, _ := strconv.Atoi(matches[i])
			highVal, _ := strconv.Atoi(matches[i+1])
			validator.ValidRanges = append(validator.ValidRanges, IdRange{LowNum: lowVal, HighNum: highVal})
		}
	}
	return validator
}

func LoadTickets(filename string) []Ticket {
	fp, _ := os.Open(filename)
	tickets := []Ticket{}

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		ticket := Ticket{}
		lineText := scanner.Text()
		values := strings.Split(lineText, ",")
		for _, v := range values {
			numVal, _ := strconv.Atoi(v)
			ticket.Numbers = append(ticket.Numbers, numVal)
		}
		tickets = append(tickets, ticket)
	}
	return tickets
}

func main() {
	tickets := LoadTickets("tickets.txt")
	validator := LoadValidRanges("ranges.txt")
	invalidValues := []int{}
	for _, ticket := range tickets {
		log.Infof("Validating ticket %v", ticket)
		for _, num := range ticket.Numbers {
			valid := false
			for _, idRange := range validator.ValidRanges {
				if idRange.LowNum <= num && idRange.HighNum >= num {
					valid = true
				}
			}
			if !valid {
				log.Infof("Number %d in ticket %v is not valid", num, ticket)
				invalidValues = append(invalidValues, num)
			}
		}
	}
	log.Infof("Invlaid Values: %v", invalidValues)
	errorRate := 0
	for _, v := range invalidValues {
		errorRate += v
	}
	log.Infof("Error rate: %d", errorRate)
}
