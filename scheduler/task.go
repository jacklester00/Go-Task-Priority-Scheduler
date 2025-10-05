// Package scheduler provides a priority queue implementation for task scheduling.
// It uses a max-heap data structure to efficiently manage tasks by priority.
//
// Example usage:
//
//	pq := scheduler.NewPriorityQueue()
//	pq.Add(&scheduler.Task{Name: "Buy groceries", Priority: 5})
//	pq.Add(&scheduler.Task{Name: "Write code", Priority: 10})
//	
//	task := pq.Next() // Returns "Write code" (highest priority)
//	fmt.Printf("Next task: %s\n", task.Name)
package scheduler

// Task represents a single task with a name, priority, and heap index.
// Tasks with higher priority numbers are processed first.
type Task struct {
	Name     string // The name/description of the task
	Priority int    // Priority level (higher numbers = higher priority)
	Index    int    // Heap index maintained by heap.Interface
}
