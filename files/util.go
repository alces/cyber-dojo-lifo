package lifo

import (
    "fmt"
)

func index(items []string, element string) (int, error) {
    for i := 0; i < len(items); i++ {
        if items[i] == element {
            return i, nil
        }
    }
    
    return 0, fmt.Errorf("Not found")
}

func remove(items []string, index int) ([]string, error) {
    size := len(items)
    
    if index >= size {
        return nil, fmt.Errorf("Index out of slice")
    }
    
    if index == 0 {
        return items[1:], nil
    } else if index == size - 1 {
        return items[:size - 1], nil
    }
    
    return []string{}, nil
}