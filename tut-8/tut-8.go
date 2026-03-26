package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Simple function to simulate an I/O task (like DB/API call)
func ioTask(id int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("IO Task Done:", id)
}

// Demonstrates race condition (unsafe)
func unsafeWrite(res *[]int, id int) {
	*res = append(*res, id) // multiple goroutines writing → unsafe
}

// Safe write using Mutex
var mu sync.Mutex

func safeWrite(res *[]int, id int) {
	mu.Lock() // only one goroutine enters critical section
	*res = append(*res, id)
	mu.Unlock()
}

// RWMutex example
var rwmu sync.RWMutex

func read(res *[]int) {
	rwmu.RLock() // multiple reads allowed
	fmt.Println("Read:", *res)
	rwmu.RUnlock()
}

func write(res *[]int, val int) {
	rwmu.Lock() // exclusive write
	*res = append(*res, val)
	rwmu.Unlock()
}

// CPU heavy task
func cpuTask() {
	defer wg.Done()
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
}

// Sequential execution (no goroutines)
func sequential() {
	start := time.Now()

	for i := 0; i < 5; i++ {
		ioTask(i)
	}

	fmt.Println("Sequential Time:", time.Since(start))
}

// Concurrent execution using goroutines
func concurrent() {
	start := time.Now()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ioTask(id)
		}(i)
	}

	wg.Wait() // wait for all goroutines
	fmt.Println("Concurrent Time:", time.Since(start))
}

func main() {

	// 1. Sequential vs Concurrent
	sequential()
	concurrent()

	// 2. Unsafe shared memory (DON'T DO THIS)
	res1 := []int{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			unsafeWrite(&res1, id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Unsafe Result:", res1) // may be corrupted

	// 3. Safe shared memory using Mutex
	res2 := []int{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			safeWrite(&res2, id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Safe Result:", res2)

	// 4. RWMutex (multiple reads, exclusive write)
	res3 := []int{}

	for i := 0; i < 3; i++ {
		write(&res3, i)
	}

	for i := 0; i < 3; i++ {
		go read(&res3)
	}

	time.Sleep(1 * time.Second)

	// 5. CPU-bound goroutines (limited by cores)
	start := time.Now()

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go cpuTask()
	}

	wg.Wait()
	fmt.Println("CPU Tasks Time:", time.Since(start))

}
