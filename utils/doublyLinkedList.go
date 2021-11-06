package utils

import "fmt"

// Doubly Linked List Code Derived from
//  https://golangbyexample.com/doubly-linked-list-golang/
type Node struct {
	Data string
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Len  int
	Tail *Node
	Head *Node
}

func InitDoublyList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Update the DLL by moving a known existing key to the HEAD
func (d *DoublyLinkedList) UpdateDLL(key string) error {
	nodeToMove, err := d.SearchDLL(key)
	if err != nil {
		return fmt.Errorf("UpdateError: Could not find node with provided key")
	}

	if nodeToMove == d.Head {
		return nil
	}

	// First, connect the link that will happen after nodeToMove is moved
	if nodeToMove.Prev != nil {
		nodeToMove.Prev.Next = nodeToMove.Next
	} else {
		// nodeToMove is HEAD, should have been caught
		fmt.Println("Had a case when nodeToMove is HEAD but didn't get caught")
	}

	if nodeToMove.Next != nil {
		nodeToMove.Next.Prev = nodeToMove.Prev
	}

	// Now move nodeToMove to the HEAD
	nodeToMove.Prev = nil
	nodeToMove.Next = d.Head
	d.Head = nodeToMove

	return nil
}

func (d *DoublyLinkedList) SearchDLL(key string) (*Node, error) {
	if d.Head == nil {
		return nil, fmt.Errorf("SearchError: List is empty")
	}
	temp := d.Head
	for temp != nil {
		if temp.Data == key {
			return temp, nil
		}
		temp = temp.Next
	}

	return nil, fmt.Errorf("SearchError: key does not exist")
}

func (d *DoublyLinkedList) AddFrontNodeDLL(Data string) {
	newNode := &Node{
		Data: Data,
	}
	if d.Head == nil {
		d.Head = newNode
		d.Tail = newNode
	} else {
		newNode.Next = d.Head
		d.Head.Prev = newNode
		d.Head = newNode
	}
	d.Len++
	return
}

func (d *DoublyLinkedList) AddEndNodeDLL(Data string) {
	newNode := &Node{
		Data: Data,
	}
	if d.Head == nil {
		d.Head = newNode
		d.Tail = newNode
	} else {
		currentNode := d.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		newNode.Prev = currentNode
		currentNode.Next = newNode
		d.Tail = newNode
	}
	d.Len++
	return
}

// Update the DLL by removing the current HEAD
func (d *DoublyLinkedList) RemoveHeadDLL() error {
	if d.Head == nil {
		return fmt.Errorf("RemoveHeadError: List is empty")
	}

	if d.Tail == d.Head {
		d.Tail = nil
		d.Head = nil
	} else {
		// Show current Head
		fmt.Printf("value = %v, Prev = %v, Next = %v\n", d.Head.Data, d.Head.Prev, d.Head.Next)

		temp := d.Head

		d.Head.Next.Prev = nil
		d.Head = d.Head.Next

		// nil out Prev Head to line up for GC (move to WHITE Zone)
		temp.Prev = nil
		temp.Next = nil
		temp.Data = ""

		// Show current Tail after updates
		fmt.Printf("value = %v, Prev = %v, Next = %v\n", d.Head.Data, d.Head.Prev, d.Head.Next)
	}

	d.Len--
	return nil
}

// Update the DLL by removing the current TAIL
func (d *DoublyLinkedList) RemoveTailDLL() error {
	if d.Tail == nil {
		return fmt.Errorf("RemoveTailError: There is no TAIL node")
	}

	if d.Tail == d.Head {
		d.Tail = nil
		d.Head = nil
	} else {
		// Show current Tail
		fmt.Printf("value = %v, Prev = %v, Next = %v\n", d.Tail.Data, d.Tail.Prev, d.Tail.Next)

		temp := d.Tail

		d.Tail.Prev.Next = nil
		d.Tail = d.Tail.Prev

		// nil out Prev Tail to line up for GC (move to WHITE Zone)
		temp.Prev = nil
		temp.Next = nil
		temp.Data = ""

		// Show current Tail after updates
		fmt.Printf("value = %v, Prev = %v, Next = %v\n", d.Tail.Data, d.Tail.Prev, d.Tail.Next)
	}
	d.Len--

	return nil
}

func (d *DoublyLinkedList) TraverseForward() error {
	if d.Head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.Head
	for temp != nil {
		fmt.Printf("value = %v, Prev = %v, Next = %v\n", temp.Data, temp.Prev, temp.Next)
		temp = temp.Next
	}
	fmt.Println()
	return nil
}

func (d *DoublyLinkedList) TraverseReverse() error {
	if d.Head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.Tail
	for temp != nil {
		fmt.Printf("value = %v, Prev = %v, Next = %v\n", temp.Data, temp.Prev, temp.Next)
		temp = temp.Prev
	}
	fmt.Println()
	return nil
}

func (d *DoublyLinkedList) Size() int {
	return d.Len
}

// func main() {
// 	doublyList := initDoublyList()
// 	fmt.Printf("Add Front Node: C\n")
// 	doublyList.AddFrontNodeDLL("C")
// 	fmt.Printf("Add Front Node: B\n")
// 	doublyList.AddFrontNodeDLL("B")
// 	fmt.Printf("Add Front Node: A\n")
// 	doublyList.AddFrontNodeDLL("A")
// 	fmt.Printf("Add End Node: D\n")
// 	doublyList.AddEndNodeDLL("D")
// 	fmt.Printf("Add End Node: E\n")
// 	doublyList.AddEndNodeDLL("E")

// 	fmt.Printf("Size of doubly linked ist: %d\n", doublyList.Size())

// 	err := doublyList.TraverseForward()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	err = doublyList.TraverseReverse()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }
