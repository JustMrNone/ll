package test

import (
	"testing"

	"github.com/JustMrNone/ll/doubly"
	"github.com/JustMrNone/ll/singly"
)

// Singly Linked List Tests
func TestSinglyLinkedList(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		list := singly.NewSinglyLinkedList()

		// Test Append
		list.Append(1)
		list.Append(2)
		list.Append(3)
		if list.Size != 3 {
			t.Errorf("Expected size 3, got %d", list.Size)
		}

		// Test Search - only one return value
		index := list.Search(2)
		if index != 1 {
			t.Errorf("Expected index 1 for value 2, got %d", index)
		}

		// Test Delete - no return value
		list.Delete(2)
		if list.Size != 2 {
			t.Errorf("Expected size 2 after delete, got %d", list.Size)
		}
	})

	t.Run("Slice Operations", func(t *testing.T) {
		list := singly.NewSinglyLinkedList()
		initialSlice := []any{1, 2, 3, 4}

		// Test FromSlice - no return value
		list.FromSlice(initialSlice)

		// Test IntoSlice
		resultSlice := list.IntoSlice()
		if len(resultSlice) != len(initialSlice) {
			t.Errorf("Expected slice length %d, got %d", len(initialSlice), len(resultSlice))
		}
	})

	t.Run("Sort Operation", func(t *testing.T) {
		list := singly.NewSinglyLinkedList()
		list.Append(3)
		list.Append(1)
		list.Append(2)

		// Sort - no return value
		list.Sort()

		result := list.IntoSlice()
		for i := 0; i < len(result)-1; i++ {
			if result[i].(int) > result[i+1].(int) {
				t.Error("List is not properly sorted")
			}
		}
	})
}

// Doubly Linked List Tests
func TestDoublyLinkedList(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		list := doubly.NewDoublyLinkedList()

		// Test Append
		err := list.Append(1)
		if err != nil {
			t.Errorf("Append failed: %v", err)
		}
		err = list.Append(2)
		if err != nil {
			t.Errorf("Append failed: %v", err)
		}

		if list.Size != 2 {
			t.Errorf("Expected size 2, got %d", list.Size)
		}

		// Test Search
		idx, err := list.Search(2)
		if err != nil || idx != 1 {
			t.Errorf("Expected index 1 for value 2, got %d", idx)
		}

		// Test Delete
		err = list.Delete(2)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
		}
		if list.Size != 1 {
			t.Errorf("Expected size 1 after delete, got %d", list.Size)
		}
	})

	t.Run("List Validation", func(t *testing.T) {
		list := doubly.NewDoublyLinkedList()
		list.Append(1)
		list.Append(2)
		list.Append(3)

		err := list.Validate()
		if err != nil {
			t.Errorf("List validation failed: %v", err)
		}
	})

	t.Run("Edge Cases", func(t *testing.T) {
		list := doubly.NewDoublyLinkedList()

		// Test operations on empty list
		_, err := list.Search(1)
		if err == nil {
			t.Error("Expected error when searching empty list")
		}

		err = list.Delete(1)
		if err == nil {
			t.Error("Expected error when deleting from empty list")
		}

		// Test with nil values
		err = list.FromSlice(nil)
		if err == nil {
			t.Error("Expected error when creating from nil slice")
		}
	})
}

// Benchmark Tests
func BenchmarkSinglyLinkedList(b *testing.B) {
	list := singly.NewSinglyLinkedList()
	b.Run("Append", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list.Append(i)
		}
	})
}

func BenchmarkDoublyLinkedList(b *testing.B) {
	list := doubly.NewDoublyLinkedList()
	b.Run("Append", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			list.Append(i)
		}
	})
}
