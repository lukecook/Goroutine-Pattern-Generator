# Goroutine Pattern Generator 🎨

A visual demonstration of Go's concurrent execution and scheduler behavior using goroutines and channels.

## Background

I've been diving into concurrent programming in Go, working through examples and exercises from *Head First Go* (Chapter 13). One exercise had me running two goroutines that print letters ('a' and 'b') to demonstrate how the Go scheduler interleaves execution.

It worked great, but I wanted something more visually interesting! So I extended it to use four goroutines with colorful emoji symbols instead of plain letters. The result makes the scheduler's behavior much easier to see.

## What It Does

This program launches 4 concurrent goroutines, each printing a different colored circle emoji 30 times:
- 🔴 Red circles
- 🔵 Blue circles  
- 🟢 Green circles
- 🟡 Yellow circles

All four compete for stdout, and the Go scheduler decides which one gets CPU time at any moment. **Every run produces a completely different pattern**.

## Key Concepts

### Concurrent vs Sequential
- **Sequential**: One task completes before the next begins. Predictable but slower.
- **Concurrent** (this program): Multiple tasks make progress during overlapping time periods, interleaved by the scheduler or running truly parallel on multiple cores.

### Race Condition on Shared Resource
Four goroutines (multiple tasks) compete for stdout (one shared resource). Since they share the output, their prints get mixed together based on scheduler timing.
This is a classic **race condition** where multiple goroutines race to access the same resource.

### Channel Synchronization
Channels provide a way to communicate between goroutines and synchronize their execution.

- When you read from an empty channel (`<-done`), the goroutine **blocks** (waits) until another goroutine sends a value into that channel.
- By reading from the channel four times, `main()` waits for **all four** goroutines to complete.
- Without this synchronization, `main()` could exit immediately and result in an incomplete output.

## Running the Program
```bash
go run goroutine-patterns.go
```
Run it multiple times to see different patterns!

## What I Learned

- How goroutines enable lightweight concurrency
- Using channels to block and synchronize concurrent tasks
- Why concurrent programs produce unpredictable results
- How race conditions occur when tasks compete for shared resources
- The difference between concurrent and parallel execution
