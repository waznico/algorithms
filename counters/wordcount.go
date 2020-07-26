package counters

import (
	"strings"

	"golang.org/x/tour/wc"
)

//WordCount counts all upcoming words within the string
func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	wcMap := make(map[string]int)
	for i := 0; i < len(fields); i++ {
		wcMap[fields[i]]++
	}
	return wcMap
}

//ExampleWordCountRun shows an example output with a test set
func ExampleWordCountRun() {
	wc.Test(WordCount)
}
