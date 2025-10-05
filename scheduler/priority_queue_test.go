package scheduler

import (
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()

	if pq == nil {
		t.Fatal("NewPriorityQueue() returned nil")
	}

	if pq.Len() != 0 {
		t.Errorf("Expected empty queue, got length %d", pq.Len())
	}
}

func TestPriorityQueue_Add(t *testing.T) {
	pq := NewPriorityQueue()

	task1 := &Task{Name: "task1", Priority: 5}
	task2 := &Task{Name: "task2", Priority: 10}
	task3 := &Task{Name: "task3", Priority: 3}

	pq.Add(task1)
	pq.Add(task2)
	pq.Add(task3)

	if pq.Len() != 3 {
		t.Errorf("Expected queue length 3, got %d", pq.Len())
	}
}

func TestPriorityQueue_Next(t *testing.T) {
	pq := NewPriorityQueue()

	// Test empty queue
	if task := pq.Next(); task != nil {
		t.Errorf("Expected nil from empty queue, got %v", task)
	}

	// Add tasks with different priorities
	pq.Add(&Task{Name: "low", Priority: 1})
	pq.Add(&Task{Name: "high", Priority: 10})
	pq.Add(&Task{Name: "medium", Priority: 5})

	// Should return highest priority task first
	task := pq.Next()
	if task == nil {
		t.Fatal("Expected task from non-empty queue")
	}

	if task.Name != "high" || task.Priority != 10 {
		t.Errorf("Expected high priority task, got %v", task)
	}

	// Queue should now have 2 tasks
	if pq.Len() != 2 {
		t.Errorf("Expected queue length 2, got %d", pq.Len())
	}
}

func TestPriorityQueue_Peek(t *testing.T) {
	pq := NewPriorityQueue()

	// Test empty queue
	if task := pq.Peek(); task != nil {
		t.Errorf("Expected nil from empty queue, got %v", task)
	}

	// Add tasks
	pq.Add(&Task{Name: "low", Priority: 1})
	pq.Add(&Task{Name: "high", Priority: 10})

	// Peek should return highest priority without removing
	task := pq.Peek()
	if task == nil {
		t.Fatal("Expected task from non-empty queue")
	}

	if task.Name != "high" || task.Priority != 10 {
		t.Errorf("Expected high priority task, got %v", task)
	}

	// Queue should still have 2 tasks
	if pq.Len() != 2 {
		t.Errorf("Expected queue length 2, got %d", pq.Len())
	}
}

func TestPriorityQueue_Len(t *testing.T) {
	pq := NewPriorityQueue()

	if pq.Len() != 0 {
		t.Errorf("Expected length 0, got %d", pq.Len())
	}

	pq.Add(&Task{Name: "task1", Priority: 5})
	if pq.Len() != 1 {
		t.Errorf("Expected length 1, got %d", pq.Len())
	}

	pq.Add(&Task{Name: "task2", Priority: 3})
	if pq.Len() != 2 {
		t.Errorf("Expected length 2, got %d", pq.Len())
	}
}

func TestPriorityQueue_Tasks(t *testing.T) {
	pq := NewPriorityQueue()

	// Test empty queue
	tasks := pq.Tasks()
	if len(tasks) != 0 {
		t.Errorf("Expected empty tasks slice, got length %d", len(tasks))
	}

	// Add tasks
	task1 := &Task{Name: "task1", Priority: 5}
	task2 := &Task{Name: "task2", Priority: 10}
	pq.Add(task1)
	pq.Add(task2)

	tasks = pq.Tasks()
	if len(tasks) != 2 {
		t.Errorf("Expected tasks slice length 2, got %d", len(tasks))
	}
}

func TestPriorityQueue_Ordering(t *testing.T) {
	pq := NewPriorityQueue()

	// Add tasks in random order
	tasks := []*Task{
		{Name: "low1", Priority: 1},
		{Name: "high", Priority: 10},
		{Name: "medium", Priority: 5},
		{Name: "low2", Priority: 2},
		{Name: "highest", Priority: 15},
	}

	for _, task := range tasks {
		pq.Add(task)
	}

	// Extract all tasks and verify they come out in priority order
	extracted := make([]*Task, 0, len(tasks))
	for pq.Len() > 0 {
		extracted = append(extracted, pq.Next())
	}

	// Check that priorities are in descending order
	for i := 1; i < len(extracted); i++ {
		if extracted[i-1].Priority < extracted[i].Priority {
			t.Errorf("Tasks not in priority order: %v came before %v",
				extracted[i-1], extracted[i])
		}
	}

	// Verify we got the highest priority task first
	if extracted[0].Name != "highest" || extracted[0].Priority != 15 {
		t.Errorf("Expected highest priority task first, got %v", extracted[0])
	}
}

func TestPriorityQueue_HeapProperty(t *testing.T) {
	pq := NewPriorityQueue()

	// Add many tasks to test heap property
	for i := 0; i < 100; i++ {
		pq.Add(&Task{Name: "task", Priority: i})
	}

	// Verify heap property: parent priority >= child priority
	verifyHeapProperty(t, pq.tasks, 0)
}

func verifyHeapProperty(t *testing.T, h TaskHeap, i int) {
	if i >= len(h) {
		return
	}

	left := 2*i + 1
	right := 2*i + 2

	if left < len(h) && h.Less(left, i) {
		t.Errorf("Heap property violated: parent %d has lower priority than left child %d",
			h[i].Priority, h[left].Priority)
	}

	if right < len(h) && h.Less(right, i) {
		t.Errorf("Heap property violated: parent %d has lower priority than right child %d",
			h[i].Priority, h[right].Priority)
	}

	verifyHeapProperty(t, h, left)
	verifyHeapProperty(t, h, right)
}
