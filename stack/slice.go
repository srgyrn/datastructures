package stack

type stackslice []string

func newStackSlice() stackslice {
	return stackslice{}
}

func (s stackslice) pop() (string, stack) {
	v := s[len(s)-1]
	return v, s[:len(s)-1]
}

func (s stackslice) push(v string) stack {
	return append(s, v)
}

func (s stackslice) peek() string {
	return s[len(s)-1]
}

func (s stackslice) isEmpty() bool {
	return len(s) == 0
}

