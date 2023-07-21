package main

import (
	"fmt"
	"log"
	"time"
)

// function to calculate fibonacci number using recursion & uint64
func fibonacci(n int) uint64 {
	// if n is less than 2 return n
	if n <= 1 {
		return uint64(n)
	}
	// return the sum of the previous two fibonacci numbers
	return fibonacci(n-1) + fibonacci(n-2)
}

// function to calculate fibonacci number using dynamic programming & uint64
func fibonacciDP(n int) uint64 {
	// create a slice of uint64 with length n+1
	f := make([]uint64, n+1)
	// set the first two values of the slice to 0 and 1
	f[0] = 0
	f[1] = 1
	// loop through the slice starting at the third value
	for i := 2; i <= n; i++ {
		// set the value of the current index to the sum of the previous two values
		f[i] = f[i-1] + f[i-2]
	}
	// return the last value in the slice
	return f[n]
}

// function to calculate the amount of time it took to run a fuction
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// main function
func main() {
	// seperately run both functions and track the amount of time for each
	defer timeTrack(time.Now(), "fibonacci")
	fmt.Println(fibonacci(40000))
	defer timeTrack(time.Now(), "fibonacciDP")
	fmt.Println(fibonacciDP(40000))

}
