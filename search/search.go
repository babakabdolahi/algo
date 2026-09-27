// A simple package implementing different kinds of search algorithms for integers.
package search

func LinearSearch(s []int, k int) bool {
	for _, v := range s {
		if v == k {
			return true
		}
	}
	return false
}

// binary search assumes that the input array is ordered.
func BinarySearch(s []int, k int) bool {
	l := len(s)
	// return false if the input slice is empty.
	if l == 0 {
		return false
	}

	for {
		l = l / 2

		if k == s[l] {
			return true
		} else if k < s[l] {
			s = s[:l+1]
		} else {
			s = s[l:]
		}

		if l == 1 || l == 0 {
			return k == s[0]
		}
	}
}
