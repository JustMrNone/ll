// Package doubly implements a doubly linked list data structure.
package doubly

import "fmt"

// Node represents a node in the doubly linked list.
type Node struct {
	Value any   // Value stored in the node
	Next  *Node // Pointer to the next node
	Prev  *Node // Pointer to the previous node
}

// LinkedList represents a doubly linked list data structure.
type LinkedList struct {
	Head *Node // First node in the list
	Tail *Node // Last node in the list
	Size int   // Number of nodes in the list
}

// NewDoublyLinkedList creates and returns an empty doubly linked list.
func NewDoublyLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

// Prepend adds a new node with the given value at the beginning of the list.
func (ll *LinkedList) Prepend(value any) error {
	newNode := &Node{Value: value, Next: ll.Head, Prev: nil}
	if ll.Head != nil {
		if ll.Head.Prev != nil {
			return fmt.Errorf("head node's prev pointer is not nil")
		}
		ll.Head.Prev = newNode
	} else {
		if ll.Tail != nil {
			return fmt.Errorf("inconsistent list state: head is nil but tail is not")
		}
		ll.Tail = newNode
	}
	ll.Head = newNode
	ll.Size++
	return nil
}

// Append adds a new node with the given value at the end of the list.
func (ll *LinkedList) Append(value any) error {
	newNode := &Node{Value: value, Next: nil, Prev: ll.Tail}
	if ll.Tail != nil {
		if ll.Tail.Next != nil {
			return fmt.Errorf("tail node's next pointer is not nil")
		}
		ll.Tail.Next = newNode
	} else {
		if ll.Head != nil {
			return fmt.Errorf("inconsistent list state: tail is nil but head is not")
		}
		ll.Head = newNode
	}
	ll.Tail = newNode
	ll.Size++
	return nil
}

// Print displays the list elements from head to tail.
func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Value)
		if current.Next != nil {
			fmt.Print(" <-> ")
		}
		current = current.Next
	}
	fmt.Print("\n")
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
		if err := ll.Append(value); err != nil {
			return fmt.Errorf("failed to append value: %v", err)
		}
	}
	return nil
}

func (ll *LinkedList) IntoArray() []any {
	arr := make([]any, ll.Size)
	current := ll.Head
	for i := 0; current != nil; i++ {
		arr[i] = current.Value
		current = current.Next
	}
	return arr
}

func (ll *LinkedList) FromArray(arr []any) error {
	if arr == nil {
		return fmt.Errorf("cannot create list from nil array")
	}

	ll.Clear()
	for _, value := range arr {
		if err := ll.Append(value); err != nil {
			return fmt.Errorf("failed to append value: %v", err)
		}
	}
	return nil
}

// Search finds the first occurrence of a value in the list and returns its index.
func (ll *LinkedList) Search(value any) (int, error) {
	index := 0
	current := ll.Head
	for current != nil {
		if current.Value == value {
			return index, nil
		}
		current = current.Next
		index++
	}
	return -1, fmt.Errorf("value not found")
}

// Shift removes the first element from the list.
func (ll *LinkedList) Shift() error {
	if ll.Head == nil {
		return fmt.Errorf("list is empty nothing to delete")
	}
	ll.Head = ll.Head.Next
	if ll.Head != nil {
		ll.Head.Prev = nil
	} else {
		ll.Tail = nil // List is now empty
	}
	ll.Size--
	return nil
}

// Pop removes the last element from the list.
func (ll *LinkedList) Pop() error {
	if ll.Head == nil {
		return fmt.Errorf("list is empty nothing to delete")
	}
	if ll.Head == ll.Tail {
		ll.Head = nil
		ll.Tail = nil
	} else {
		ll.Tail = ll.Tail.Prev
		ll.Tail.Next = nil
	}
	ll.Size--
	return nil
}

// Delete removes the first occurrence of the specified value from the list.
func (ll *LinkedList) Delete(value any) error {
	if ll.Head == nil {
		return fmt.Errorf("list is empty nothing to delete")
	}

	if ll.Head.Value == value {
		return ll.Shift()
	}

	if ll.Tail.Value == value {
		return ll.Pop()
	}

	current := ll.Head
	for current != nil && current.Value != value {
		current = current.Next
	}

	if current == nil {
		return fmt.Errorf("value not found in the list")
	}

	current.Prev.Next = current.Next
	if current.Next != nil {
		current.Next.Prev = current.Prev
	}
	ll.Size--
	return nil
}

// IsEmpty returns true if the list has no elements.
func (ll *LinkedList) IsEmpty() bool {
	return ll.Size == 0
}

// Clear removes all elements from the list.
func (ll *LinkedList) Clear() {
	ll.Head = nil
	ll.Tail = nil
	ll.Size = 0
}

// Get returns the value at the specified index.
func (ll *LinkedList) Get(index int) (any, error) {
	if index < 0 || index >= ll.Size {
		return nil, fmt.Errorf("index out of bounds")
	}
	current := ll.Head

	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value, nil
}

// Insert adds a new value at the specified index.
func (ll *LinkedList) Insert(value any, index int) error {
	if index < 0 || index > ll.Size {
		return fmt.Errorf("index out of bounds: index %d, size %d", index, ll.Size)
	}
	if index == 0 {
		ll.Prepend(value)
		return nil
	}
	if index == ll.Size {
		ll.Append(value)
		return nil
	}

	current := ll.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	newNode := &Node{
		Value: value,
		Next:  current,
		Prev:  current.Prev,
	}

	current.Prev.Next = newNode
	current.Prev = newNode
	ll.Size++

	return nil
}

// DeleteAt removes the element at the specified index..
func (ll *LinkedList) DeleteAt(index int) error {
	if index < 0 || index > ll.Size {
		return fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		ll.Shift()
		return nil
	}

	if index == ll.Size-1 {
		ll.Pop()
		return nil
	}

	current := ll.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	current.Prev.Next = current.Next
	if current.Next != nil {
		current.Next.Prev = current.Prev
	}
	ll.Size--
	return nil
}

// Reverse reverses the order of elements in the list.
func (ll *LinkedList) Reverse() error {
	if ll.Head == nil {
		return fmt.Errorf("cannot reverse empty list")
	}
	if ll.Size <= 1 {
		return nil
	}

	current := ll.Head
	ll.Tail = current

	for current != nil {
		nextTemp := current.Next
		current.Next = current.Prev
		current.Prev = nextTemp

		if nextTemp == nil {
			ll.Head = current
		}
		current = nextTemp
	}
	return nil
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

// PrintReverse displays the list elements from tail to head.
func (ll *LinkedList) PrintReverse() {
	current := ll.Tail
	for current != nil {
		fmt.Print(current.Value)
		if current.Prev != nil {
			fmt.Print(" <-> ")
		}
		current = current.Prev
	}
	fmt.Print("\n")
}

// Merge combines the current list with another list.
func (ll *LinkedList) Merge(list *LinkedList) error {
	if list == nil {
		return fmt.Errorf("cannot merge with nil list")
	}

	current := list.Head
	for current != nil {
		if err := ll.Append(current.Value); err != nil {
			return fmt.Errorf("merge failed: %v", err)
		}
		current = current.Next
	}
	return nil
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
			if current.Next != nil {
				current.Next.Prev = current
			} else {
				ll.Tail = current
			}
			ll.Size--
		} else {
			visited[current.Next.Value] = true
			current = current.Next
		}
	}
	return nil
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

// Validate checks the integrity of the list structure.
func (ll *LinkedList) Validate() error {
	// Check if empty list is valid
	if ll.Head == nil {
		if ll.Tail != nil {
			return fmt.Errorf("head is nil but tail is not")
		}
		if ll.Size != 0 {
			return fmt.Errorf("empty list has non-zero size")
		}
		return nil
	}

	// Check head node's prev pointer
	if ll.Head.Prev != nil {
		return fmt.Errorf("head node has non-nil prev pointer")
	}

	// Check tail node's next pointer
	if ll.Tail.Next != nil {
		return fmt.Errorf("tail node has non-nil next pointer")
	}

	// Count nodes and verify links
	count := 0
	current := ll.Head
	var lastNode *Node

	for current != nil {
		count++
		if count > ll.Size {
			return fmt.Errorf("list contains more nodes than Size indicates")
		}

		// Verify prev/next links
		if current.Next != nil && current.Next.Prev != current {
			return fmt.Errorf("broken bidirectional link found")
		}

		lastNode = current
		current = current.Next
	}

	// Verify size
	if count != ll.Size {
		return fmt.Errorf("actual node count (%d) differs from Size (%d)", count, ll.Size)
	}

	// Verify tail pointer
	if lastNode != ll.Tail {
		return fmt.Errorf("tail pointer does not point to last node")
	}

	return nil
}
