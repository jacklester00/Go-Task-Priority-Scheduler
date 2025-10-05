package scheduler

import (
	"testing"
)

func TestTask_Initialization(t *testing.T) {
	task := &Task{
		Name:     "test task",
		Priority: 5,
		Index:    0,
	}

	if task.Name != "test task" {
		t.Errorf("Expected name 'test task', got '%s'", task.Name)
	}

	if task.Priority != 5 {
		t.Errorf("Expected priority 5, got %d", task.Priority)
	}

	if task.Index != 0 {
		t.Errorf("Expected index 0, got %d", task.Index)
	}
}

func TestTask_ZeroValues(t *testing.T) {
	task := &Task{}

	if task.Name != "" {
		t.Errorf("Expected empty name, got '%s'", task.Name)
	}

	if task.Priority != 0 {
		t.Errorf("Expected priority 0, got %d", task.Priority)
	}

	if task.Index != 0 {
		t.Errorf("Expected index 0, got %d", task.Index)
	}
}

func TestTask_FieldModification(t *testing.T) {
	task := &Task{Name: "original", Priority: 1, Index: 0}

	// Test modifying fields
	task.Name = "modified"
	task.Priority = 10
	task.Index = 5

	if task.Name != "modified" {
		t.Errorf("Expected name 'modified', got '%s'", task.Name)
	}

	if task.Priority != 10 {
		t.Errorf("Expected priority 10, got %d", task.Priority)
	}

	if task.Index != 5 {
		t.Errorf("Expected index 5, got %d", task.Index)
	}
}
