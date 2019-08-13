package regex

import "fmt"
import "testing"
import "time"

func TestCharacters(t *testing.T) {
	tc := []struct {
		re     string
		target string
	}{
		{
			re:     "(a|b)*a",
			target: "ab",
		},
		{
			re:     "abcdefg",
			target: "abcdefg",
		},
		{
			re:     "",
			target: "",
		},
		{
			re:     "   ",
			target: " ",
		},
		{
			re:     "(a*)(a*)(a*)(a*)(a*)",
			target: "a",
		},
	}

	for _, c := range tc {
		rs := charactersDupRemoval(c.re)

		if rs != c.target {
			t.Fatalf("expected %s and actual %s", c.target, rs)
		}
	}
}

func TestDfaMatch(t *testing.T) {
	for _, c := range regcases {
		//fmt.Printf("========DFA:start check re %s and target %s============\n", c.re, c.target)
		n := post2nfa(re2post(c.re))

		d := nfa2dfa(charactersDupRemoval(c.re), n)

		start := time.Now()
		r := d.match(c.target)

		if r != c.match {
			t.Fatalf("expected %v for re %s and target %s,actual %v\n", c.match, c.re, c.target, r)
		}

		cost := time.Now().Sub(start).Nanoseconds()

		fmt.Printf("========DFA:case [%s:%s] is ok and cost %d ns================\n", c.re, c.target, cost)
	}
}
