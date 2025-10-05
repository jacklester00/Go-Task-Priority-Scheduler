package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"

    "priority-queue-scheduler/scheduler"
)

func main() {
    pq := scheduler.NewPriorityQueue()
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("Task Scheduler — type 'help' for commands")
    for {
        fmt.Print("> ")
        scanner.Scan()
        input := scanner.Text()
        parts := strings.Fields(input)
        if len(parts) == 0 { continue }

        cmd := parts[0]
        switch cmd {
        case "add":
            if len(parts) < 3 {
                fmt.Println("Usage: add <name> <priority>")
                continue
            }
            name := strings.Join(parts[1:len(parts)-1], " ")
            priority, _ := strconv.Atoi(parts[len(parts)-1])
            pq.Add(&scheduler.Task{Name: name, Priority: priority})
            fmt.Println("Added task:", name)
        case "next":
            task := pq.Next()
            if task == nil { fmt.Println("No tasks.") } else {
                fmt.Printf("Next task: %s (priority %d)\n", task.Name, task.Priority)
            }
        case "list":
            fmt.Println("Current tasks:")
            for i, t := range pq.Tasks() {
                fmt.Printf("%d. %s (priority %d)\n", i+1, t.Name, t.Priority)
            }
        case "quit", "exit":
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Commands: add, next, list, quit")
        }
    }
}
