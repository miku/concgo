package main

import (
	"log"
	"strings"
	"sync"
)

// Exercise: A worker pool example.
//
//

func work(queue chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range queue {
		log.Println(strings.ToUpper(item))
	}
}

func main() {

}
