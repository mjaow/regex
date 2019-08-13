package regex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var regcases = []struct {
	re     string
	target string
	match  bool
}{
	{
		re:     "a(b|c)kk",
		target: "abkk",
		match:  true,
	},
	{
		re:     "abcdefg",
		target: "abcdefg",
		match:  true,
	},
	{
		re:     "a(b|c)*",
		target: "abbbbbbbbbbbbbbbbbbb",
		match:  true,
	},
	{
		re:     "(a|b)*a",
		target: "ababababab",
		match:  false,
	},
	{
		re:     "(a|b)*a",
		target: "aaaaaaaaba",
		match:  true,
	},
	{
		re:     "(a|b)*a",
		target: "aaaaaabac",
		match:  false,
	},
	{
		re:     "a(b|c)*d",
		target: "abccbcccd",
		match:  true,
	},
	{
		re:     "a(b|c)*d",
		target: "abccbcccde",
		match:  false,
	},
	{
		re:     "a(b|c)+d",
		target: "acd",
		match:  true,
	},
	{
		re:     "a(b|c)+d",
		target: "ad",
		match:  false,
	},
	{
		re:     "a(b|c)+d",
		target: "abbbbd",
		match:  true,
	},
	{
		re:     "a(b|c)?d",
		target: "acd",
		match:  true,
	},
	{
		re:     "a(b|c)?d",
		target: "accd",
		match:  false,
	},
	{
		re:     "a(b|c)?d",
		target: "ad",
		match:  true,
	},
	{
		re:     "(a*)(a*)(a*)(a*)(a*)",
		target: "",
		match:  true,
	},
	{
		re:     "(a*)(a*)(a*)(a*)(a*)",
		target: "a",
		match:  true,
	},
	{
		re:     "(a*)(a*)(a*)(a*)(a*)",
		target: "aaaaaaa",
		match:  true,
	},
	{
		re:     "(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)(a*)",
		target: "aba",
		match:  false,
	},
	{
		re:     "",
		target: "",
		match:  true,
	},
	{
		re:     "    ",
		target: "",
		match:  false,
	},
	{
		re:     "    ",
		target: "    ",
		match:  true,
	},
}

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
