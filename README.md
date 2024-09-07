## tinystack: A Stack Data Structure in Go

This document describes the `tinystack` package, which implements a simple Last-In-First-Out (LIFO) stack data structure in Go.

### Usage

The `tinystack` package provides functionalities for creating, pushing, popping, peeking, and checking the size and emptiness of a stack. All functions operate on a pointer to a `Stack_t` struct, which represents the stack instance.

**1. Creating a Stack:**

* `NewStack()`: Creates a new stack with the default initial capacity of 10 elements.
* `NewStackCap(capacity int)`: Creates a new stack with a specified initial capacity.

**2. Adding Elements (Pushing):**

* `Push(element string)`: Adds an element (`string`) to the top of the stack.

**3. Removing Elements (Popping):**

* `Pop() (string, error)`: Removes and returns the top element of the stack. Returns an error if the stack is empty.

**4. Peeking at the Top Element:**

* `PeekWithErr() (string, error)`: Returns the top element of the stack without removing it. Returns an error if the stack is empty.
* `Peek() string`: Returns the top element of the stack without removing it. Returns an empty string if the stack is empty.  **Use `PeekWithErr` for error handling.**

**5. Checking Stack Status:**

* `IsEmpty() bool`: Returns true if the stack is empty.
* `Size() int`: Returns the number of elements in the stack.

**6. Copying a Stack:**

* `Copy() *Stack_t`: Returns a new stack with a copy of all elements from the original stack.


### Example Usage:

```go
package main

import (
  "fmt"
  "tinystack" // Assuming tinystack source code is in your project directory
)

func main() {
  // Create a new stack
  stack := tinystack.NewStack()

  // Push some elements
  stack.Push("apple")
  stack.Push("banana")
  stack.Push("cherry")

  // Check the size of the stack
  size := stack.Size()
  fmt.Printf("Stack size: %d\n", size) // Output: Stack size: 3

  // Peek at the top element
  topElement, err := stack.PeekWithErr()
  if err != nil {
    fmt.Println(err) // Output: stack is empty (if Peek is used instead)
  } else {
    fmt.Println("Top element:", topElement) // Output: Top element: cherry
  }

  // Pop elements
  for !stack.IsEmpty() {
    element, _ := stack.Pop()
    fmt.Println("Popped:", element)
  }
  // Output: Popped: cherry
  // Output: Popped: banana
  // Output: Popped: apple
}
```

### Additional Notes

* The `Peek` function should be used cautiously as it returns an empty string if the stack is empty. Prefer `PeekWithErr` for error handling.
* The `Copy` function creates a deep copy of the stack, ensuring that modifications to the new stack do not affect the original stack.


This documentation provides a basic overview of the `tinystack` package. Feel free to explore the source code for further details on the implementation.
