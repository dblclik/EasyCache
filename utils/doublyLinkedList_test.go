package utils

import (
	"testing"
)

func TestAddFrontNode(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData := "TestNode"

	// Add node with `Data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData)
	if nodeData != TestDLL.Head.Data {
		t.Fatalf(`After AddFrontNodeDLL("TestNode"), expected TestDLL.Head.Data = %q, instead TestDLL.Head.Data = %v`, nodeData, TestDLL.Head.Data)
	}
}

func TestAddTailNode(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData := "TestNode"

	// Add node with `Data := "TestNode"`
	TestDLL.AddEndNodeDLL(nodeData)
	if nodeData != TestDLL.Tail.Data {
		t.Fatalf(`After AddEndNodeDLL("TestNode"), expected TestDLL.Tail.Data = %q, instead TestDLL.Tail.Data = %v`, nodeData, TestDLL.Tail.Data)
	}
}

func TestRemoveTailNode(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"
	nodeData2 := "TestNode2"

	// Add node with `Data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	TestDLL.AddFrontNodeDLL(nodeData2)
	err := TestDLL.RemoveTailDLL()
	if err != nil || nodeData2 != TestDLL.Tail.Data {
		t.Fatalf(`After adding TestNode1 then TestNode2 then removing Tail, expected TestDLL.Tail.Data = %q, instead TestDLL.Tail.Data = %v`, nodeData2, TestDLL.Tail.Data)
	}
}

func TestRemoveTailNodeSingle(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"

	// Add node with `Data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	err := TestDLL.RemoveTailDLL()
	if err != nil || TestDLL.Tail != nil {
		t.Fatalf(`After adding TestNode1 then removing Tail, expected TestDLL.Tail to be nil, instead TestDLL.Tail.Data = %v`, TestDLL.Tail.Data)
	}
}

func TestRemoveHeadNode(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"
	nodeData2 := "TestNode2"

	// Add node with `Data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	TestDLL.AddFrontNodeDLL(nodeData2)
	err := TestDLL.RemoveHeadDLL()
	if err != nil || nodeData1 != TestDLL.Head.Data {
		t.Fatalf(`After adding TestNode1 then TestNode2 then removing Head, expected TestDLL.Head.Data = %q, instead TestDLL.Head.Data = %v`, nodeData1, TestDLL.Head.Data)
	}
}

func TestRemoveHeadNodeSingle(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"

	// Add node with `Data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	err := TestDLL.RemoveHeadDLL()
	if err != nil || nil != TestDLL.Head {
		t.Fatalf(`After adding TestNode1 then removing Head, expected TestDLL.Head to be nil, instead TestDLL.Head.Data = %v`, TestDLL.Head.Data)
	}
}

func TestSearchSuccess(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"
	nodeData2 := "TestNode2"

	// Add node with `Data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	TestDLL.AddFrontNodeDLL(nodeData2)
	nodePtr, err := TestDLL.SearchDLL(nodeData1)
	if err != nil || nodeData1 != nodePtr.Data {
		t.Fatal("Unable to find node with Data key 'TestNode1'")
	}
}

func TestSearchFail(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"
	nodeData2 := "TestNode2"

	// Add node with `Data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	TestDLL.AddFrontNodeDLL(nodeData2)
	nodePtr, err := TestDLL.SearchDLL("TestNode123")
	if err == nil || nodePtr != nil {
		t.Fatal("Unexpectedly found a node with Data key 'TestNode123'")
	}
}

func TestUpdateDLL(t *testing.T) {
	var TestDLL *DoublyLinkedList = InitDoublyList()
	nodeData1 := "TestNode1"
	nodeData2 := "TestNode2"

	// Add node with `Data := "TestNode"`
	TestDLL.AddFrontNodeDLL(nodeData1)
	TestDLL.AddFrontNodeDLL(nodeData2)
	err := TestDLL.UpdateDLL(nodeData1)
	if err != nil || TestDLL.Head.Data != nodeData1 {
		t.Fatalf("After UpdateDLL, expected HEAD Data to be: %q; actual HEAD Data is: %v", nodeData1, TestDLL.Head.Data)
	}
}
