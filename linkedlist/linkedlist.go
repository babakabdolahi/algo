package linkedlist

import "fmt"

type Node struct {
	Item int
	Next *Node
}

func Init(v int) *Node {
	return &Node{
		Next: nil,
		Item: v,
	}
}

func Insert(prevNode *Node, v int) *Node {
	newNode := Init(v)
	newNode.Next = prevNode.Next
	prevNode.Next = newNode

	return newNode
}

func DeleteNode(prevNode *Node) {
	if prevNode.Next == nil {
		return
	}

	prevNode.Next = prevNode.Next.Next
}

func Traverse(head *Node) {
	for t := head; t != nil; t = t.Next {
		fmt.Println(t.Item)
	}
}
