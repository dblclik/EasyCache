package utils

import (
	"testing"
)

func TestAddFrontNode(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData := "TestNode"

	// Add node with `data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData)
	if nodeData != TestDLL.head.data {
		t.Fatalf(`After AddFrontNodeDLL("TestNode"), expected TestDLL.head.data = %q, instead TestDLL.head.data = %v`, nodeData, TestDLL.head.data)
	}
}

func TestAddTailNode(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData := "TestNode"

	// Add node with `data := "TestNode"`
	TestDLL.AddEndNodeDLL(nodeData)
	if nodeData != TestDLL.tail.data {
		t.Fatalf(`After AddEndNodeDLL("TestNode"), expected TestDLL.tail.data = %q, instead TestDLL.tail.data = %v`, nodeData, TestDLL.tail.data)
	}
}

func TestRemoveTailNode(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"
	nodeData2 := "TestNode2"

	// Add node with `data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	TestDLL.AddFrontNodeDLL(nodeData2)
	err := TestDLL.RemoveTailDLL()
	if err != nil || nodeData2 != TestDLL.tail.data {
		t.Fatalf(`After adding TestNode1 then TestNode2 then removing tail, expected TestDLL.tail.data = %q, instead TestDLL.tail.data = %v`, nodeData2, TestDLL.tail.data)
	}
}

func TestRemoveTailNodeSingle(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"

	// Add node with `data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	err := TestDLL.RemoveTailDLL()
	if err != nil || TestDLL.tail != nil {
		t.Fatalf(`After adding TestNode1 then removing tail, expected TestDLL.tail to be nil, instead TestDLL.tail.data = %v`, TestDLL.tail.data)
	}
}
