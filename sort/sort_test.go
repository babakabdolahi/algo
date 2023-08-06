package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{2, 1, 0}, []int{0, 1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1, 5, 100}, []int{1, 5, 100}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{2}, []int{2}},
	}
	for _, test := range tests {
		if got := BubbleSort(test.input); !sliceComparison(test.input, got) {
			t.Errorf("BubbleSort(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{2, 1, 0}, []int{0, 1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1, 5, 100}, []int{1, 5, 100}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{2}, []int{2}},
	}
	for _, test := range tests {
		if got := SelectionSort(test.input); !sliceComparison(test.input, got) {
			t.Errorf("SelectionSort(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func sliceComparison(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}
