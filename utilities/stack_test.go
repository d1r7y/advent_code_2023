package utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackInt(t *testing.T) {
	stack := &Stack[int]{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
}

func TestStackString(t *testing.T) {
	stack := &Stack[string]{}

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")

	assert.Equal(t, "3", stack.Pop())
	assert.Equal(t, "2", stack.Pop())
	assert.Equal(t, "1", stack.Pop())
}

func TestStackIsEmpty(t *testing.T) {
	stack := &Stack[string]{}
	assert.True(t, stack.IsEmpty())

	stack.Push("1")
	assert.False(t, stack.IsEmpty())

	assert.Equal(t, "1", stack.Pop())
	assert.True(t, stack.IsEmpty())
}
