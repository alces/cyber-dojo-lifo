package lifo

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
    stack := New()

    assert.Equal(t, 0, stack.size)
}

func TestStackAdd(t *testing.T) {
    stack := New()
       
    assert.Error(t, stack.Add(""))
    assert.Equal(t, 0, stack.size)
    
    assert.Nil(t, stack.Add("one"))
    assert.Equal(t, 1, stack.size)
}

func TestStackSize(t *testing.T) {
    stack := New()
    
    stack.Add("test")
    assert.Equal(t, 1, stack.Size())
}