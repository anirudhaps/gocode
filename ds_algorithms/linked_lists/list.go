package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	data int
	next *Node
}

func CreateNode(data int) *Node {
	node := new(Node) //allocates space for a Node and returns a pointer to it
	node.data = data
	node.next = nil // zero value of a pointer is nil
	return node
}

func AddNode(node *Node, head *Node, atBegin bool) *Node {
	if node == nil {
		return nil // invalid node for adding
	}

	if head == nil {
		head = node
	} else if atBegin {
		node.next = head
		head = node
	} else {
		var save *Node = head
		for save.next != nil {
			save = save.next
		}
		save.next = node
	}
	return head
}

func DisplayList(head *Node) {
	if head == nil {
		fmt.Println("Error: Invalid list")
		return
	}

	for head != nil {
		fmt.Print(head.data)
		if head.next != nil {
			fmt.Print("->")
		}
		head = head.next
	}
	fmt.Printf("\n")
}

func MiddleNode(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	var one *Node = head
	var two *Node = head.next
	for two != nil {
		one = one.next
		two = two.next
		if two == nil {
			break
		}
		two = two.next
	}
	return one
}

func DeleteNonTailNode(node *Node) {
	if node == nil || node.next == nil {
		return
	}
	// swap data with next node
	var temp int = node.data
	node.data = node.next.data
	node.next.data = temp
	// delete next node
	nextNode := node.next
	node.next = nextNode.next
	nextNode.next = nil
	nextNode = nil // garbage collector will take care about freeing the space of deleted node
	// we just have to make that space as non-pointed, that is, no pointer should be pointing to it
}

func main() {
	command := "y"
	var inp int
	var head *Node
	atBegin := true

	if len(os.Args) > 2 {
		os.Exit(1)
	}

	// -e for insertion at the end
	if len(os.Args) == 2 && strings.Compare(os.Args[1], "-e") == 0 {
		atBegin = false
	}

	for strings.Compare(strings.ToLower(command), "y") == 0 {
		fmt.Print("Enter any number: ")
		fmt.Scanf("%d", &inp)

		node := CreateNode(inp)
		head = AddNode(node, head, atBegin)

		fmt.Print("More? (y/n): ")
		fmt.Scanf("%s", &command)
	}

	DisplayList(head)
	midNode := MiddleNode(head)
	fmt.Printf("Middle node data value: %d\n", midNode.data)
	DeleteNonTailNode(midNode)
	DisplayList(head)
}
