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
}