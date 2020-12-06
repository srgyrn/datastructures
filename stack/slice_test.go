package stack

import (
	"reflect"
	"testing"
)

func Test_stackslice_isEmpty(t *testing.T) {
	s := newStackSlice()
	if got := s.isEmpty(); got != true {
		t.Errorf("isEmpty() = %v, want %v", got, true)
	}

	got := s.push("first node").isEmpty()
	if got != false {
		t.Errorf("isEmpty() = %v, want %v", got, false)
	}
}

func Test_stackslice_peek(t *testing.T) {
	s := newStackSlice().push("first node").push("second node")
	want := "second node"
	if got := s.peek(); got != want {
		t.Errorf("peek() = %v, want %v", got, want)
	}
}

func Test_stackslice_pop(t *testing.T) {
	s := newStackSlice().push("first node").push("second node").push("third node")
	var got string

	for i := 0; i < 2; i++ {
		want := s.peek()
		if got, s = s.pop(); got != want {
			t.Fatalf("pop(): %s, want: %s", got, want)
		}
	}

	want := stackslice{"first node"}
	if !reflect.DeepEqual(want, s) {
		t.Errorf("unexpected stack:\n%+v", s)
	}
}

func Test_stackslice_push(t *testing.T) {
	s := newStackSlice()
	got := s.push("first node").push("second node").push("third node")
	want := stackslice{
		"first node",
		"second node",
		"third node",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("push() got = %v, want %v", got, want)
	}
}
