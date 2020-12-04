package passport

import (
	"testing"
)

func TestLoadingMap(t *testing.T) {
	passports, err := LoadPassports("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if len(passports) != 4 {
		t.Errorf("Expected 4 passports, found %d", len(passports))
	}

	if passports[0].PassportId != "860033327" {
		t.Errorf("Expected passport id of 860033327, Got %s", passports[0].PassportId)
	}
}

func TestIsValidPassport(t *testing.T) {
	passport := Passport{EyeColor: "gry", PassportId: "860033327", ExpirationYear: 2020, HairColor: "#fffffd", BirthYear: 1937, IssueYear: 2017, CountryId: "147", Height: "183cm"}
	if !IsValidPassport(&passport) {
		t.Errorf("Expected passport to be valid, was invalid")
	}

	passportWithoutCid := Passport{EyeColor: "gry", PassportId: "860033327", ExpirationYear: 2020, HairColor: "#fffffd", BirthYear: 1937, IssueYear: 2017, Height: "183cm"}
	if !IsValidPassport(&passportWithoutCid) {
		t.Errorf("Expected passport to be valid, was invalid")
	}

	passportWithoutPid := Passport{EyeColor: "gry", ExpirationYear: 2020, HairColor: "#fffffd", BirthYear: 1937, IssueYear: 2017, Height: "183cm"}
	if IsValidPassport(&passportWithoutPid) {
		t.Errorf("Expected passport to be invalid, was valid")
	}

	passportWithMissingFields := Passport{EyeColor: "brn", ExpirationYear: 2025, HairColor: "#cfa07d", IssueYear: 2011, Height: "59in", PassportId: "166559648"}
	if IsValidPassport(&passportWithMissingFields) {
		t.Errorf("Expected passport to be invalid, was valid")
	}

	passportWithInvalidHeight := Passport{EyeColor: "gry", PassportId: "860033327", ExpirationYear: 2020, HairColor: "#fffffd", BirthYear: 1937, IssueYear: 2017, CountryId: "147", Height: "512cm"}
	if IsValidPassport(&passportWithInvalidHeight) {
		t.Errorf("Expected passport to be invalid, was valid")
	}
}
