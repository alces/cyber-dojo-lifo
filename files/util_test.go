package lifo

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {    
    _, err := index([]string{}, "nonexistent")
    assert.Error(t, err)
        
    index, err := index([]string{"something", "existing"}, "existing")
    assert.Nil(t, err)
    assert.Equal(t, 1, index)
}