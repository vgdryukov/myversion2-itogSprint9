package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{"Size 0", 0, 0},
		{"Size 5", 5, 5},
		{"Size 100", 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			if len(result) != tt.expected {
				t.Errorf("generateRandomElements(%d) = длина %d; ожидалась длина %d",
					tt.size, len(result), tt.expected)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{42}, 42},
		{"Multiple elements", []int{1, 5, 3, 9, 2}, 9},
		{"All equal", []int{7, 7, 7, 7}, 7},
		{"Max at beginning", []int{10, 1, 2, 3}, 10},
		{"Max at end", []int{1, 2, 3, 10}, 10},
		{"Negative numbers", []int{-5, -1, -10}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			if result != tt.expected {
				t.Errorf("maximum(%v) = %d; ожидалось %d",
					tt.input, result, tt.expected)
			}
		})
	}
}
