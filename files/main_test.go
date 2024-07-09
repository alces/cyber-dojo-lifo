package lifo

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
    stack := New()

    assert.Equal(t, 0, stack.Size())
}

func TestStackAdd(t *testing.T) {
    stack := New()
       
    assert.Error(t, stack.Add(""))
    assert.Equal(t, 0, stack.Size())
    
    assert.Nil(t, stack.Add("one"))
    assert.Equal(t, 1, stack.Size())
}

func TestStackGet(t *testing.T) {
    stack := New()
    
    _, err := stack.Get(0)
    assert.Error(t, err)
    
    firstItem := "first"
    stack.Add(firstItem)
    getItem, err := stack.Get(0)
    assert.Nil(t, err)
    assert.Equal(t, firstItem, getItem)
    
    secondItem := "second"
    stack.Add(secondItem)
    
    getItem, err = stack.Get(0)
    assert.Nil(t, err)
    assert.Equal(t, secondItem, getItem)
    
    getItem, err = stack.Get(1)
    assert.Nil(t, err)
    assert.Equal(t, firstItem, getItem)
}

func TestStackIndex(t *testing.T) {
    stack := New()
    assert.Equal(t, -1, stack.index("nonexistent"))
}
    
func TestStackSize(t *testing.T) {
    stack := New()
    
    stack.Add("test")
    assert.Equal(t, 1, stack.Size())
}