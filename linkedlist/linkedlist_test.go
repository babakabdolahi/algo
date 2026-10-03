package linkedlist_test

import (
	"testing"

	"github.com/babakabdolahi/algo.git/linkedlist"
)

func TestGeneral(t *testing.T) {
	head := linkedlist.Init(0)

	newNode1 := linkedlist.Insert(head, 1)
	newNode2 := linkedlist.Insert(newNode1, 2)
	_ = linkedlist.Insert(newNode2, 3)

	linkedlist.Traverse(head)

	linkedlist.DeleteNode(head)

	linkedlist.Traverse(head)
}
