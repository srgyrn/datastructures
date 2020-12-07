package queue

type node struct {
	value string
	next  *node
}

type queue struct {
	length int
	front  *node
}

func (q *queue) enqueue(v string) {
	n := node{value: v}
	q.length++

	if q.front == nil {
		q.front = &n
		return
	}

	cur := q.front
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = new(node)
	cur.next = &n
}

func (q *queue) dequeue() string {
	res := q.front.value
	q.length--
	q.front = q.front.next
	return res
}

func (q *queue) isEmpty() bool {
	return q.length == 0
}
