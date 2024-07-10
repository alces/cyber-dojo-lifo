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
    if index >= len(items) {
        return nil, fmt.Errorf("Index out of slice")
    }
    
    return []string{}, nil
}