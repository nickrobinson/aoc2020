package customs

import (
	"testing"
)

func TestCorrectCustomsCount(t *testing.T) {
	var tests = []struct {
		name        string
		answers     string
		groupCount  int
		answerCount int
	}{
		{"Basic valid passport", "abc", 1, 3},
		{"Group validation", "abcab", 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answerCount := GetCustomsCount(tt.answers, tt.groupCount)
			if answerCount != tt.answerCount {
				t.Errorf("GetCustomsCount(%s) got %v, want %v", tt.name, answerCount, tt.answerCount)
			}
		})
	}
}
