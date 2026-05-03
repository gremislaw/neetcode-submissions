type Node struct {
	val  rune
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(el rune) {
	s.top = &Node{val: el, next: s.top}
}

func (s *Stack) Pop() *rune {
	if s.top == nil {
		return nil
	}
	val := s.top.val
	s.top = s.top.next
	return &val
}

func isValid(s string) bool {
	st := &Stack{}
	for _, el := range s {
		if el == ')' || el == ']' || el == '}' {
			p := st.Pop()
			if p == nil {
				return false
			}
			if !((*p == '(' && el == ')') || (*p == '[' && el == ']') || (*p == '{' && el == '}')) {
				return false
			}
		} else {
			st.Push(el)
		}
	}
	return st.top == nil
}