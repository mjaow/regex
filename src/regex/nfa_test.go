package regex

import (
	"fmt"
	"testing"
	"time"
)

func TestRe2post(t *testing.T) {
	tc := []struct {
		re   string
		post string
	}{
		{
			re:   "a(b|c)d",
			post: "abc|.d.",
		},
		{
			re:   "a(b|c)*",
			post: "abc|*.",
		},
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

func TestNfaMatch(t *testing.T) {
	for _, c := range regcases {
		n := post2nfa(re2post(c.re))

		start := time.Now()
		r := n.match(c.target)

		if r != c.match {
			t.Fatalf("expected %v for re %s and target %s,actual %v\n", c.match, c.re, c.target, r)
		}

		cost := time.Now().Sub(start)

		fmt.Printf("========NFA:case [%s:%s] is ok and cost %d ns================\n", c.re, c.target, cost)

	}
}
