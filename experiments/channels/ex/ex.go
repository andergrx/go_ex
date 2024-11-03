package main

import (
	"log"
	"sync"
)

const Jobs = 5

type comm struct {
	ch     chan int
	exCh   chan int
	exDone chan struct{}
	count  int
	wg     sync.WaitGroup
	mut    sync.Mutex
}

func (c *comm) Inc() {
	c.mut.Lock()
	c.count++
	c.mut.Unlock()
}

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	com := &comm{
		make(chan int, Jobs),
		make(chan int, Jobs),
		make(chan struct{}),
		0,
		sync.WaitGroup{},
		sync.Mutex{},
	}

	com.wg.Add(1)
	go tx(com)

	for i := 0; i < Jobs; i++ {
		com.wg.Add(1)
		go rx(i+1, com)
	}

	com.wg.Add(1)
	go ExtendTx(com)

	for i := 0; i < Jobs; i++ {
		com.wg.Add(1)
		go ExtendRx(i+1, com)
	}

	com.wg.Add(1)
	go Cleaner(com)

	com.wg.Wait()
	log.Println("Main End:", com.count)

}

func tx(com *comm) {
	defer com.wg.Done()
	for i := 0; i < 30; i++ {
		log.Println("Tx:", i)
		com.ch <- i
	}
	close(com.ch)
}

func rx(n int, com *comm) {
	defer com.wg.Done()
	for {
		c, ok := <-com.ch
		if !ok {
			return
		}
		com.Inc()
		log.Printf("Rx%d: %d\n", n, c)
	}
}

func ExtendTx(com *comm) {
	defer com.wg.Done()
	for i := 30; i < 40; i++ {
		log.Println("ExtendTx:", i)
		com.exCh <- i
	}
	close(com.exCh)
	com.exDone <- struct{}{}
}

func ExtendRx(n int, com *comm) {
	defer com.wg.Done()
	for {
		select {
		case c, ok := <-com.exCh:
			if !ok {
				return
			}
			com.Inc()
			log.Printf("ExtendRx%d: %d\n", n, c)

		case <-com.exDone:
			log.Printf("ExtendRx%d Done\n", n)
		}
	}
}

func Cleaner(com *comm) {
	defer com.wg.Done()

	<-com.exDone

	for c := range com.exCh {
		com.Inc()
		log.Println("Cleaning up:", c)
	}
}
