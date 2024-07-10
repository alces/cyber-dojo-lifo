package lifo

import (
    "fmt"
)

type Stack struct {
    byContent  map[string]int
    byPosition map[int]string
}

func New() *Stack {
    return &Stack{
        byContent:  make(map[string]int),
        byPosition: make(map[int]string),
    }
}

func (s *Stack) Add(element string) error {
    if element == "" {
        return fmt.Errorf("Element cannot be empty")
    }
    
    index := s.Size()
    
    if _, ok := s.byContent[element]; ok {
        index--
    }
    
    s.byContent[element] = index
    s.byPosition[index] = element
    
    return nil
}

func (s *Stack) Get(index int) (string, error) {
    item, ok := s.byPosition[s.Size() - index - 1]    
    if !ok {
        return "", fmt.Errorf("Not found")
    }
    
    return item, nil
}

func (s *Stack) Size() int {
    return len(s.byContent)
}
