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
	for i, c := range password {
		if string(c) == p.Character {
			if i+1 == p.FirstCheckedPosition || i+1 == p.SecondCheckedPosition {
				return true
			}
		}
	}
	return false
}
