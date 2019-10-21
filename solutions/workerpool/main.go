package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

// Exercise: A worker pool example. Some data and a basic worker is already
// there.
//
// (1) Complete the main function, setup a number of workers, the queue.
// (2) Iterate over the data and put the strings of the queue.

// data returns a slice of strings simulating tasks.
func data() (result []string) {
	for i := 0; i < 1000; i++ {
		result = append(result, fmt.Sprintf("item-%d", i))
	}
	return
}

// worker has an id, a queue to receive tasks from, WaitGroup for joining.
func worker(id int, queue chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range queue {
		log.Printf("[%d] %s", id, strings.ToUpper(item))
	}
}

func main() {
	// TODO.
}
