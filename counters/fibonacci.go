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

// fibonacciChan takes an amount and writes the result into channel c
func fibonacciChan(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// GetFibonacciChain returns the given amount of fibonacci numbers
func GetFibonacciChain(i int) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
