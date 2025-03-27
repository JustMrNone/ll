package main

import (
	"fmt"

	"github.com/JustMrNone/ll/doubly"
	"github.com/JustMrNone/ll/singly"
)

func main() {
	fmt.Println("=== Singly Linked List Demo ===")
	demonstrateSinglyLinkedList()

	fmt.Println("\n=== Doubly Linked List Demo ===")
	demonstrateDoublyLinkedList()
}

func demonstrateSinglyLinkedList() {
	list := singly.NewSinglyLinkedList()

	// Basic operations
	fmt.Println("Adding elements...")
	list.Append(1)
	list.Append(2)
	list.Prepend(0)
	list.Print() // Expected: 0 -> 1 -> 2

	// Insert and Delete operations
	fmt.Println("\nInserting 1.5 at index 2...")
	list.Insert(1.5, 2)
	list.Print() // Expected: 0 -> 1 -> 1.5 -> 2

	fmt.Println("Deleting value 1.5...")
	list.Delete(1.5)
	list.Print() // Expected: 0 -> 1 -> 2

	// Slice operations
	slice := list.IntoSlice()
	fmt.Printf("As slice: %v\n", slice)

	// Array operations
	array := list.IntoArray()
	fmt.Printf("As array: %v\n", array)

	// Search and Get
	fmt.Printf("Index of 1: %d\n", list.Search(1))
	if val, err := list.Get(1); err == nil {
		fmt.Printf("Value at index 1: %v\n", val)
	}

	// Stack-like operations
	list.Pop()
	fmt.Print("After Pop: ")
	list.Print()

	list.Shift()
	fmt.Print("After Shift: ")
	list.Print()

	// Advanced operations
	list.FromSlice([]any{5, 3, 4, 1, 2})
	fmt.Print("New list from slice: ")
	list.Print()

	list.Sort()
	fmt.Print("After sorting: ")
	list.Print()

	list.Reverse()
	fmt.Print("After reversal: ")
	list.Print()

	middle, _ := list.GetMiddle()
	fmt.Printf("Middle element: %v\n", middle)
}

func demonstrateDoublyLinkedList() {
	list := doubly.NewDoublyLinkedList()

	// Basic operations
	fmt.Println("Adding elements...")
	list.Append(1)
	list.Append(2)
	list.Prepend(0)
	list.Print() // Expected: 0 <-> 1 <-> 2

	// Insert and Delete operations
	fmt.Println("\nInserting 1.5 at index 2...")
	list.Insert(1.5, 2)
	list.Print() // Expected: 0 <-> 1 <-> 1.5 <-> 2

	fmt.Println("Deleting value 1.5...")
	list.Delete(1.5)
	list.Print() // Expected: 0 <-> 1 <-> 2

	// Slice operations
	slice := list.IntoSlice()
	fmt.Printf("As slice: %v\n", slice)

	// Search and validation
	if idx, err := list.Search(1); err == nil {
		fmt.Printf("Index of 1: %d\n", idx)
	}

	if err := list.Validate(); err == nil {
		fmt.Println("List structure is valid")
	}

	// Stack-like operations
	list.Pop()
	fmt.Print("After Pop: ")
	list.Print()

	list.Shift()
	fmt.Print("After Shift: ")
	list.Print()

	// Advanced operations
	list.FromSlice([]any{5, 3, 4, 1, 2})
	fmt.Print("New list from slice: ")
	list.Print()

	list.Sort()
	fmt.Print("After sorting: ")
	list.Print()

	list.Reverse()
	fmt.Print("After reversal: ")
	list.Print()

	fmt.Println("Printing in reverse order:")
	list.PrintReverse()

	// Unique elements
	list.Append(1)
	list.Append(2)
	list.Unique()
	fmt.Print("After removing duplicates: ")
	list.Print()

	// Check if empty and clear
	fmt.Printf("Is empty? %v\n", list.IsEmpty())
	list.Clear()
	fmt.Printf("Is empty after clear? %v\n", list.IsEmpty())
}
