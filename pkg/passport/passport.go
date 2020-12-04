package passport

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	validator "github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var validate = validator.New()

const MinHeightCm = 150
const MaxHeightCm = 193
const MinHeightIn = 59
const MaxHeightIn = 76

func init() {
	// Register custom height validation
	_ = validate.RegisterValidation("height", func(fl validator.FieldLevel) bool {
		fieldVal := fl.Field().String()
		height, _ := strconv.Atoi(string(fieldVal[:len(fieldVal)-2]))
		if strings.HasSuffix(fieldVal, "in") {
			if height >= MinHeightIn && height <= MaxHeightIn {
				return true
			}
		} else if strings.HasSuffix(fieldVal, "cm") {
			if height >= MinHeightCm && height <= MaxHeightCm {
				return true
			}
		}
		return false
	})
}

type Passport struct {
	BirthYear      int    `validate:"required,gte=1920,lte=2002"`
	IssueYear      int    `validate:"required,gte=2010,lte=2020"`
	ExpirationYear int    `validate:"required,gte=2020,lte=2030"`
	Height         string `validate:"required,endswith=cm|endswith=in,height"`
	HairColor      string `validate:"required,hexcolor"`
	EyeColor       string `validate:"required,eq=amb|eq=blu|eq=brn|eq=gry|eq=grn|eq=hzl|eq=oth"`
	PassportId     string `validate:"required,number,len=9"`
	CountryId      string
}

func LoadPassports(filepath string) ([]Passport, error) {
	fp, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Got error while trying to open %s: %v", filepath, err)
		return nil, err
	}

	scanner := bufio.NewScanner(fp)
	passports := []Passport{}
	passport := Passport{}

	for scanner.Scan() {
		lineText := scanner.Text()

		if lineText == "" {
			passports = append(passports, passport)
			passport = Passport{}
			continue
		}

		passportKv := strings.Split(lineText, " ")
		for _, p := range passportKv {
			passportData := strings.Split(p, ":")
			passportVal := passportData[1]
			switch passportKey := passportData[0]; passportKey {
			case "byr":
				passport.BirthYear, _ = strconv.Atoi(passportVal)
			case "iyr":
				passport.IssueYear, _ = strconv.Atoi(passportVal)
			case "eyr":
				passport.ExpirationYear, _ = strconv.Atoi(passportVal)
			case "hgt":
				passport.Height = passportVal
			case "hcl":
				passport.HairColor = passportVal
			case "ecl":
				passport.EyeColor = passportVal
			case "pid":
				passport.PassportId = passportVal
			case "cid":
				passport.CountryId = passportVal
			default:
				log.Errorf("Unknown key %s", passportKey)
			}
		}
	}

	// Make sure we cover last passport in the file
	passports = append(passports, passport)

	return passports, nil
}

// IsValidPassport uses validations provided in Passport struct
// to validate that a passport has required parameters in expected format
func IsValidPassport(passport *Passport) bool {
	err := validate.Struct(passport)
	if err != nil {
		log.Warn(err)
		return false
	}
	return true
}
