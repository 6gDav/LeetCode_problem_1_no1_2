package main

import (
	"slices"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tc := range tests {
		result := twoSum(tc.nums, tc.target)
		if result != nil && !slices.Equal(result, tc.expected) {
			t.Errorf("reverse(%d) = %d; exepted: %d", tc.nums, result, tc.expected)
		}
	}
}
