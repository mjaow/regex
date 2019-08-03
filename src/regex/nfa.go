package regex

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

//operator:
// + ? * ( ) . |
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

// + ? * . | 字符 -> nfa
func post2nfa(post string) *nfa {
	stack := NewNfaStack()
	for _, c := range post {
		cur := string(c)

		var s *nfa
		switch cur {
		case "+":
			s = NewOneOrMore(stack.pop())
		case "?":
			s = NewZeroOrOne(stack.pop())
		case "*":
			s = NewZeroOrMore(stack.pop())
		case ".":
			r := stack.pop()
			l := stack.pop()
			s = NewCombine(l, r)
		case "|":
			r := stack.pop()
			l := stack.pop()
			s = NewUnion(l, r)
		default:
			s = NewSymbol(cur)
		}
		stack.push(s)
	}
	return stack.pop()
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

func (n *nfa) match(target string) bool {
	if n == nil {
		return target == ""
	}

	var curState []*state

	dfsNextState(n.start, &curState, make(map[int]bool))

	for _, c := range target {
		cur := string(c)

		var nextState []*state
		for _, s := range curState {
			if next, ok := s.transition[cur]; ok {
				dfsNextState(next, &nextState, make(map[int]bool))
			}
		}

		if len(nextState) == 0 {
			return false
		}

		curState = nextState
	}

	for _, s := range curState {
		if s.isEnd {
			return true
		}
	}

	return false
}
