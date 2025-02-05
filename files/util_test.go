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

func TestRemove(t *testing.T) {
    _, err := remove([]string{"a", "b"}, 2)
    assert.Error(t, err)
    
    result, err := remove([]string{"a", "b"}, 0)
    assert.Nil(t, err)
    assert.Equal(t, []string{"b"}, result)
    
    result, err = remove([]string{"a", "b"}, 1)
    assert.Nil(t, err)
    assert.Equal(t, []string{"a"}, result)
    
    result, err = remove([]string{"a", "b", "c"}, 1)
    assert.Nil(t, err)
    assert.Equal(t, []string{"a", "c"}, result)
}