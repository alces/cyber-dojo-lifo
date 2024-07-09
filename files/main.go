package lifo

import (
    "fmt"
)

type Stack struct {
    size  int
    items []string
}

func New() *Stack {
    return &Stack{}
}

func (s *Stack) Add(element string) error {
    if element == "" {
        return fmt.Errorf("Element cannot be empty")
    }
    
    index := s.index(element)
    
    if index > -1 {
        s.items = append(s.items[index - 1:], s.items[index + 1:]...)
    } else { 
        s.size++
    }
    s.items = append(s.items, element)
    
    return nil
}

func (s *Stack) Get(index int) (string, error) {
    if index >= s.Size() {
        return "", fmt.Errorf("Index outside of stack")
    }
    
    return s.items[s.Size() - index - 1], nil
}

func (s *Stack) Size() int {
    return s.size
}

func (s *Stack) index(element string) int {
    for i := 0; i < s.Size(); i++ {
        if s.items[i] == element {
            return i
        }
    }
    
    return -1
}