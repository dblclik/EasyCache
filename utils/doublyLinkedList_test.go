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

func TestRemoveHeadNode(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"
	nodeData2 := "TestNode2"

	// Add node with `data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	TestDLL.AddFrontNodeDLL(nodeData2)
	err := TestDLL.RemoveHeadDLL()
	if err != nil || nodeData1 != TestDLL.head.data {
		t.Fatalf(`After adding TestNode1 then TestNode2 then removing head, expected TestDLL.head.data = %q, instead TestDLL.head.data = %v`, nodeData1, TestDLL.head.data)
	}
}

func TestRemoveHeadNodeSingle(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"

	// Add node with `data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	err := TestDLL.RemoveHeadDLL()
	if err != nil || nil != TestDLL.head {
		t.Fatalf(`After adding TestNode1 then removing head, expected TestDLL.head to be nil, instead TestDLL.head.data = %v`, TestDLL.head.data)
	}
}

func TestSearchSuccess(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"
	nodeData2 := "TestNode2"

	// Add node with `data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	TestDLL.AddFrontNodeDLL(nodeData2)
	nodePtr, err := TestDLL.SearchDLL(nodeData1)
	if err != nil || nodeData1 != nodePtr.data {
		t.Fatal("Unable to find node with data key 'TestNode1'")
	}
}

func TestSearchFail(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"
	nodeData2 := "TestNode2"

	// Add node with `data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	TestDLL.AddFrontNodeDLL(nodeData2)
	nodePtr, err := TestDLL.SearchDLL("TestNode123")
	if err == nil || nodePtr != nil {
		t.Fatal("Unexpectedly found a node with data key 'TestNode123'")
	}
}
