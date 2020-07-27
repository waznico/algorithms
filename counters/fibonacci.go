package counters

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	val := 0
	return func() int {
		val = val + (val + 1)
		return val
	}
}

// GetFibonacciChain returns the given amount of fibonacci numbers
func GetFibonacciChain(i int) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
