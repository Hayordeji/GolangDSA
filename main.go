package main

func main() {

}

type Stack struct {
	items []int
}

// Push adds element to top
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// Pop removes and returns top element
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false // Stack is empty
	}

	index := len(s.items) - 1
	element := s.items[index]
	s.items = s.items[:index] // Remove last element

	return element, true
}

// Peek returns top element without removing
func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

// IsEmpty checks if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns number of elements
func (s *Stack) Size() int {
	return len(s.items)
}
