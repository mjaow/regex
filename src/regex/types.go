package regex

import (
	"container/list"
	"fmt"
)

var (
	stateId int
)

type queue struct {
	list *list.List
}

func NewQueue() *queue {
	return &queue{list.New()}
}

func (s *queue) empty() bool {
	return s.list.Len() == 0
}

func (s *queue) push(e string) {
	s.list.PushBack(e)
}

func (s *queue) peek() (string, bool) {
	r := s.list.Front()

	if r == nil {
		return "", false
	} else {
		return r.Value.(string), true
	}
}

func (s *queue) pop() (string, bool) {
	r := s.list.Front()

	if r == nil {
		return "", false
	} else {
		rs := r.Value.(string)
		s.list.Remove(r)
		return rs, true
	}
}

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
	p1.start.isEnd = false
	p1.end.isEnd = false
	p2.start.isEnd = false
	p2.end.isEnd = true
	p1.end.addEpsilonTransition(p2.start)

	return &nfa{p1.start, p2.end}
}

func NewZeroOrOne(p *nfa) *nfa {
	start := NewState(false)
	end := NewState(true)

	p.start.isEnd = false
	p.end.isEnd = false

	start.addEpsilonTransition(p.start)
	start.addEpsilonTransition(end)

	p.end.addEpsilonTransition(end)

	return &nfa{start, end}
}

func NewZeroOrMore(p *nfa) *nfa {
	start := NewState(false)
	end := NewState(true)

	p.start.isEnd = false
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

	p.start.isEnd = false
	p.end.isEnd = false

	start.addEpsilonTransition(p.start)

	p.end.addEpsilonTransition(end)
	p.end.addEpsilonTransition(p.start)

	return &nfa{start, end}
}

func NewUnion(p1, p2 *nfa) *nfa {
	start := NewState(false)
	end := NewState(true)

	p1.start.isEnd = false
	p1.end.isEnd = false
	p2.start.isEnd = false
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
