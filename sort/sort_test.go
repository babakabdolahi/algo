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

func TestInsertionSort(t *testing.T) {
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
		if got := InsertionSort(test.input); !sliceComparison(test.input, got) {
			t.Errorf("SelectionSort(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestPartition(t *testing.T) {
	var tests = []struct {
		inputSlice  []int
		firstIndex  int
		lastIndex   int
		pivotIndex  int
		outputSlice []int
	}{
		{[]int{2, 1, 0}, 0, 2, 0, []int{0, 1, 2}},
		{[]int{2, 1}, 0, 1, 0, []int{1, 2}},
		{[]int{1, 5, 100}, 0, 2, 2, []int{1, 5, 100}},
		{[]int{1, 1, 1}, 0, 2, 1, []int{1, 1, 1}},
		{[]int{2}, 0, 0, 0, []int{2}},
		{[]int{0, 5, 2, 1, 6, 3}, 0, 5, 3, []int{0, 1, 2, 3, 6, 5}},
		{[]int{20, 18, 10, 8}, 0, 3, 0, []int{8, 18, 10, 20}},
		{[]int{3, 1, 1}, 0, 2, 0, []int{1, 1, 3}},
	}
	for _, test := range tests {
		originalInputSlice := make([]int, len(test.inputSlice))
		copy(originalInputSlice, test.inputSlice)
		pivotIndex := Partition(test.inputSlice, test.firstIndex, test.lastIndex)
		if !(pivotIndex == test.pivotIndex && sliceComparison(test.inputSlice, test.outputSlice)) {
			t.Errorf("Partition(%v) = (%v, %d), want %v, %d", originalInputSlice, test.inputSlice, pivotIndex, test.outputSlice, test.pivotIndex)
		}
	}
}

func TestQuickSort(t *testing.T) {
	var tests = []struct {
		inputSlice []int
		firstIndex int
		lastIndex  int
		want       []int
	}{
		{[]int{2, 1, 0}, 0, 2, []int{0, 1, 2}},
		{[]int{2, 1}, 0, 1, []int{1, 2}},
		{[]int{1, 5, 100}, 0, 2, []int{1, 5, 100}},
		{[]int{1, 1, 1}, 0, 2, []int{1, 1, 1}},
		{[]int{2}, 0, 0, []int{2}},
		{[]int{0, 5, 2, 1, 6, 3}, 0, 5, []int{0, 1, 2, 3, 5, 6}},
		{[]int{20, 18, 10, 8}, 0, 3, []int{8, 10, 18, 20}},
		{[]int{20, 18, 10, 8, 2, 0}, 0, 5, []int{0, 2, 8, 10, 18, 20}},
		{[]int{20, 18, 8, 10, 8, 2, 0}, 0, 5, []int{0, 2, 8, 8, 10, 18, 20}},
	}
	for _, test := range tests {
		originalInputSlice := make([]int, len(test.inputSlice))
		copy(originalInputSlice, test.inputSlice)
		QuickSort(test.inputSlice, test.firstIndex, test.lastIndex)
		if !sliceComparison(test.inputSlice, test.want) {
			t.Errorf("Partition(%v) = %v, want %v", originalInputSlice, test.inputSlice, test.want)
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
