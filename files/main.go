package lifo

import (
    "fmt"
)

type Stack struct {
    size int
}

func New() *Stack {
    return &Stack{}
}

func (s *Stack) Add(element string) error {
    if element == "" {
        return fmt.Errorf("Element cannot be empty")
    }
    
    s.size++
    
    return nil
}

func (s *Stack) Size() int {
    return s.size
}