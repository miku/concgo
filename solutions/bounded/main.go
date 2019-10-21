package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func work(id int) {
	log.Printf("[%d] starting work", id)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	log.Printf("[%d] done work", id)
}

func main() {
	N := 100
	sem := make(chan struct{}, 4)

	var wg sync.WaitGroup
	wg.Add(N)

	for i := 0; i < N; i++ {
		i := i
		sem <- struct{}{}
		go func() {
			work(i)
			wg.Done()
			<-sem
		}()
		if i%10 == 0 {
			log.Printf("%d goroutines running", runtime.NumGoroutine())
		}
	}

	wg.Wait()
	log.Printf("%d goroutines running", runtime.NumGoroutine())
}
