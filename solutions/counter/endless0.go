package main

import (
	"fmt"
)

//
func main() {
	var c int
	for i := 0; i < 100; i++ {
		go func() {
			var k int
			for {
				k = k + i*2
				k--
			}

		}()
	}
	fmt.Println(c)
}
