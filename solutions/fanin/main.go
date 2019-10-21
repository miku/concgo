package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Exercise: Worker pool with fan-in. Instead of printing the results in the
// worker, the worker will put the result on an output channel. This output
// channel should just print the data.
//
//

// tasks returns a slice of strings simulating tasks.
func tasks() (result []string) {
	for i := 0; i < 1000; i++ {
		result = append(result, fmt.Sprintf("item-%d", i))
	}
	return
}

// worker has an id, a queue to receive tasks from, WaitGroup for joining.
func worker(id int, queue chan string, out chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range queue {
		result := fmt.Sprintf("[%d] %s", id, strings.ToUpper(task))
		out <- result
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func fanIn(done chan bool, out chan string) {
	for result := range out {
		log.Printf("fanIn: %s", result)
	}
	done <- true
}

func main() {
	queue := make(chan string)
	out := make(chan string)
	done := make(chan bool)

	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go worker(i, queue, out, &wg)
	}

	go fanIn(done, out)

	for _, task := range tasks() {
		queue <- task
	}
	close(queue)
	wg.Wait()
	close(out)
	<-done
}
