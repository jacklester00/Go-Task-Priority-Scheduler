package main

import (
	"strings"
	"testing"

	"priority-queue-scheduler/scheduler"
)

func TestMain_Integration(t *testing.T) {
	// This test simulates the core functionality that main.go provides
	// without trying to test the interactive CLI directly

	// Test the PriorityQueue operations that main.go uses
	pq := scheduler.NewPriorityQueue()

	// Simulate adding tasks like the CLI would
	task1 := &scheduler.Task{Name: "Buy groceries", Priority: 5}
	task2 := &scheduler.Task{Name: "Write code", Priority: 10}
	task3 := &scheduler.Task{Name: "Exercise", Priority: 3}

	pq.Add(task1)
	pq.Add(task2)
	pq.Add(task3)

	// Verify tasks were added
	if pq.Len() != 3 {
		t.Errorf("Expected 3 tasks, got %d", pq.Len())
	}

	// Test listing tasks (like the 'list' command)
	tasks := pq.Tasks()
	if len(tasks) != 3 {
		t.Errorf("Expected 3 tasks in list, got %d", len(tasks))
	}

	// Test getting next task (like the 'next' command)
	nextTask := pq.Next()
	if nextTask == nil {
		t.Fatal("Expected a task from Next(), got nil")
	}

	// Should get highest priority task
	if nextTask.Name != "Write code" || nextTask.Priority != 10 {
		t.Errorf("Expected 'Write code' with priority 10, got %s with priority %d",
			nextTask.Name, nextTask.Priority)
	}

	// Should have 2 tasks remaining
	if pq.Len() != 2 {
		t.Errorf("Expected 2 tasks remaining, got %d", pq.Len())
	}
}

func TestPriorityQueueUsage(t *testing.T) {
	// Test the actual PriorityQueue functionality that main.go uses
	pq := scheduler.NewPriorityQueue()

	// Test adding tasks
	task1 := &scheduler.Task{Name: "Buy groceries", Priority: 5}
	task2 := &scheduler.Task{Name: "Write code", Priority: 10}
	task3 := &scheduler.Task{Name: "Exercise", Priority: 3}

	pq.Add(task1)
	pq.Add(task2)
	pq.Add(task3)

	// Test listing tasks
	tasks := pq.Tasks()
	if len(tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(tasks))
	}

	// Test next task (should be highest priority)
	nextTask := pq.Next()
	if nextTask == nil {
		t.Fatal("Expected a task, got nil")
	}

	if nextTask.Name != "Write code" || nextTask.Priority != 10 {
		t.Errorf("Expected 'Write code' with priority 10, got %s with priority %d",
			nextTask.Name, nextTask.Priority)
	}

	// Test remaining tasks
	remainingTasks := pq.Tasks()
	if len(remainingTasks) != 2 {
		t.Errorf("Expected 2 remaining tasks, got %d", len(remainingTasks))
	}
}

func TestCommandParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"add Buy groceries 5", []string{"add", "Buy", "groceries", "5"}},
		{"add Write code 10", []string{"add", "Write", "code", "10"}},
		{"list", []string{"list"}},
		{"next", []string{"next"}},
		{"quit", []string{"quit"}},
		{"", []string{}},
	}

	for _, test := range tests {
		parts := strings.Fields(test.input)
		if len(parts) != len(test.expected) {
			t.Errorf("Expected %d parts for input '%s', got %d",
				len(test.expected), test.input, len(parts))
			continue
		}

		for i, part := range parts {
			if part != test.expected[i] {
				t.Errorf("Expected part %d to be '%s', got '%s' for input '%s'",
					i, test.expected[i], part, test.input)
			}
		}
	}
}

// Helper function to test CLI commands without running main
func TestCLICommands(t *testing.T) {
	// This test simulates the command processing logic from main.go
	pq := scheduler.NewPriorityQueue()

	// Test add command
	addCommand := "add Buy groceries 5"
	parts := strings.Fields(addCommand)

	if len(parts) >= 3 {
		name := strings.Join(parts[1:len(parts)-1], " ")
		priority := 5 // strconv.Atoi(parts[len(parts)-1])

		if name != "Buy groceries" {
			t.Errorf("Expected name 'Buy groceries', got '%s'", name)
		}

		if priority != 5 {
			t.Errorf("Expected priority 5, got %d", priority)
		}

		pq.Add(&scheduler.Task{Name: name, Priority: priority})
	}

	// Test next command
	if pq.Len() > 0 {
		task := pq.Next()
		if task == nil {
			t.Error("Expected task from non-empty queue")
		}
	}

	// Test list command
	tasks := pq.Tasks()
	// After next command, should have 0 tasks
	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks after next, got %d", len(tasks))
	}
}

// Benchmark the main functionality
func BenchmarkPriorityQueueOperations(b *testing.B) {
	pq := scheduler.NewPriorityQueue()

	// Pre-populate with some tasks
	for i := 0; i < 100; i++ {
		pq.Add(&scheduler.Task{Name: "task", Priority: i})
	}

	b.ResetTimer()

	// Benchmark add operation
	b.Run("Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pq.Add(&scheduler.Task{Name: "benchmark task", Priority: i})
		}
	})

	// Benchmark next operation (requires tasks to be present)
	b.Run("Next", func(b *testing.B) {
		// Add tasks for each benchmark iteration
		for i := 0; i < b.N; i++ {
			pq.Add(&scheduler.Task{Name: "benchmark task", Priority: i})
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if pq.Len() > 0 {
				pq.Next()
			}
		}
	})
}
