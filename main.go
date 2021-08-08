// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func main() {
	// Create initial bytes array with 1KiB
	list := make([]byte, 1024)
	dur := time.Second * 2
	for {
		// The first loop increase memory allocation until become much slower
		start := time.Now()
		for i := 0; i < 4; i++ {
			flip(list)
			reverse(list)
		}
		list = grow(list)
		fmt.Printf("Growth to %v\n", len(list))
		newDur := time.Since(start)
		if newDur/dur > 4 {
			// Most probably, we reached swap
			break
		}
	}
	fmt.Printf("Looping, size: %v\n", len(list))
	for {
		// The second loop works forever
		start := time.Now()
		flip(list)
		reverse(list)
		newDur := time.Since(start)
		fmt.Printf("duration: %v\n", newDur)
		time.Sleep(time.Second)
	}
}

func grow(list []byte) []byte {
	list = append(list, list...)
	return list
}

func flip(list []byte) {
	for i, l := range list {
		list[i] = ^l
	}
}

func reverse(list []byte) {
	for i := 0; i < len(list)/2; i++ {
		list[i], list[len(list)-i-1] = list[len(list)-i-1], list[i]
	}
}
