package joltage

import (
	"sort"
	"testing"
)

func TestFindJoltageDifference(t *testing.T) {
	joltages := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	difference := FindJoltageDifference(joltages)
	if difference[1] != 7 {
		t.Errorf("Expected 7 differences of 1 jolt, got %d", difference[1])
	}
	if difference[3] != 5 {
		t.Errorf("Expected 5 differences of 3 jolt, got %d", difference[3])
	}
}

func TestFindJoltageDifferenceAdvanced(t *testing.T) {
	joltages := []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
	difference := FindJoltageDifference(joltages)
	if difference[1] != 22 {
		t.Errorf("Expected 22 differences of 1 jolt, got %d", difference[1])
	}
	if difference[3] != 10 {
		t.Errorf("Expected 10 differences of 3 jolt, got %d", difference[3])
	}
}

func TestFindPossibleAdapterSets(t *testing.T) {
	visited := make(map[int]int)
	joltages := []int{0, 16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	sort.Ints(joltages)
	sets := FindPossibleAdapterSets(0, joltages, visited)
	if sets != 8 {
		t.Errorf("Expected 8 arrangements, got %d", sets)
	}
}

func TestFindPossibleAdapterSetsAdvanced(t *testing.T) {
	visited := make(map[int]int)
	joltages := []int{52, 0, 28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
	sort.Ints(joltages)
	sets := FindPossibleAdapterSets(0, joltages, visited)
	if sets != 19208 {
		t.Errorf("Expected 19208 arrangements, got %d", sets)
	}
}
