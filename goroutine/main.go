package main

import (
	"fmt"
	"runtime"
)

func main() {
	names := []string{"Eric", "Harry", "Robbert", "Jim", "Mark"}
	for _, name := range names {
		fmt.Printf("Hello!, %s\n", name)
		go func(who string) {
			fmt.Printf("Hello, %s\n", who)
		}(name)
		runtime.Gosched()
	}
}
