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
	var tests = []struct {
		name     string
		passport Passport
		valid    bool
	}{
		{"Basic valid passport", Passport{EyeColor: "gry", PassportId: "860033327", ExpirationYear: 2020, HairColor: "#fffffd", BirthYear: 1937, IssueYear: 2017, CountryId: "147", Height: "183cm"}, true},
		{"Passport without CID", Passport{EyeColor: "gry", PassportId: "860033327", ExpirationYear: 2020, HairColor: "#fffffd", BirthYear: 1937, IssueYear: 2017, Height: "183cm"}, true},
		{"Passport without PID", Passport{EyeColor: "gry", ExpirationYear: 2020, HairColor: "#fffffd", BirthYear: 1937, IssueYear: 2017, Height: "183cm"}, false},
		{"Passport with missing fields", Passport{EyeColor: "brn", ExpirationYear: 2025, HairColor: "#cfa07d", IssueYear: 2011, Height: "59in", PassportId: "166559648"}, false},
		{"Passport with invalid height", Passport{EyeColor: "gry", PassportId: "860033327", ExpirationYear: 2020, HairColor: "#fffffd", BirthYear: 1937, IssueYear: 2017, CountryId: "147", Height: "512cm"}, false},
		{"Passport with invalid hair color", Passport{EyeColor: "gry", PassportId: "860033327", ExpirationYear: 2020, HairColor: "#fffff", BirthYear: 1937, IssueYear: 2017, CountryId: "147", Height: "183cm"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := IsValidPassport(&tt.passport)
			if valid != tt.valid {
				t.Errorf("IsValidPassport(%s) got %v, want %v", tt.name, valid, tt.valid)
			}
		})
	}
}
