package main

func main() {

}

type CircularQueue struct {
	items    []int
	front    int
	rear     int
	size     int
	capacity int
}

func (q *CircularQueue) Enqueue(val int) bool {
	if q.IsFull() {
		return false
	}

	q.rear = (q.rear + 1) % q.capacity
	q.items[q.rear] = val
	q.size++
	return true
}

func (q *CircularQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	element := q.items[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return element, true
}

func (q *CircularQueue) Front() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.items[q.front], true
}

func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *CircularQueue) IsFull() bool {
	return q.size == q.capacity
}

func (q *CircularQueue) Size() int {
	return q.size
}
