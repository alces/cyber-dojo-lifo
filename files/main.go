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
    
    s.items = append(s.items, element)   
    s.size++
    
    return nil
}

func (s *Stack) Get(index int) (string, error) {
    if index >= s.Size() {
        return "", fmt.Errorf("Index outside of stack")
    }
    
    return s.items[index], nil
}

func (s *Stack) Size() int {
    return s.size
}