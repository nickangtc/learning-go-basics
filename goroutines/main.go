package main

import (
	"fmt"
	"time"
)

func greet(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(500 * time.Millisecond)
	}
}

// illustrates `go` keyword being used for multi-threading, aka goroutines
func main() {
	// 🐢 slowest due to single thread + sleep
	// greet("Nick Ang")
	// greet("Yondu BIllars")

	// ⚡️ multi-threading with `go`
	go greet("Nick Ang")
	greet("Yondu BIllars")

	// 🚨 will not work because main process exits before goroutines finish executing
	// unless you put infinite loop to ensure main process stays running (last line)
	// go greet("Nick Ang")
	// go greet("Yondu BIllars")
	// for {
	// }
}
