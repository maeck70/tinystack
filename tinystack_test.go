package tinystack

import (
	"fmt"
	"testing"
)

func TestPop(t *testing.T) {
	s := NewStack()

	// #1 Test when the stack is empty
	_, err := s.Pop()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	// Stack is empty

	// #2 Test when the stack has one element
	s.Push("element")
	element, err := s.Pop()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if element != "element" {
		t.Errorf("Expected 'element', got %s", element)
	}
	// Stack is empty

	// #3 Test when the stack has multiple elements
	s.Push("element1")
	s.Push("element2")
	element, err = s.Pop()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if element != "element2" {
		t.Errorf("Expected 'element2', got %s", element)
	}
	// Stack has one element (element1)

	// #4 Test poping the next one
	element, err = s.Pop()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if element != "element1" {
		t.Errorf("Expected 'element1', got %s", element)
	}
	// Stack is empty

	// #5 Test pushing element3 and poping it
	s.Push("element1")
	s.Push("element3")
	element, err = s.Pop()

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if element != "element3" {
		t.Errorf("Expected 'element3', got %s", element)
	}
	// Stack has one element (element1)

	// #6 Pop the stack empty
	element, err = s.Pop()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if element != "element1" {
		t.Errorf("Expected 'element1', got %s", element)
	}

	if err != nil {
		t.Errorf("Expected error, got nil")
	}

}
func TestPeekWithErr(t *testing.T) {
	s := NewStack()

	// #1 Test when the stack is empty
	_, err := s.PeekWithErr()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	// Stack is empty

	// #2 Test when the stack has one element
	s.Push("element")
	element, err := s.PeekWithErr()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if element != "element" {
		t.Errorf("Expected 'element', got %s", element)
	}
	// Stack has one element (element)

	// #3 Test when the stack has multiple elements
	s.Push("element1")
	s.Push("element2")
	element, err = s.PeekWithErr()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if element != "element2" {
		t.Errorf("Expected 'element2', got %s", element)
	}
	// Stack has multiple elements (element2, element1)

	// #4 Test peeking again
	element, err = s.PeekWithErr()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if element != "element2" {
		t.Errorf("Expected 'element2', got %s", element)
	}
	// Stack has multiple elements (element2, element1)

	// #5 Test pushing element3 and peeking it
	s.Push("element3")
	element, err = s.PeekWithErr()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if element != "element3" {
		t.Errorf("Expected 'element3', got %s", element)
	}
	// Stack has multiple elements (element3, element2, element1)

	// #6 Peek the stack empty
	for i := s.Size(); i > 0; i-- {
		s.Pop()
	}
	_, err = s.PeekWithErr()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	// Stack is empty
}
func TestIsEmpty(t *testing.T) {
	s := NewStack()

	// #1 Test when the stack is empty
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
	// Stack is empty

	// #2 Test when the stack has one element
	s.Push("element")
	if s.IsEmpty() {
		t.Errorf("Expected stack not to be empty")
	}
	// Stack has one element

	// #3 Test when the stack has multiple elements
	s.Push("element1")
	s.Push("element2")
	if s.IsEmpty() {
		t.Errorf("Expected stack not to be empty")
	}
	// Stack has multiple elements

	// #4 Test after popping all elements
	for i := s.Size(); i > 0; i-- {
		s.Pop()
	}
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
	// Stack is empty
}
func TestSize(t *testing.T) {
	s := NewStack()

	// #1 Test when the stack is empty
	size := s.Size()
	if size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}
	// Stack is empty

	// #2 Test when the stack has one element
	s.Push("element")
	size = s.Size()
	if size != 1 {
		t.Errorf("Expected size 1, got %d", size)
	}
	// Stack has one element

	// #3 Test when the stack has multiple elements
	s.Push("element1")
	s.Push("element2")
	size = s.Size()
	if size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}
	// Stack has multiple elements

	// #4 Test after popping all elements
	s.Pop()
	s.Pop()
	size = s.Size()
	if size != 1 {
		t.Errorf("Expected size 1, got %d", size)
	}
	// Stack has one element

	// #5 Test after pushing additional elements
	s.Push("element3")
	s.Push("element4")
	size = s.Size()
	if size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}
	// Stack has multiple elements
}
func TestCopy(t *testing.T) {
	s := NewStack()
	s.Push("element1")
	s.Push("element2")
	s.Push("element3")

	// Test copying the stack
	copyStack := s.Copy()

	// Verify that the copied stack has the same elements
	for i := 3; i > 0; i-- {
		element, err := copyStack.Pop()
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
		expectedElement := fmt.Sprintf("element%d", i)
		if element != expectedElement {
			t.Errorf("Expected '%s', got '%s'", expectedElement, element)
		}
	}

	// Verify that the original stack is not modified
	if !copyStack.IsEmpty() {
		t.Errorf("Expected copy stack to be empty")
	}
	if s.IsEmpty() {
		t.Errorf("Expected original stack to be not modified")
	}
	if s.Size() != 3 {
		t.Errorf("Expected original stack to have 3 elements, found %d", s.Size())
	}
}
func TestPeek(t *testing.T) {
	s := NewStack()

	// #1 Test when the stack is empty
	expected := ""
	result := s.Peek()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
	// Stack is empty

	// #2 Test when the stack has one element
	s.Push("element")
	expected = "element"
	result = s.Peek()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
	// Stack has one element (element)

	// #3 Test when the stack has multiple elements
	s.Push("element1")
	s.Push("element2")
	expected = "element2"
	result = s.Peek()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
	// Stack has multiple elements (element2, element1)

	// #4 Test peeking again
	expected = "element2"
	result = s.Peek()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
	// Stack has multiple elements (element2, element1)

	// #5 Test pushing element3 and peeking it
	s.Push("element3")
	expected = "element3"
	result = s.Peek()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
	// Stack has multiple elements (element3, element2, element1)

	// #6 Peek the stack empty
	for i := s.Size(); i > 0; i-- {
		s.Pop()
	}
	expected = ""
	result = s.Peek()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
	// Stack is empty
}
