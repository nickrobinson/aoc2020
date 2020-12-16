package main

import "testing"

func TestMemoryGame(t *testing.T) {
	mg := NewMemoryGame([]int{0, 3, 6})
	result := mg.Step()
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
	result = mg.Step()
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
	result = mg.Step()
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
	result = mg.Step()
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
