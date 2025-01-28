// implementation of the list data structure based on the "Data Structures And Algorithms" by A. Aho eta.
package list

import "slices"

const MaxListLength = 100

type List struct {
	Elements       []int
	LastElPosition int
}

// inserts x into position p
func (list *List) Insert(x, p int) error {
	list.Elements = slices.Insert(list.Elements, p, x)

	list.LastElPosition = len(list.Elements) - 1

	return nil
}
