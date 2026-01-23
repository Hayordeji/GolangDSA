package main

func main() {

}

type Task struct {
	Name     string
	Priority int // Lower number = Higher priority (Min-Heap)
}

type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }

// Less determines priority. For a Min-Heap, we use <.
// If you wanted a Max-Heap, you would use > here.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push and Pop use pointer receivers because they modify the slice's length
func (pq *PriorityQueue) Push(x any) {
	item := x.(*Task)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
