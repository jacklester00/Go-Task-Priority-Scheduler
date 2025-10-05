package scheduler

import "container/heap"

// PriorityQueue implements a priority queue using a max-heap.
// Tasks with higher priority numbers are processed first.
type PriorityQueue struct {
	tasks TaskHeap
}

// NewPriorityQueue creates and initializes a new priority queue.
func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}
	heap.Init(&pq.tasks)
	return pq
}

// Add inserts a new task into the priority queue.
func (pq *PriorityQueue) Add(task *Task) {
	heap.Push(&pq.tasks, task)
}

// Next removes and returns the highest priority task from the queue.
// Returns nil if the queue is empty.
func (pq *PriorityQueue) Next() *Task {
	if pq.Len() == 0 {
		return nil
	}
	return heap.Pop(&pq.tasks).(*Task)
}

// Peek returns the highest priority task without removing it from the queue.
// Returns nil if the queue is empty.
func (pq *PriorityQueue) Peek() *Task {
	if pq.Len() == 0 {
		return nil
	}
	return pq.tasks[0]
}

// Len returns the number of tasks in the queue.
func (pq *PriorityQueue) Len() int {
	return pq.tasks.Len()
}

// Tasks returns a slice of all tasks currently in the queue.
// The tasks are returned in heap order (not necessarily priority order).
func (pq *PriorityQueue) Tasks() []*Task {
	return []*Task(pq.tasks)
}
