// Implementation of different Data structures.
package datastructures

type Node struct {
	Data     int
	NextNode *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Read(index int) int {
	currentNode := l.Head
	currentIndex := 0

	for currentIndex < index {
		currentNode = currentNode.NextNode
		currentIndex++
		//we reach the tail here
		if currentNode.NextNode == nil {
			break
		}
	}
	return currentNode.Data
}

func (l *LinkedList) Insert(index, value int) {
	currentNode := l.Head
	currentIndex := 0

	if index == 0 {
		newHead := Node{
			Data:     value,
			NextNode: currentNode,
		}
		l.Head = &newHead
		return
	}

	for currentIndex <= index {
		if currentIndex == index-1 {
			newNode := Node{
				Data:     value,
				NextNode: currentNode.NextNode,
			}
			currentNode.NextNode = &newNode
			break
		}
		currentNode = currentNode.NextNode
		if currentNode == nil {

		}
		currentIndex++
	}

}
