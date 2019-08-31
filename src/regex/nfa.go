package regex

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

func (n *nfa) match(target string) bool {
	if n == nil {
		return target == ""
	}

	var curState []*state

	// 获取下一个step可能的所有状态
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
