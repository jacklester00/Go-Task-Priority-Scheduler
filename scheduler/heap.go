package scheduler

// TaskHeap implements heap.Interface for Task structs using a max-heap.
// Tasks with higher priority numbers have higher priority in the heap.
type TaskHeap []*Task

// Len returns the number of tasks in the heap.
func (h TaskHeap) Len() int { return len(h) }

// Less reports whether the task at index i has higher priority than the task at index j.
// This implements a max-heap where higher priority numbers are "less" (closer to root).
func (h TaskHeap) Less(i, j int) bool { return h[i].Priority > h[j].Priority }

// Swap swaps the tasks at indices i and j and updates their heap indices.
func (h TaskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

// Push adds a new task to the heap and sets its index.
func (h *TaskHeap) Push(x interface{}) {
	task := x.(*Task)
	task.Index = len(*h)
	*h = append(*h, task)
}

// Pop removes and returns the last task from the heap.
// Note: heap.Pop calls this after heap.Fix, so this returns the correct task.
func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	task := old[n-1]
	task.Index = -1 // for safety
	*h = old[0 : n-1]
	return task
}
