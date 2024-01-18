package utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFifoInt(t *testing.T) {
	fifo := &FIFO[int]{}

	fifo.Push(1)
	fifo.Push(2)
	fifo.Push(3)

	assert.Equal(t, 1, fifo.Pop())
	assert.Equal(t, 2, fifo.Pop())
	assert.Equal(t, 3, fifo.Pop())
}

func TestFifoString(t *testing.T) {
	fifo := &FIFO[string]{}

	fifo.Push("1")
	fifo.Push("2")
	fifo.Push("3")

	assert.Equal(t, "1", fifo.Pop())
	assert.Equal(t, "2", fifo.Pop())
	assert.Equal(t, "3", fifo.Pop())
}

func TestFifoIsEmpty(t *testing.T) {
	fifo := &FIFO[string]{}
	assert.True(t, fifo.IsEmpty())

	fifo.Push("1")
	assert.False(t, fifo.IsEmpty())

	assert.Equal(t, "1", fifo.Pop())
	assert.True(t, fifo.IsEmpty())
}
