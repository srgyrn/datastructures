package stack

type stack interface {
	pop() (string, stack)
	push(v string) stack
	peek() string
	isEmpty() bool
}