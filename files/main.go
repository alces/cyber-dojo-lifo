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
    
    index, err := index(s.items, element)    
    if err == nil {
        s.items, err = remove(s.items, index)
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
    return len(s.items)
}
