package regex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {

	s := NewStack()
	assert.True(t, s.empty())

	s.push("a")
	s.push("b")
	s.push("c")
	s.push("d")

	var v string
	var ok bool
	v, ok = s.pop()
	assert.True(t, ok)
	assert.Equal(t, "d", v)
	v, ok = s.pop()
	assert.True(t, ok)
	assert.Equal(t, "c", v)
	v, ok = s.pop()
	assert.True(t, ok)
	assert.Equal(t, "b", v)
	v, ok = s.pop()
	assert.True(t, ok)
	assert.Equal(t, "a", v)
	v, ok = s.pop()
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestQueue(t *testing.T) {

	s := NewQueue()
	assert.True(t, s.empty())

	s.push("a")
	s.push("b")
	s.push("c")
	s.push("d")

	var v string
	var ok bool
	v, ok = s.pop()
	assert.True(t, ok)
	assert.Equal(t, "a", v)
	v, ok = s.pop()
	assert.True(t, ok)
	assert.Equal(t, "b", v)
	v, ok = s.pop()
	assert.True(t, ok)
	assert.Equal(t, "c", v)
	v, ok = s.pop()
	assert.True(t, ok)
	assert.Equal(t, "d", v)
	v, ok = s.pop()
	assert.False(t, ok)
	assert.Equal(t, "", v)
}
