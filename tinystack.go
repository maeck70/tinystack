package tinystack

import "fmt"

// INITIAL_STACK_CAPACITY is the initial size of the stack
// It will automatically grow if needed
const INITIAL_STACK_CAPACITY = 10

// Stack_t is a stack data structure
type Stack_t struct {
	data []string
	top  *string
}

// NewStack creates a new stack of strings
func NewStack() *Stack_t {
	return NewStackCap(INITIAL_STACK_CAPACITY)
}

// NewStackCap creates a new stack of strings with a specified capacity
func NewStackCap(capacity int) *Stack_t {
	return &Stack_t{
		data: make([]string, 0, capacity),
		top:  nil,
	}
}

// Push adds an element to the top of the stack
func (s *Stack_t) Push(element string) {
	s.data = append(s.data, element)
	s.top = &s.data[len(s.data)-1]
}

// Pop removes and returns the top element of the stack
func (s *Stack_t) Pop() (string, error) {
	if s.top == nil {
		return "", fmt.Errorf("stack is empty")
	}

	element := *s.top
	s.data = s.data[:len(s.data)-1]
	if len(s.data) > 0 {
		s.top = &s.data[len(s.data)-1]
	} else {
		s.top = nil
	}

	return element, nil
}

// PeekWithErr returns the top element of the stack without removing it
// It returns an error if the stack is empty
func (s *Stack_t) PeekWithErr() (string, error) {
	if s.top == nil {
		return "", fmt.Errorf("stack is empty")
	}

	return *s.top, nil
}

// Peek returns the top element of the stack without removing it
// It returns an empty string if the stack is empty, use PeekWithErr to get an error when empty
func (s *Stack_t) Peek() string {
	if s.top == nil {
		return ""
	}

	return *s.top
}

// IsEmpty returns true if the stack is empty
func (s *Stack_t) IsEmpty() bool {
	r := s.top == nil
	return r
}

// Size returns the number of elements in the stack
func (s *Stack_t) Size() int {
	return len(s.data)
}

// Copy returns a new stack with the same elements as the original stack
func (sourceStack *Stack_t) Copy() *Stack_t {
	newStack := NewStack()
	newStack.data = append(newStack.data, sourceStack.data...)
	newStack.top = sourceStack.top
	return newStack
}
