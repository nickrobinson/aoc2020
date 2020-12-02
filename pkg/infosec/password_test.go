package infosec

import "testing"

func TestValidateCountPolicy(t *testing.T) {
	policy := PasswordCountPolicy{MinAppearanceCount: 1, MaxAppearanceCount: 3, Character: "a"}
	if policy.ValidatePassword("abcde") == false {
		t.Errorf("Expected password to be valid, but it was marked as invalid")
	}

	if policy.ValidatePassword("abacdaea") == true {
		t.Errorf("Expected password to be invalid, but it was marked as valid")
	}
}

func TestValidatePositionPolicy(t *testing.T) {
	policy := PasswordPoisitionPoilcy{FirstCheckedPosition: 1, SecondCheckedPosition: 3, Character: "a"}
	if policy.ValidatePassword("abcde") == false {
		t.Errorf("Expected password to be valid, but it was marked as invalid")
	}

	if policy.ValidatePassword("babdcc") == true {
		t.Errorf("Expected password to be invalid, but it was marked as valid")
	}
}
