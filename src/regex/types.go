package regex

import (
	"container/list"
	"fmt"
)

const (
	sign = "+?*()|"
)

var (
	stateId int
)

type stack struct {
	list *list.List
}

func (s *stack) empty() bool {
	return s.list.Len() == 0
}

func NewStack() *stack {
	return &stack{list.New()}
}

func (s *stack) push(e string) {
	s.list.PushBack(e)
}

func (s *stack) peek() (string, bool) {
	r := s.list.Back()

	if r == nil {
		return "", false
	} else {
		return r.Value.(string), true
	}
}

func (s *stack) pop() (string, bool) {
	r := s.list.Back()

	if r == nil {
		return "", false
	} else {
		rs := r.Value.(string)
		s.list.Remove(r)
		return rs, true
	}
}

type state struct {
	isEnd             bool
	transition        map[string]*state
	epsilonTransition []*state
	id                int
}

func (s *state) String() string {
	var c string

	for k := range s.transition {
		c += "-" + k + "->,"
	}

	return fmt.Sprintf("id %d;end %v;transition %s", s.id, s.isEnd, c)
}

func NewState(isEnd bool) *state {
	t := map[string]*state{}
	var e []*state
	stateId++
	return &state{isEnd, t, e, stateId}
}

func (s *state) addEpsilonTransition(t *state) {
	s.epsilonTransition = append(s.epsilonTransition, t)
}

func (s *state) addSymbolTransition(c string, t *state) {
	s.transition[c] = t
}

type nfa struct {
	start *state
	end   *state
}

func (n *nfa) String() string {
	return "start:" + n.start.String() + ",end:" + n.end.String()
}

func NewEpsilon() *nfa {
	start := NewState(false)
	end := NewState(true)

	start.addEpsilonTransition(end)

	return &nfa{start, end}
}

func NewSymbol(s string) *nfa {
	start := NewState(false)
	end := NewState(true)

	start.addSymbolTransition(s, end)

	return &nfa{start, end}
}

func NewCombine(p1, p2 *nfa) *nfa {
	p1.end.isEnd = false
	p2.end.isEnd = true
	p1.end.addEpsilonTransition(p2.start)

	return &nfa{p1.start, p2.end}
}

func NewZeroOrOne(p *nfa) *nfa {
	start := NewState(false)
	end := NewState(true)

	p.end.isEnd = false

	start.addEpsilonTransition(p.start)
	start.addEpsilonTransition(end)

	p.end.addEpsilonTransition(end)

	return &nfa{start, end}
}

func NewZeroOrMore(p *nfa) *nfa {
	start := NewState(false)
	end := NewState(true)

	p.end.isEnd = false

	start.addEpsilonTransition(p.start)
	start.addEpsilonTransition(end)

	p.end.addEpsilonTransition(end)
	p.end.addEpsilonTransition(p.start)

	return &nfa{start, end}
}

func NewOneOrMore(p *nfa) *nfa {
	start := NewState(false)
	end := NewState(true)

	p.end.isEnd = false

	start.addEpsilonTransition(p.start)

	p.end.addEpsilonTransition(end)
	p.end.addEpsilonTransition(p.start)

	return &nfa{start, end}
}

func NewUnion(p1, p2 *nfa) *nfa {
	start := NewState(false)
	end := NewState(true)

	p1.end.isEnd = false
	p2.end.isEnd = false

	start.addEpsilonTransition(p1.start)
	start.addEpsilonTransition(p2.start)

	p1.end.addEpsilonTransition(end)
	p2.end.addEpsilonTransition(end)

	return &nfa{start, end}
}

type nfaStack struct {
	list *list.List
}

func (s *nfaStack) empty() bool {
	return s.list.Len() == 0
}

func NewNfaStack() *nfaStack {
	return &nfaStack{list.New()}
}

func (s *nfaStack) push(n *nfa) {
	s.list.PushBack(n)
}

func (s *nfaStack) peek() *nfa {
	r := s.list.Back()

	if r == nil {
		return nil
	} else {
		return r.Value.(*nfa)
	}
}

func (s *nfaStack) pop() *nfa {
	r := s.list.Back()

	if r == nil {
		return nil
	} else {
		rs := r.Value.(*nfa)
		s.list.Remove(r)
		return rs
	}
}

func preprocess(re string) string {
	var rs string
	for i, c := range re {
		// 字符类型：
		// 1. 字母和数字
		// 2. +
		// 3. ?
		// 4. *
		// 5. (
		// 6. )
		// 7. |

		cur := string(c)
		rs += cur

		if i < len(re)-1 {
			next := string(re[i+1])

			if cur == "(" || cur == "|" {
				continue
			}

			if next == "?" || next == "*" || next == "+" || next == ")" || next == "|" {
				continue
			}

			rs += "."
		}
	}

	return rs
}

func priority(s string) int {
	switch s {
	case "+":
		fallthrough
	case "?":
		fallthrough
	case "*":
		return 10
	case ".":
		return 8
	case "|":
		return 6
	default:
		return 0
	}
}

// infix to postfix by shunting-yard algorithm (Edsger Dijkstra)
// operator   + ? * ( ) . |
func re2post(re string) string {
	r := preprocess(re)

	s := NewStack()

	var rs string
	for _, c := range r {
		cur := string(c)

		switch cur {
		case "(":
			s.push(cur)
		case ")":
			for {
				if v, ok := s.pop(); !ok || v == "(" {
					break
				} else {
					rs += v
				}
			}
		case "+":
			fallthrough
		case "?":
			fallthrough
		case "*":
			fallthrough
		case ".":
			fallthrough
		case "|":
			for {
				if v, ok := s.peek(); ok && priority(v) >= priority(cur) {
					v, _ = s.pop()
					rs += v
				} else {
					break
				}
			}
			s.push(cur)
		default:
			rs += cur
		}
	}

	for !s.empty() {
		v, _ := s.pop()
		rs += v
	}

	return rs
}

func dfsNextState(start *state, list *[]*state, dict map[int]bool) {
	if len(start.epsilonTransition) > 0 {
		for _, v := range start.epsilonTransition {
			if _, ok := dict[v.id]; !ok {
				dfsNextState(v, list, dict)
			}
		}
	} else {
		dict[start.id] = true
		*list = append(*list, start)
	}

}
