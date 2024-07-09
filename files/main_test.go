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

func TestGet(t *testing.T) {
    stack := New()
    
    item, err := stack.Get(0)
    assert.Error(t, err)
}

func TestStackSize(t *testing.T) {
    stack := New()
    
    stack.Add("test")
    assert.Equal(t, 1, stack.Size())
}