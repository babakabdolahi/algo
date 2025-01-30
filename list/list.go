// implementation of the list data structure based on the "Data Structures And Algorithms" by A. Aho eta.
package list

import (
	"fmt"
	"slices"
)

type List struct {
	Elements       []int
	LastElPosition int
	MaxLength      int
}

func New(maxLength int) List {
	arr := make([]int, 0, maxLength)

	return List{
		Elements:       arr,
		LastElPosition: 0,
		MaxLength:      maxLength,
	}
}

// Inserts x into position p.
func (list *List) Insert(x, p int) error {
	if list.Elements == nil {
		return fmt.Errorf("nil list")
	}

	if list.isPositionNegative(p) {
		return fmt.Errorf("negative position")
	}

	if list.isListFull() {
		return fmt.Errorf("list is full")
	}

	if list.isPositionOutOfRange(p) {
		return fmt.Errorf("position is out of range")
	}

	list.Elements = slices.Insert(list.Elements, p, x)

	list.LastElPosition = len(list.Elements) - 1

	return nil
}

func (list *List) isPositionNegative(p int) bool {
	return p < 0
}

func (list *List) isListFull() bool {
	return list.LastElPosition == list.MaxLength-1
}

func (list *List) isPositionOutOfRange(p int) bool {
	return p > list.LastElPosition+1
}
