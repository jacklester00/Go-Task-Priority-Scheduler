package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"priority_queue_scheduler/scheduler"
)

// main runs the interactive task scheduler CLI.
func main() {
	pq := scheduler.NewPriorityQueue()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Task Scheduler — type 'help' for commands")

	// Main command loop
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue // Skip empty input
		}

		cmd := parts[0]
		switch cmd {
		case "add":
			// Parse: add <name> <priority>
			if len(parts) < 3 {
				fmt.Println("Usage: add <name> <priority>")
				continue
			}
			// Join all parts except last as task name (supports multi-word names)
			name := strings.Join(parts[1:len(parts)-1], " ")
			priority, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				fmt.Printf("Error: '%s' is not a valid priority number\n", parts[len(parts)-1])
				continue
			}
			pq.Add(&scheduler.Task{Name: name, Priority: priority})
			fmt.Println("Added task:", name)

		case "next":
			// Get and remove the highest priority task
			task := pq.Next()
			if task == nil {
				fmt.Println("No tasks.")
			} else {
				fmt.Printf("Next task: %s (priority %d)\n", task.Name, task.Priority)
			}

		case "list":
			// Display all current tasks
			fmt.Println("Current tasks:")
			for i, t := range pq.Tasks() {
				fmt.Printf("%d. %s (priority %d)\n", i+1, t.Name, t.Priority)
			}

		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  add <name> <priority>  - Add a new task")
			fmt.Println("  next                   - Get the next highest priority task")
			fmt.Println("  list                   - Show all current tasks")
			fmt.Println("  help                   - Show this help message")
			fmt.Println("  quit, exit             - Exit the program")

		case "quit", "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}
