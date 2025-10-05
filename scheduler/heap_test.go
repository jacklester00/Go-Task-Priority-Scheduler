package scheduler

import (
	"testing"
)

func TestTaskHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		heap TaskHeap
		want int
	}{
		{"empty heap", TaskHeap{}, 0},
		{"single task", TaskHeap{&Task{Name: "test", Priority: 5}}, 1},
		{"multiple tasks", TaskHeap{
			&Task{Name: "task1", Priority: 5},
			&Task{Name: "task2", Priority: 10},
			&Task{Name: "task3", Priority: 3},
		}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.heap.Len(); got != tt.want {
				t.Errorf("TaskHeap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHeap_Less(t *testing.T) {
	tests := []struct {
		name string
		heap TaskHeap
		i    int
		j    int
		want bool
	}{
		{
			"higher priority should be less (max-heap)",
			TaskHeap{
				&Task{Name: "high", Priority: 10},
				&Task{Name: "low", Priority: 5},
			},
			0, 1, true,
		},
		{
			"lower priority should not be less",
			TaskHeap{
				&Task{Name: "low", Priority: 5},
				&Task{Name: "high", Priority: 10},
			},
			0, 1, false,
		},
		{
			"equal priorities",
			TaskHeap{
				&Task{Name: "task1", Priority: 5},
				&Task{Name: "task2", Priority: 5},
			},
			0, 1, false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.heap.Less(tt.i, tt.j); got != tt.want {
				t.Errorf("TaskHeap.Less(%d, %d) = %v, want %v", tt.i, tt.j, got, tt.want)
			}
		})
	}
}

func TestTaskHeap_Swap(t *testing.T) {
	heap := TaskHeap{
		&Task{Name: "task1", Priority: 5, Index: 0},
		&Task{Name: "task2", Priority: 10, Index: 1},
	}

	heap.Swap(0, 1)

	// Check that tasks are swapped
	if heap[0].Name != "task2" || heap[1].Name != "task1" {
		t.Errorf("Tasks not swapped correctly")
	}

	// Check that indices are updated
	if heap[0].Index != 0 || heap[1].Index != 1 {
		t.Errorf("Indices not updated correctly")
	}
}

func TestTaskHeap_Push(t *testing.T) {
	heap := TaskHeap{}
	task := &Task{Name: "test", Priority: 5}

	heap.Push(task)

	if len(heap) != 1 {
		t.Errorf("Expected heap length 1, got %d", len(heap))
	}

	if heap[0] != task {
		t.Errorf("Task not added to heap")
	}

	if task.Index != 0 {
		t.Errorf("Task index not set correctly, got %d", task.Index)
	}
}

func TestTaskHeap_Pop(t *testing.T) {
	heap := TaskHeap{
		&Task{Name: "task1", Priority: 5, Index: 0},
		&Task{Name: "task2", Priority: 10, Index: 1},
	}

	originalLength := len(heap)
	task := heap.Pop().(*Task)

	if len(heap) != originalLength-1 {
		t.Errorf("Expected heap length %d, got %d", originalLength-1, len(heap))
	}

	if task.Name != "task2" {
		t.Errorf("Expected to pop 'task2', got %s", task.Name)
	}

	if task.Index != -1 {
		t.Errorf("Expected popped task index to be -1, got %d", task.Index)
	}
}
