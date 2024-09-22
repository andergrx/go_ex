package main

import (
	"log"
	"sync"
)

const NumGo = 10

func main() {

	ch := make([]chan int, NumGo)
	for i := 0; i < NumGo; i++ {
		ch[i] = make(chan int, 1)
	}
	event := make(chan struct{})
	var wg sync.WaitGroup

	for i := 1; i <= NumGo; i++ {
		wg.Add(1)
		go func(n int, ci chan int) {
			defer wg.Done()
			c := <-ci
			log.Printf("Go%d read: %d\n", n, c)
		}(i, ch[i-1])
	}

	for i := 1; i <= NumGo; i++ {
		wg.Add(1)
		go func(n int, ev <-chan struct{}) {
			defer wg.Done()
			tmp := <-ev
			log.Printf("G2 Go%d got event %v", n, tmp)
		}(i, event)
	}

	for i := 0; i < NumGo; i++ {
		ch[i] <- i + 1
		close(ch[i])
	}
	// Send event with close for all chans to rcv
	close(event)

	wg.Wait()
	log.Println("End main.")
}
