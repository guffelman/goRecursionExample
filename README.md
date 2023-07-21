# goRecursionExample

## Introduction to Recursion

Recursion is a programming technique where a function calls itself to solve a problem. Let's explain recursion in the context of the repo's Fibonacci code.

The Fibonacci sequence is a series of numbers where each number is the sum of the two preceding ones, usually starting with 0 and 1. The Fibonacci functions implemented here calculate the nth Fibonacci number using two different approaches: recursion and dynamic programming.

1. Recursive Fibonacci Function:
```go
func fibonacci(n int) uint64 {
	if n <= 1 {
		return uint64(n)
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
```

**Explanation of the Recursive Function:**
- The `fibonacci` function takes an integer `n` as input and returns the nth Fibonacci number as a `uint64`.
- It starts with a base case: if `n` is less than or equal to 1, it returns `n`. The base case is essential in recursion to stop the function from calling itself indefinitely and avoid infinite recursion.
- When `n` is greater than 1, the function calls itself twice with `n-1` and `n-2` as arguments and adds the results of these recursive calls. This is the recursive step, where the function breaks down the problem into smaller subproblems and solves them.

**Example: Calculating the 5th Fibonacci number (n = 5):**
1. `fibonacci(5)` is called. It's not a base case, so we go to the next step.
2. `fibonacci(5)` calls `fibonacci(4) + fibonacci(3)`.
3. `fibonacci(4)` calls `fibonacci(3) + fibonacci(2)`, and `fibonacci(3)` calls `fibonacci(2) + fibonacci(1)`.
4. The process continues, and eventually, the base cases are reached: `fibonacci(2)` returns 1, `fibonacci(1)` returns 1, and `fibonacci(0)` returns 0.
5. Now, the calculations are done in reverse order:
   - `fibonacci(1) + fibonacci(0)` becomes `1 + 0`, which is 1.
   - `fibonacci(2) + fibonacci(1)` becomes `1 + 1`, which is 2.
   - `fibonacci(3) + fibonacci(2)` becomes `2 + 1`, which is 3.
   - `fibonacci(4) + fibonacci(3)` becomes `3 + 2`, which is 5.
   - Finally, `fibonacci(5)` returns `fibonacci(4) + fibonacci(3)` as 5.

Therefore, `fibonacci(5)` returns 5, which is the 5th Fibonacci number.

Recursion is a powerful concept, but it can lead to performance issues for large values of `n`, as the function will repeatedly compute the same values multiple times. That's where dynamic programming comes in handy, as it optimizes the computation by storing the previously calculated Fibonacci numbers in an array (slice) and reusing them instead of recomputing.

## Exponential Processing Requirement

The processing requirement becomes exponential in recursive algorithms like the one used to calculate the Fibonacci sequence due to redundant and repeated computations. Let's explore why this happens.

In the recursive Fibonacci function:

```go
func fibonacci(n int) uint64 {
	if n <= 1 {
		return uint64(n)
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
```

When you call `fibonacci(n)` for some value of `n`, the function will branch out and make two more recursive calls: `fibonacci(n-1)` and `fibonacci(n-2)`. Each of these calls will again branch out, leading to further recursive calls, and so on, until the base cases are reached. This creates a tree-like structure of function calls, known as a recursion tree.

For example, if you call `fibonacci(5)`, the recursion tree will look like this:

```
                fibonacci(5)
               /           \
      fibonacci(4)       fibonacci(3)
       /      \           /       \
fibonacci(3)  fibonacci(2)  fibonacci(2)  fibonacci(1)
 /     \
fibonacci(2)  fibonacci(1)
```

As you can see, some of the function calls are repeated multiple times in different branches of the recursion tree. For instance, `fibonacci(3)` is called twice, and `fibonacci(2)` is called three times. The higher the value of `n`, the more repetitions and redundant computations occur.

Now, let's understand the processing requirement for `fibonacci(n)` in terms of time complexity:

- Each call to `fibonacci(n)` branches into two more calls, leading to 2 recursive calls for each level of the recursion tree.
- The depth of the recursion tree is `n` because in each call, the value of `n` decreases by 1 until it reaches the base cases (`n <= 1`).
- As a result, the number of nodes (function calls) in the recursion tree grows exponentially with `n`.

The time complexity of the recursive Fibonacci function can be expressed as O(2^n). This is because, in the worst case, the function makes two recursive calls for each level of the recursion tree, and there are approximately 2^n nodes in the tree. So, the number of operations required to compute `fibonacci(n)` grows exponentially with the input `n`.

This exponential growth becomes a problem for larger values of `n`, leading to a significant increase in processing time and making the recursive approach inefficient. This is where dynamic programming or memoization comes into play, as it stores intermediate results and avoids redundant computations, effectively reducing the time complexity to O(n) in the case of the Fibonacci sequence.

## Dynamic Programming and Memoization

Dynamic programming is a technique that optimizes recursive algorithms by storing intermediate results in a data structure, usually an array, so that the same calculations don't need to be repeated multiple times. This approach avoids redundant computations and improves the time complexity of the algorithm.

In the case of the Fibonacci sequence, dynamic programming can be applied using memoization. Memoization involves caching the results of expensive function calls and reusing them when the same inputs occur again. By storing the previously calculated Fibonacci numbers in a slice, we can access and reuse them when needed, rather than recomputing them from scratch.

Let's take a look at the dynamic programming implementation of the Fibonacci function:

```go
func fibonacciDP(n int) uint64 {
	f := make([]uint64, n+1) // create a slice of uint64 with length n+1
	f[0] = 0                  // set the first two values of the slice to 0 and 1
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2] // calculate and store the Fibonacci number at index i
	}

	return f[n] // return the nth Fibonacci number
}
```

By using the dynamic programming approach with memoization, we avoid redundant computations and significantly reduce the time complexity to O(n). This makes the Fibonacci function much more efficient for larger values of `n` compared to the recursive approach.



In conclusion, dynamic programming, along with memoization, is a powerful technique to optimize recursive algorithms and avoid the exponential processing requirement that arises from redundant computations.
