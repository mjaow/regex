package regex

import (
	"fmt"
	"testing"
)

func TestRe2post(t *testing.T) {
	tc := []struct {
		re   string
		post string
	}{
		{
			re:   "a(b|c)*d",
			post: "abc|*.d.",
		},
		{
			re:   "a(bc|de)(fg|hjk)*(bc)?www",
			post: "abc.de.|.fg.hj.k.|*.bc.?.w.w.w.",
		},
	}

	for _, c := range tc {
		r := re2post(c.re)

		if r != c.post {
			t.Fatalf("re2post %s.expected %s and actual %s\n", c.re, c.post, r)
		}
	}
}

func TestPreprocess(t *testing.T) {
	tc := []struct {
		re string
		rs string
	}{
		{
			re: "a(b|c)*d",
			rs: "a.(b|c)*.d",
		},
		{
			re: "a(bcd)d",
			rs: "a.(b.c.d).d",
		},
		{
			re: "abcde",
			rs: "a.b.c.d.e",
		},
		{
			re: "a(bcd)(cded)?e",
			rs: "a.(b.c.d).(c.d.e.d)?.e",
		},
		{
			re: "a",
			rs: "a",
		},
		{
			re: "",
			rs: "",
		},
	}

	for _, c := range tc {
		r := preprocess(c.re)
		if r != c.rs {
			t.Fatalf("expected %s and actual %s\n", c.rs, r)
		}
	}
}

func TestMatch(t *testing.T) {
	tc := []struct {
		re     string
		target string
		match  bool
	}{
		{
			re:     "abcdefg",
			target: "abcdefg",
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
			re:     "(a*)(a*)(a*)(a*)(a*)",
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

	for _, c := range tc {
		fmt.Printf("start check re %s and target %s\n", c.re, c.target)
		n := post2nfa(re2post(c.re))

		r := n.match(c.target)

		if r != c.match {
			t.Fatalf("expected %v for re %s and target %s,actual %v\n", c.match, c.re, c.target, r)
		}
	}
}
