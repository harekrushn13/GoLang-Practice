In Go, the `sync.WaitGroup` is used for synchronizing goroutines. It allows you to wait for a collection of goroutines to complete their execution. It is a powerful tool to manage concurrency in Go programs when you need to ensure that multiple goroutines finish their work before proceeding.

### Key Methods of `sync.WaitGroup`

1. **`Add(int)`**:
    - This method is used to set the number of goroutines the `WaitGroup` should wait for.
    - The argument to `Add` is an integer that specifies the number of goroutines to wait for. You usually call this before starting the goroutines.
    - It can be positive, negative, or zero, but in most cases, it's used with a positive integer to signal the number of goroutines.

2. **`Done()`**:
    - This method should be called by a goroutine when it finishes its work.
    - Each time a goroutine finishes its execution, it calls `Done()` to decrement the `WaitGroup`'s counter by one.
    - The `Done()` method is typically called at the end of the goroutine.

3. **`Wait()`**:
    - This method blocks the calling goroutine until the `WaitGroup` counter is decremented to zero.
    - It is called by the main goroutine (or any other goroutine) that wants to wait for the other goroutines to finish their work.

### Example: Using `sync.WaitGroup`

Let’s go through a simple example that demonstrates how `WaitGroup` works.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when goroutine completes

	fmt.Printf("Worker %d starting\n", id)
	// Simulating some work with a sleep
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Add 3 to the WaitGroup counter (we have 3 workers)
	wg.Add(3)

	// Launch 3 goroutines
	for i := 1; i <= 3; i++ {
		go worker(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Now all workers are done
	fmt.Println("All workers finished.")
}
```

### Step-by-Step Explanation:

1. **Initialize a `sync.WaitGroup`:**
   ```go
   var wg sync.WaitGroup
   ```

   This creates a `WaitGroup` that will be used to wait for all the goroutines to finish.

2. **Add goroutines to the `WaitGroup`:**
   ```go
   wg.Add(3)
   ```

   This tells the `WaitGroup` that we will be waiting for 3 goroutines. You call `Add` with the number of goroutines you plan to launch. Here, we want to wait for 3 workers to finish.

3. **Start the goroutines:**
   ```go
   go worker(i, &wg)
   ```

   We start 3 goroutines in the loop, each calling the `worker` function. The `worker` function takes an `id` and a pointer to the `WaitGroup` (`&wg`).

4. **In the worker goroutine:**
   ```go
   defer wg.Done()
   ```

   The `defer` keyword ensures that the `Done()` method will be called when the goroutine finishes executing. This decreases the `WaitGroup` counter by 1.

5. **Wait for all goroutines to complete:**
   ```go
   wg.Wait()
   ```

   The `Wait()` method blocks the main goroutine until the `WaitGroup` counter reaches zero (i.e., all goroutines have called `Done()`).

6. **Output:**
   The program will print the following output (order may vary because goroutines run concurrently):

   ```
   Worker 1 starting
   Worker 2 starting
   Worker 3 starting
   Worker 1 done
   Worker 2 done
   Worker 3 done
   All workers finished.
   ```

### Detailed Breakdown of Methods:

1. **`Add(delta int)`**:
    - **Purpose**: Sets or modifies the counter of the `WaitGroup`.
    - **Usage**: You call `Add` before starting any goroutines, specifying how many goroutines to wait for.
    - **Can it be negative?**: Yes, you can call `Add` with negative numbers, but it must be done carefully. If you decrement the counter below zero, `Wait()` will block forever.

   **Example:**
   ```go
   var wg sync.WaitGroup
   wg.Add(3)  // Start with a count of 3
   ```

   The counter here is set to 3, meaning `Wait()` will block until 3 calls to `Done()` are made.

2. **`Done()`**:
    - **Purpose**: Decreases the `WaitGroup` counter by 1.
    - **Usage**: Each goroutine that is part of the `WaitGroup` should call `Done()` when it completes its task. This signals that the goroutine has finished its work.
    - **How to ensure it’s called?**: `defer wg.Done()` is commonly used inside goroutines to make sure the `Done()` method is called when the goroutine finishes.

   **Example:**
   ```go
   defer wg.Done()
   ```

3. **`Wait()`**:
    - **Purpose**: Blocks the calling goroutine until the `WaitGroup` counter reaches zero. This is typically called in the main goroutine to wait for all other goroutines to finish.
    - **Usage**: This method is used when you want the program to wait for the completion of all goroutines before continuing.

   **Example:**
   ```go
   wg.Wait()  // Wait until all goroutines have called Done()
   ```

### Edge Cases and Important Considerations:

1. **Calling `Add()` After `Wait()`**:
    - You should **not** call `Add()` after calling `Wait()`. Doing so can lead to a runtime panic or undefined behavior because `Wait()` is designed to block until the `WaitGroup` counter reaches zero. Changing the counter after `Wait()` could cause the program to wait indefinitely or cause a panic if the counter goes negative.

   **Incorrect usage:**
   ```go
   wg.Wait()   // Blocking
   wg.Add(1)   // Adding after Wait() has started is incorrect!
   ```

2. **Negative `Add()`**:
    - **Be careful with negative values** passed to `Add()`. Decrementing the counter beyond zero causes `Wait()` to wait forever (deadlock). You must ensure that the number of `Add()` calls and `Done()` calls match.

   **Potential Deadlock Example**:
   ```go
   var wg sync.WaitGroup
   wg.Add(1) // Add 1
   wg.Done() // Done after goroutine is finished
   wg.Wait() // Wait indefinitely because the count was zero
   ```

3. **Zero `Add()`**:
    - If you call `Add(0)`, `Wait()` will immediately return without blocking, which is useful in some scenarios where you don’t need to wait for any goroutines.

   ```go
   var wg sync.WaitGroup
   wg.Add(0)  // Don't wait for any goroutines
   wg.Wait()  // No blocking occurs
   ```

4. **Using `Done()` Multiple Times**:
    - **Calling `Done()` multiple times** (without an equivalent `Add()` call) will cause a runtime panic, as the counter will go negative.

   **Incorrect usage example**:
   ```go
   wg.Done()  // This will panic if Add() hasn't been called
   ```

### Conclusion:
- `sync.WaitGroup` is a fundamental synchronization tool in Go for managing the execution of multiple goroutines. By using `Add()` to set the number of goroutines, `Done()` to signal when each goroutine finishes, and `Wait()` to block until all goroutines have completed, you can efficiently coordinate concurrent tasks.
- Careful management of the counter (using `Add()` and `Done()` properly) is essential to prevent race conditions, deadlocks, or panics.
