package regex

import (
	"sort"
	"strconv"
	"strings"
)

var (
	dfaId   int
	keyDict map[int]string
)

func init() {
	keyDict = make(map[int]string)
}

type stateList struct {
	stateList []*state
	isEnd     bool
	id        int
}

type dfa struct {
	T     map[string]*stateList
	start *stateList
}

func (s *stateList) key() string {
	if v, ok := keyDict[s.id]; ok {
		return v
	}
	var l []string

	for _, v := range s.stateList {
		l = append(l, strconv.Itoa(v.id))
	}

	sort.Strings(l)

	rs := strings.Join(l, "-")

	keyDict[s.id] = rs

	return rs
}

func NewStateList(curState []*state) *stateList {
	dfaId++
	return &stateList{curState, isEndState(curState), dfaId}
}

func isEndState(states []*state) bool {
	for _, s := range states {
		if s.isEnd {
			return true
		}
	}

	return false
}

func epsilon_closure(from *stateList) *stateList {
	dict := make(map[int]bool)
	var curState []*state
	for _, s := range from.stateList {
		dfsNextState(s, &curState, dict)
	}

	return NewStateList(curState)
}

func delta(from *stateList, c string) *stateList {
	var rs []*state

	for _, f := range from.stateList {
		if val, ok := f.transition[c]; ok {
			rs = append(rs, val)
		}
	}

	return NewStateList(rs)
}

func checkSign(c rune) bool {
	for _, s := range sign {
		if c == s {
			return true
		}
	}

	return false
}

func charactersDupRemoval(re string) string {
	var count [256]int
	var rs string
	for _, c := range re {

		if !checkSign(c) {
			i := int(c)

			if count[i] == 0 {
				rs += string(c)
			}

			count[i]++
		}
	}

	return rs
}

func nfa2dfa(rs string, n *nfa) *dfa {
	if n == nil {
		return nil
	}

	n0 := []*state{n.start}
	q0 := epsilon_closure(NewStateList(n0))

	existMap := make(map[string]bool)
	existMap[q0.key()] = true

	var workList []*stateList
	workList = append(workList, q0)

	T := make(map[string]*stateList)

	for len(workList) > 0 {
		q := workList[0]
		workList = workList[1:]

		for _, c := range rs {
			t := epsilon_closure(delta(q, string(c)))
			if len(t.stateList) > 0 {
				T[q.key()+"-"+string(c)] = t
				k := t.key()

				if _, ok := existMap[k]; !ok {
					existMap[k] = true
					workList = append(workList, t)
				}
			}

		}
	}

	return &dfa{T: T, start: q0}
}

func (d *dfa) match(target string) bool {
	if d == nil {
		return target == ""
	}

	cur := d.start

	for i := range target {
		ch := string(target[i])

		if next, ok := d.T[cur.key()+"-"+ch]; ok {
			cur = next
		} else {
			return false
		}

	}

	return cur.isEnd
}
