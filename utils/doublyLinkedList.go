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

	return nil
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

	return nil
}

// Update the DLL by removing the current TAIL
func (d *DoublyLinkedList) RemoveTailDLL() error {
	if d.tail == nil {
		return fmt.Errorf("RemoveTrailError: There is no TAIL node")
	}

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
