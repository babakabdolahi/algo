// A simple package implementing different kinds of sort algorithms for integers.
package sort

func BubbleSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				temp := s[j+1]
				s[j+1] = s[j]
				s[j] = temp
			}
		}
	}
	return s
}

func SelectionSort(s []int) []int {
	var lowestIndex int
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[i] {
				lowestIndex = j
			}
		}
		temp := s[i]
		s[i] = s[lowestIndex]
		s[lowestIndex] = temp
	}
	return s
}

func InsertionSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		temp := s[i]
		for j := i - 1; j >= 0; j-- {
			if s[j] > temp {
				s[j+1] = s[j]
			} else {
				s[j+1] = temp
				break
			}
		}
	}
	return s
}

func Partition(s []int) []int {
	lastIndex := len(s) - 1
	leftIndex := 0
	rightIndex := lastIndex
	pivot := s[lastIndex]

	for {
		for {
			if s[leftIndex] >= pivot || leftIndex == lastIndex {
				break
			} else {
				leftIndex++
			}
		}
		for {
			if s[rightIndex] <= pivot || rightIndex < 0 {
				break
			} else {
				rightIndex--
			}
		}
		if leftIndex >= rightIndex {
			s[leftIndex], s[lastIndex] = pivot, s[leftIndex]
			break
		} else {
			s[leftIndex], s[rightIndex] = s[rightIndex], s[leftIndex]
			leftIndex++
			rightIndex--
		}
	}
	return s
}
