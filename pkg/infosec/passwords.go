package infosec

import (
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type PasswordPoilcy interface {
	ValidatePassword(string) bool
}

type PasswordCountPolicy struct {
	MinAppearanceCount int
	MaxAppearanceCount int
	Character          string
}

type PasswordPoisitionPoilcy struct {
	FirstCheckedPosition  int
	SecondCheckedPosition int
	Character             string
}

func (p PasswordCountPolicy) ValidatePassword(password string) bool {
	appearanceCount := strings.Count(password, p.Character)
	if appearanceCount >= p.MinAppearanceCount && appearanceCount <= p.MaxAppearanceCount {
		return true
	}
	return false
}

func (p PasswordPoisitionPoilcy) ValidatePassword(password string) bool {
	foundMatch := false
	for i := 0; i < p.SecondCheckedPosition; i++ {
		if string(password[i]) == p.Character {
			// offset by 1 since password policies are not zero-indexed
			if i+1 == p.FirstCheckedPosition || i+1 == p.SecondCheckedPosition {
				if foundMatch {
					return false
				}
				foundMatch = true
			}
		}
	}
	if foundMatch {
		log.Debugf("Found match for password '%s' with policy: %v", password, p)
	}
	return foundMatch
}
