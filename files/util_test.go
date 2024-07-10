import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestStackIndex(t *testing.T) {
    nonexistent := "nonexistent"
    existing    := "existing"
    
    stack := New()
    assert.Equal(t, -1, stack.index(nonexistent))
        
    stack.Add(existing)
    assert.Equal(t, 0, stack.index(existing))
    assert.Equal(t, -1, stack.index(nonexistent))
}