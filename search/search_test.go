package search

import "testing"

func TestLinearSearch(t *testing.T) {
	s := []int{1, 5, 7}
	k := 7
	if !LinearSearch(s, k) {
		t.Errorf("LinearSearch(%v, %d) = false", s, k)
	}
}

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		inputSlice []int
		searchFor  int
		want       bool
	}{
		{[]int{1, 3, 6}, 3, true},
		{[]int{1, 4, 5, 7}, 1, true},
		{[]int{}, 0, false},
		{[]int{1, 4, 7}, 2, false},
		{[]int{1, 4}, 0, false},
		{[]int{1, 4}, 8, false},
		{[]int{1}, 1, true},
		{[]int{1}, 3, false},
	}
	for _, test := range tests {
		if got := BinarySearch(test.inputSlice, test.searchFor); got != test.want {
			t.Errorf("BinarySearch(%v, %d) = %v", test.inputSlice, test.searchFor, got)
		}
	}
}
