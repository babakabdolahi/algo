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
	if list.isListEmpty() {
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

// Returns the position of x
// If multiple elements exist, it'll return the position of the first one.
func (list *List) Locate(x int) (int, bool) {
	for index, value := range list.Elements {
		if value == x {
			return index, true
		}
	}

	return 0, false
}

// Returns the value at position p
func (list *List) Retrieve(p int) (int, error) {
	if list.isPositionNegative(p) {
		return 0, fmt.Errorf("negative position")
	}

	if list.isPositionOutOfRange(p) {
		return 0, fmt.Errorf("position is out of range")
	}

	return list.Elements[p], nil
}

// Deletes the value at position p
func (list *List) Delete(p int) error {
	if list.isPositionNegative(p) {
		return fmt.Errorf("negative position")
	}

	if list.isPositionOutOfRange(p) {
		return fmt.Errorf("position is out of range")
	}

	list.Elements = slices.Delete(list.Elements, p, p+1)

	return nil
}

// Returns the position following the last element
func (list *List) End() int {
	return len(list.Elements)
}

// Causes the list to become and empty one
func (list *List) MakeEmpty() {
	list.Elements = slices.Delete(list.Elements, 0, len(list.Elements))
}

// Prints elements to the default output
func (list *List) PrintElements() {
	fmt.Println(list)
}

func (list *List) isListEmpty() bool {
	return list.Elements == nil
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
