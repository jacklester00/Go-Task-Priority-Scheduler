# Go Task Priority Scheduler

A simple task priority scheduler built in Go using a max-heap data structure.

## Overview

This project implements a task scheduling system with priority-based execution. Tasks are stored in a priority queue (max-heap) where higher priority numbers indicate more important tasks that should be executed first.

## Features

- **Priority-based task scheduling**: Tasks with higher priority numbers are executed first
- **Interactive CLI**: Simple command-line interface for managing tasks
- **Heap-based implementation**: Efficient O(log n) insertion and O(log n) removal
- **Real-time task listing**: View all current tasks in priority order
- **Comprehensive testing**: Full unit test coverage with benchmarks
- **CI/CD pipeline**: Automated testing with GitHub Actions

## Getting Started

### Prerequisites

- Go 1.21 or later

### Installation

1. Clone the repository:
```bash
git clone https://github.com/jacklester00/Go-Task-Priority-Scheduler.git
cd Go-Task-Priority-Scheduler
```

2. Build the project:
```bash
go build -o scheduler .
```

Or run directly:
```bash
go run .
```

## Usage

Start the scheduler:
```bash
./scheduler
# or
go run .
```

### Available Commands

- `add <name> <priority>` - Add a new task (higher numbers = higher priority)
- `next` - Get and remove the next highest priority task
- `list` - Show all current tasks
- `help` - Show available commands
- `quit` or `exit` - Exit the program

### Example Session

```
Task Scheduler — type 'help' for commands
> add Buy groceries 5
Added task: Buy groceries
> add Write code 10
Added task: Write code
> add Exercise 3
Added task: Exercise
> list
Current tasks:
1. Write code (priority 10)
2. Buy groceries (priority 5)
3. Exercise (priority 3)
> next
Next task: Write code (priority 10)
> list
Current tasks:
1. Buy groceries (priority 5)
2. Exercise (priority 3)
> quit
Goodbye!
```

## Testing

The project includes comprehensive unit tests and benchmarks:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run benchmarks
go test -bench=. ./...

# Run tests with coverage
go test -cover ./...
```

### Test Coverage

- **Unit tests**: Complete coverage of all scheduler components
- **Integration tests**: End-to-end testing of CLI functionality
- **Benchmark tests**: Performance validation of heap operations
- **Edge case testing**: Empty queues, single tasks, priority ordering

## CI/CD

Automated testing with GitHub Actions:

- **Tests**: Run on every push and pull request
- **Build verification**: Ensures project compiles successfully
- **Benchmarks**: Performance validation

![CI Status](https://github.com/jacklester00/Go-Task-Priority-Scheduler/workflows/Tests/badge.svg)

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests to ensure everything works (`go test ./...`)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Submit a pull request

### Development Setup

```bash
# Clone the repository
git clone https://github.com/jacklester00/Go-Task-Priority-Scheduler.git
cd Go-Task-Priority-Scheduler

# Install dependencies
go mod download

# Run tests
go test ./...

# Build the project
go build .
```

## License

This project is licensed under the MIT License.
