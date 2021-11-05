package utils

import "fmt"

// Doubly Linked List Code Derived from
//  https://golangbyexample.com/doubly-linked-list-golang/
type Node struct {
	data string
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	len  int
	tail *Node
	head *Node
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

	// First, connect the link that will happen after nodeToMove is moved
	nodeToMove.prev.next = nodeToMove.next
	nodeToMove.next.prev = nodeToMove.prev

	// Now move nodeToMove to the HEAD
	nodeToMove.prev = nil
	nodeToMove.next = d.head
	d.head = nodeToMove

	return nil
}

func (d *DoublyLinkedList) SearchDLL(key string) (*Node, error) {
	if d.head == nil {
		return nil, fmt.Errorf("SearchError: List is empty")
	}
	temp := d.head
	for temp != nil {
		if temp.data == key {
			return temp, nil
		}
		temp = temp.next
	}

	return nil, fmt.Errorf("SearchError: key does not exist")
}

func (d *DoublyLinkedList) AddFrontNodeDLL(data string) {
	newNode := &Node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
	return
}

func (d *DoublyLinkedList) AddEndNodeDLL(data string) {
	newNode := &Node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.len++
	return
}

// Update the DLL by removing the current HEAD
func (d *DoublyLinkedList) RemoveHeadDLL() error {
	if d.head == nil {
		return fmt.Errorf("RemoveHeadError: List is empty")
	}

	if d.tail == d.head {
		d.tail = nil
		d.head = nil
	} else {
		// Show current head
		fmt.Printf("value = %v, prev = %v, next = %v\n", d.head.data, d.head.prev, d.head.next)

		temp := d.head

		d.head.next.prev = nil
		d.head = d.head.next

		// nil out prev head to line up for GC (move to WHITE Zone)
		temp.prev = nil
		temp.next = nil
		temp.data = ""

		// Show current tail after updates
		fmt.Printf("value = %v, prev = %v, next = %v\n", d.head.data, d.head.prev, d.head.next)
	}

	d.len--
	return nil
}

// Update the DLL by removing the current TAIL
func (d *DoublyLinkedList) RemoveTailDLL() error {
	if d.tail == nil {
		return fmt.Errorf("RemoveTailError: There is no TAIL node")
	}

	if d.tail == d.head {
		d.tail = nil
		d.head = nil
	} else {
		// Show current tail
		fmt.Printf("value = %v, prev = %v, next = %v\n", d.tail.data, d.tail.prev, d.tail.next)

		temp := d.tail

		d.tail.prev.next = nil
		d.tail = d.tail.prev

		// nil out prev tail to line up for GC (move to WHITE Zone)
		temp.prev = nil
		temp.next = nil
		temp.data = ""

		// Show current tail after updates
		fmt.Printf("value = %v, prev = %v, next = %v\n", d.tail.data, d.tail.prev, d.tail.next)
	}
	d.len--

	return nil
}

func (d *DoublyLinkedList) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.head
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.prev, temp.next)
		temp = temp.next
	}
	fmt.Println()
	return nil
}

func (d *DoublyLinkedList) TraverseReverse() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.tail
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.prev, temp.next)
		temp = temp.prev
	}
	fmt.Println()
	return nil
}

func (d *DoublyLinkedList) Size() int {
	return d.len
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
