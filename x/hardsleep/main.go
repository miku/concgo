package main

import (
	"fmt"
	"time"
)

func main() {
	var data int
	go func() {
		data++
	}()
	time.Sleep(5 * time.Microsecond)
	if data == 0 {
		fmt.Printf("value = %d\n", data)
	}
}
