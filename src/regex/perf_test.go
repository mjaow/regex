package regex

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestPerf(t *testing.T) {

	for i := 1; i <= 15; i++ {
		regex := strings.Repeat("a?", i) + strings.Repeat("a", i)
		target := strings.Repeat("a", i)
		n := post2nfa(re2post(regex))

		startNFA := time.Now()

		if !n.match(target) {
			t.Fatalf("NFA:expected match for re %s and target %s,actual not match\n", regex, target)
		}

		nfaCost := time.Now().Sub(startNFA)

		d := nfa2dfa(charactersDupRemoval(regex), n)

		startDFA := time.Now()

		if !d.match(target) {
			t.Fatalf("DFA:expected match for re %s and target %s,actual not match\n", regex, target)
		}

		dfaCost := time.Now().Sub(startDFA)

		fmt.Printf("%d:regex %s match target %s cost NFA:DFA %d:%d ns\n", i, regex, target, nfaCost, dfaCost)
	}
}
