package queue

import (
	"reflect"
	"testing"
)

func Test_queue_enqueue(t *testing.T) {
	q := new(queue)
	want := &queue{
		length: 1,
		front:  &node{value: "first node", next: nil},
	}

	q.enqueue(want.front.value)
	if !reflect.DeepEqual(q, want) {
		t.Errorf("q.add(\"%s\") = %v, want %v", want.front.value, q, want)
		return
	}

	n2 := node{value: "second node", next: nil}
	want = &queue{
		length: 2,
		front:  &node{value: "first node", next: &n2},
	}
	q.enqueue(n2.value)
	if !reflect.DeepEqual(q, want) {
		t.Errorf("q.add(\"%s\") = %v, want %v", n2.value, q, want)
		return
	}

	n3 := node{value: "third node", next: nil}
	n2.next = new(node)
	n2.next = &n3
	want = &queue{
		length: 3,
		front:  &node{value: "first node", next: &n2},
	}

	q.enqueue(n3.value)
	if !reflect.DeepEqual(q, want) {
		t.Errorf("q.add(\"%s\") = %v, want %v", n3.value, q, want)
		return
	}
}

func Test_queue_dequeue(t *testing.T) {
	q := new(queue)

	if got := q.dequeue(); got != "" || !reflect.DeepEqual(q, new(queue)) {
		t.Errorf("empty queue.pop() failed")
		return
	}

	q.enqueue("first node")
	q.enqueue("second node")
	q.enqueue("third node")

	results := []struct {
		want      string
		wantQueue *queue
	}{
		{
			want: "first node",
			wantQueue: &queue{
				length: 2,
				front: &node{
					value: "second node",
					next: &node{
						value: "third node",
					},
				},
			},
		},
		{
			want: "second node",
			wantQueue: &queue{
				length: 1,
				front: &node{
					value: "third node",
				},
			},
		},
	}
	for _, tt := range results {
		if got := q.dequeue(); got != tt.want {
			t.Errorf("queue.pop() = %v, want %v", got, tt.want)
			return
		}

		if !reflect.DeepEqual(q, tt.wantQueue) {
			t.Errorf("queue.pop() = %v, want %v", q, tt.wantQueue)
		}
	}
}

func Test_queue_isEmpty(t *testing.T) {
	q := new(queue)

	if !q.isEmpty() {
		t.Errorf("queue.isEmpty() = %v, want %v", q.isEmpty(), true)
		return
	}

	q.enqueue("new node here")
	if q.isEmpty() {
		t.Errorf("queue.isEmpty() = %v, want %v", q.isEmpty(), false)
	}
}
