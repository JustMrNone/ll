// Package singly implements a singly linked list data structure.
package singly

import "fmt"

// Node represents a node in the singly linked list.
type Node struct {
	Value any   // Value stored in the node
	Next  *Node // Pointer to the next node
}

// LinkedList represents a singly linked list data structure.
type LinkedList struct {
	Head *Node // First node in the list
	Size int   // Number of nodes in the list
}

// NewSinglyLinkedList creates and returns an empty singly linked list.
func NewSinglyLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Size: 0,
	}
}

// Print displays the list elements from head to tail.
func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Value)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Print("\n")
}

// Length returns the number of nodes in the list.
func (ll *LinkedList) Length() int {
	return ll.Size
}

// IsEmpty returns true if the list has no elements.
func (ll *LinkedList) IsEmpty() bool {
	return ll.Size == 0
}

// Clear removes all elements from the list.
func (ll *LinkedList) Clear() {
	ll.Head = nil
	ll.Size = 0
}

// Prepend adds a new node with the given value at the beginning of the list.
func (ll *LinkedList) Prepend(value any) {
	newNode := &Node{Value: value, Next: ll.Head}
	ll.Head = newNode
	ll.Size++
}

// Append adds a new node with the given value at the end of the list.
func (ll *LinkedList) Append(value any) {
	newNode := &Node{Value: value, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
	} else {
		lastNode := ll.Head
		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}
		lastNode.Next = newNode
	}
	ll.Size++
}

// IntoSlice converts the list into a slice.
func (ll *LinkedList) IntoSlice() []any {
	current := ll.Head
	var retSlice []any
	for current != nil {
		retSlice = append(retSlice, current.Value)
		current = current.Next
	}
	return retSlice
}

// FromSlice creates a list from the given slice.
func (ll *LinkedList) FromSlice(slice []any) error {
	if slice == nil {
		return fmt.Errorf("cannot create list from nil slice")
	}
	ll.Clear()
	for _, value := range slice {
		ll.Append(value)
	}
	return nil
}

// IntoArray converts the list into a fixed-size array.
func (ll *LinkedList) IntoArray() []any {
	arr := make([]any, ll.Size)
	current := ll.Head
	for i := 0; current != nil; i++ {
		arr[i] = current.Value
		current = current.Next
	}
	return arr
}

// FromArray creates a list from the given array.
func (ll *LinkedList) FromArray(arr []any) error {
	if arr == nil {
		return fmt.Errorf("cannot create list from nil array")
	}
	ll.Clear()
	for _, value := range arr {
		ll.Append(value)
	}
	return nil
}

// Search finds the first occurrence of a value in the list and returns its index.
func (ll *LinkedList) Search(value any) int {
	index := 0
	current := ll.Head
	for current != nil {
		if current.Value == value {
			return index
		}
		current = current.Next
		index++
	}
	return -1
}

// Shift removes and returns the first element from the list.
func (ll *LinkedList) Shift() error {
	if ll.Head == nil {
		return fmt.Errorf("list is empty, nothing to delete")
	}
	ll.Head = ll.Head.Next
	ll.Size--
	return nil
}

// Pop removes and returns the last element from the list.
func (ll *LinkedList) Pop() error {
	if ll.Head == nil {
		return fmt.Errorf("list is empty nothing to delete")
	}
	// If there's only one node in the list
	if ll.Head.Next == nil {
		ll.Head = nil
		ll.Size--
		return nil
	}
	// Traverse to the second-to-last node
	current := ll.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	// Remove the last node
	current.Next = nil
	ll.Size--
	return nil
}

// Delete removes the first occurrence of the specified value from the list.
func (ll *LinkedList) Delete(value any) error {
	if ll.Head == nil {
		return fmt.Errorf("list is empty nothing to delete")
	}

	// If the value is in the head node
	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		ll.Size--
		return nil
	}
	// Traverse the list to find the node to delete
	current := ll.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}
	// If the value was not found
	if current.Next == nil {
		return fmt.Errorf("value not found in the list")
	}
	// Delete the node
	current.Next = current.Next.Next
	ll.Size--
	return nil
}

// Insert adds a new value at the specified index.
func (ll *LinkedList) Insert(value any, index int) error {
	if index < 0 || index > ll.Size {
		return fmt.Errorf("index out of bounds")
	}
	if index == 0 {
		ll.Append(value)
		return nil
	}
	current := ll.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	newNode := &Node{Value: value, Next: current.Next}
	current.Next = newNode
	ll.Size++
	return nil
}

// DeleteAt removes the element at the specified index.
func (ll *LinkedList) DeleteAt(index int) error {
	if index < 0 || index > ll.Size {
		return fmt.Errorf("index out of bounds")
	}
	if index == 0 {
		ll.Shift()
		return nil
	}

	current := ll.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
	ll.Size--
	return nil
}

// Get returns the value at the specified index.
func (ll *LinkedList) Get(index int) (any, error) {
	if index < 0 || index > ll.Size {
		return nil, fmt.Errorf("index out of bounds")
	}
	current := ll.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value, nil
}

// Reverse reverses the order of elements in the list.
func (ll *LinkedList) Reverse() {
	var prev *Node
	current := ll.Head
	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	ll.Head = prev
}

// Contains checks if a value exists in the list.
func (ll *LinkedList) Contains(value any) bool {
	current := ll.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

// Unique removes duplicate values from the list.
func (ll *LinkedList) Unique() error {
	if ll.Head == nil {
		return fmt.Errorf("list is empty")
	}
	if ll.Size <= 1 {
		return nil
	}

	visited := make(map[any]bool)
	current := ll.Head
	visited[current.Value] = true
	for current.Next != nil {
		if visited[current.Next.Value] {
			current.Next = current.Next.Next
			ll.Size--
		} else {
			visited[current.Next.Value] = true
			current = current.Next
		}
	}
	return nil
}

// Merge combines the current list with another list.
func (ll *LinkedList) Merge(list *LinkedList) error {
	if list == nil {
		return fmt.Errorf("cannot merge with nil list")
	}
	current := list.Head
	for current != nil {
		ll.Append(current.Value)
		current = current.Next
	}
	return nil
}

// PrintReverse displays the list elements from tail to head.
func (ll *LinkedList) PrintReverse() {
	ll.reverseHelper(ll.Head)
	fmt.Print("\n")
}

// Sort orders the elements in the list (supports int, string, float64).
func (ll *LinkedList) Sort() error {
	if ll.Size <= 1 {
		return nil
	}

	swapped := true
	for swapped {
		swapped = false
		current := ll.Head

		for current.Next != nil {
			// Type assertion for comparison
			switch x := current.Value.(type) {
			case int:
				y, ok := current.Next.Value.(int)
				if !ok {
					return fmt.Errorf("mismatched types in list")
				}
				if x > y {
					current.Value, current.Next.Value = current.Next.Value, current.Value
					swapped = true
				}
			case string:
				y, ok := current.Next.Value.(string)
				if !ok {
					return fmt.Errorf("mismatched types in list")
				}
				if x > y {
					current.Value, current.Next.Value = current.Next.Value, current.Value
					swapped = true
				}
			case float64:
				y, ok := current.Next.Value.(float64)
				if !ok {
					return fmt.Errorf("mismatched types in list")
				}
				if x > y {
					current.Value, current.Next.Value = current.Next.Value, current.Value
					swapped = true
				}
			default:
				return fmt.Errorf("unsupported type for sorting")
			}
			current = current.Next
		}
	}
	return nil
}

// hasCycle detects if the list contains a cycle using Floyd's algorithm.
func (ll *LinkedList) hasCycle() bool {
	if ll.Head == nil || ll.Head.Next == nil {
		return false
	}

	slow := ll.Head
	fast := ll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// GetMiddle returns the middle element of the list.
func (ll *LinkedList) GetMiddle() (any, error) {
	if ll.Head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	slow := ll.Head
	fast := ll.Head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Value, nil
}

// Validate checks the integrity of the list structure.
func (ll *LinkedList) Validate() error {
	if ll.Head == nil && ll.Size != 0 {
		return fmt.Errorf("empty list has non-zero size")
	}

	// Count nodes to verify Size
	count := 0
	current := ll.Head
	for current != nil {
		count++
		if count > ll.Size {
			return fmt.Errorf("list contains more nodes than Size indicates")
		}
		current = current.Next
	}

	if count != ll.Size {
		return fmt.Errorf("actual node count (%d) differs from Size (%d)", count, ll.Size)
	}

	// Check for cycles in the list
	if ll.hasCycle() {
		return fmt.Errorf("list contains a cycle")
	}

	return nil
}

// reverseHelper is a recursive helper function for PrintReverse.
func (ll *LinkedList) reverseHelper(node *Node) {
	if node == nil {
		return
	}

	ll.reverseHelper(node.Next)
	fmt.Print(node.Value)
	if node != ll.Head {
		fmt.Print(" -> ")
	}
}
