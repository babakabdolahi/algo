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
	for {
		l := len(s)

		if l == 0 {
			return false
		}

		l = l / 2

		if k == s[l] {
			return true
		} else if k < s[l] {
			s = s[:l]
		} else {
			s = s[l:]
		}

		if len(s) == 1 {
			return k == s[0]
		}
	}
}
