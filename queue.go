package goqueue

type AnyType interface{}

type Node struct {
    next *Node
    data AnyType
}

type Queue struct {
    head          *Node
    tail          *Node
    length, bound int
}

func (q *Queue) Enqueue(data AnyType) {
    pushed := &Node{nil, data}
    if q.length == q.bound {
        return
    }
    if q.length == 0 {
        q.head, q.tail = pushed, pushed
    } else {
        q.tail.next = pushed
        q.tail = pushed
    }
    q.length++
}

func (q *Queue) Dequeue() AnyType {
    data := q.head.data
    q.head = q.head.next
    q.length--
    return data
}

func (q *Queue) Len() int {
    return q.length
}

func NewQueue(bound int) *Queue {
    return &Queue{nil, nil, 0, bound}
}
