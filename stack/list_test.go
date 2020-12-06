package stack

import (
	"reflect"
	"testing"
)

func Test_stacklist_push(t *testing.T) {
	n1 := createNode("first node", nil)
	n2 := createNode("second node", &n1)
	n3 := createNode("third node", &n2)
	want := stacklist{
		length: 3,
		last: &n3,
		first: &n1,
	}

	sl := newStackList()
	sl = sl.push("first node").push("second node").push("third node")

	if !reflect.DeepEqual(want, sl) {
		t.Errorf("push() failed:\nwant:\n%v\ngot:\n%v", want, sl)
	}
}

func Test_stacklist_peek(t *testing.T) {
	n1:= createNode("first node", nil)
	sl := stacklist{
		length: 2,
		last: &node{
			value: "second node",
			next:  &n1,
		},
		first: &n1,
	}

	if got, want := sl.peek(), "second node"; got != want {
		t.Errorf("peek(): %s, want: %s", got, want)
	}
}

func Test_stacklist_pop(t *testing.T) {
	sl := newStackList()
	sl = sl.push("first node").push("second node").push("third node")
	var got string

	for i := 0; i < 2; i++ {
		want := sl.peek()
		if got, sl = sl.pop(); got != want {
			t.Fatalf("pop(): %s, want: %s", got, want)
		}
	}

	n1 := createNode("first node", nil)
	want := stacklist{
		length: 1,
		last: &n1,
		first: &n1,
	}

	if !reflect.DeepEqual(want, sl) {
		t.Errorf("unexpected stack:\n%+v", sl)
	}
}

func createNode(v string, next *node) node {
	return node{
		value: v,
		next:  next,
	}
}
