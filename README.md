# ll - LinkedList Library

This is a **LinkedList Library**, a Go-based implementation of both **singly linked lists** and **doubly linked lists**.

---

## Features

### Singly Linked List
- **Basic Operations**: Append, Prepend, Delete, Insert, and more.
- **Traversal**: Print elements from head to tail.
- **Conversions**: Convert to/from slices and arrays.
- **Search**: Find the index of a value.
- **Sorting**: Sort elements (supports `int`, `string`, and `float64`).
- **Reversal**: Reverse the order of elements.
- **Advanced Features**:
  - Detect cycles using Floyd's Cycle-Finding Algorithm.
  - Get the middle element of the list.
  - Remove duplicates.
  - Merge with another list.

### Doubly Linked List
- **Bidirectional Traversal**: Print elements from head to tail or tail to head.
- **Basic Operations**: Append, Prepend, Delete, Insert, and more.
- **Conversions**: Convert to/from slices and arrays.
- **Search**: Find the index of a value.
- **Sorting**: Sort elements (supports `int`, `string`, and `float64`).
- **Reversal**: Reverse the order of elements.
- **Advanced Features**:
  - Validate the integrity of the list structure.
  - Remove duplicates.
  - Merge with another list.

---

## Installation

To use this library, you need to have Go installed on your system. Clone the repository and initialize the module:

```bash
git clone https://github.com/JustMrNone/ll.git
cd ll
go mod tidy
```

---

## Usage

### Singly Linked List Example

```go
package main

import (
    "fmt"
    "github.com/JustMrNone/ll/singly"
)

func main() {
    list := singly.NewSinglyLinkedList()

    // Basic operations
    list.Append(1)
    list.Append(2)
    list.Prepend(0)
    list.Print() // Output: 0 -> 1 -> 2

    // Advanced operations
    list.Sort()
    list.Print() // Output: 0 -> 1 -> 2 (sorted)
}
```

### Doubly Linked List Example

```go
package main

import (
    "fmt"
    "github.com/JustMrNone/ll/doubly"
)

func main() {
    list := doubly.NewDoublyLinkedList()

    // Basic operations
    list.Append(1)
    list.Append(2)
    list.Prepend(0)
    list.Print() // Output: 0 <-> 1 <-> 2

    // Print in reverse
    list.PrintReverse() // Output: 2 <-> 1 <-> 0
}
```

---

## API Documentation

### Singly Linked List

#### Methods
- `NewSinglyLinkedList() *LinkedList`: Creates a new singly linked list.
- `Append(value any)`: Adds a value to the end of the list.
- `Prepend(value any)`: Adds a value to the beginning of the list.
- `Delete(value any)`: Removes the first occurrence of a value.
- `Insert(value any, index int)`: Inserts a value at the specified index.
- `Search(value any) int`: Returns the index of the first occurrence of a value.
- `Sort()`: Sorts the list (supports `int`, `string`, and `float64`).
- `Reverse()`: Reverses the list.
- `GetMiddle() (any, error)`: Returns the middle element of the list.

### Doubly Linked List

#### Methods
- `NewDoublyLinkedList() *LinkedList`: Creates a new doubly linked list.
- `Append(value any)`: Adds a value to the end of the list.
- `Prepend(value any)`: Adds a value to the beginning of the list.
- `Delete(value any)`: Removes the first occurrence of a value.
- `Insert(value any, index int)`: Inserts a value at the specified index.
- `Search(value any) (int, error)`: Returns the index of the first occurrence of a value.
- `Sort()`: Sorts the list (supports `int`, `string`, and `float64`).
- `Reverse()`: Reverses the list.
- `PrintReverse()`: Prints the list in reverse order.

---

## Testing

This library includes comprehensive test cases for both singly and doubly linked lists. To run the tests, use:

```bash
go test -v
```

To run benchmarks:

```bash
go test -bench=.
```

---

## Acknowledgments

Special thanks to the Go community for their excellent resources.

---

Enjoy!