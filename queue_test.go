package goqueue

import "testing"

func TestNewQueue(t *testing.T) {
    for i := -1; i < 5000; i++ {
        q := NewQueue(i)
        if q.bound != i {
            t.Errorf("The bound was not matched. %d should be %d", q.bound, i)
        }
    }
}

func TestEnqueue(t *testing.T) {
    q := NewQueue(5)
    for i := 0; i <= q.bound; i++ {
        // we are asking to enqueue more than the bound.
        // everything past the bound should be ignored though
        q.Enqueue(q.length)
    }
    if q.head.data != 0 {
        t.Errorf("Head of 0|1|2|3|4 expected 0, got %d", q.head.data)
    }
    if q.tail.data != 4 {
        t.Errorf("Head of 0|1|2|3|4 expected 4, got %d", q.tail.data)
    }
    q = NewQueue(50000)
    for q.length < q.bound {
        q.Enqueue(q.length)
    }
    if q.length != 50000 {
        t.Errorf("Queue of length 50000 expected, got %d", q.length)
    }
}

func TestDequeue(t *testing.T) {
    var d int
    q := NewQueue(50000)
    for q.length < q.bound {
        q.Enqueue(q.length)
    }
    for q.length > 0 {
        d = q.Dequeue().(int)
        if d != (q.bound - q.length - 1) {
            t.Errorf("Expected to dequeue %d, got %d", q.bound-q.length-1, d)
        }
    }
}
