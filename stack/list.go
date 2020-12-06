package stack

type node struct {
	value string
	next  *node
}

type stacklist struct {
	length int
	last   *node
	first  *node
}

func newStackList() stack {
	return stacklist{
		length: 0,
		last:   nil,
		first:  nil,
	}
}

func (s stacklist) pop() (string, stack) {
	res := s.last.value
	s.length--

	s.last = s.last.next
	return res, s
}

func (s stacklist) push(v string) stack {
	newNode := node{
		value: v,
		next:  s.last,
	}

	if s.length == 0 {
		s.first = &newNode
		s.last = &newNode
	} else {
		s.last = &newNode
	}

	s.length++
	return s
}

func (s stacklist) peek() string {
	if s.length == 0 {
		return ""
	}
	return s.last.value
}

func (s stacklist) isEmpty() bool {
	return s.length == 0
}
