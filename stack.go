package mystack

type Stack []int

// Size return size of the stack
func (s Stack) Size() int {
	return len(s)
}

// IsEmpty ...
func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

// Push ...
func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

// Pop return false if stack is empty
func (s *Stack) Pop() (int, bool) {
	if (*s).IsEmpty() == true {
		return 0, false
	}

	ret := (*s)[(*s).Size()-1]
	*s = (*s)[:(*s).Size()-1]
	return ret, true
}
