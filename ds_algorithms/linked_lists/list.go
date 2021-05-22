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
}
