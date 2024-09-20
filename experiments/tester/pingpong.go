package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const MaxHits = 50

type Comms struct {
	ping chan struct{}
	pong chan struct{}
	ref  chan struct{}
	done chan struct{}
	mut  sync.Mutex

	hits int
}

func main() {

	fmt.Println("Ping Pong...")

	comms := &Comms{
		ping: make(chan struct{}, 1),
		pong: make(chan struct{}, 1),
		ref:  make(chan struct{}, 1),
		done: make(chan struct{}, 1),
		hits: 0,
	}

	go ping(comms)
	go pong(comms)
	go referee(comms)

	<-comms.done
	fmt.Println("Game Over!")

}

func ping(c *Comms) {

	for {
		select {
		case <-c.pong:
			hitBall(c)
			log.Printf("Pinging... %d", c.hits)
			time.Sleep(1 * time.Second)
			c.ref <- struct{}{}
			c.ping <- struct{}{}

		case <-c.done:
			log.Println("Ping Out!")
			return

		default:

		}
	}
}

func pong(c *Comms) {
	c.pong <- struct{}{}
	for {
		select {
		case <-c.ping:
			hitBall(c)
			log.Printf("Ponging... %d", c.hits)
			time.Sleep(1 * time.Second)
			c.ref <- struct{}{}
			c.pong <- struct{}{}

		case <-c.done:
			log.Println("Pong Out!")
			return
		}
	}
}

func hitBall(c *Comms) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.hits++
}

func referee(c *Comms) {
	for {
		select {
		case <-c.ref:
			c.mut.Lock()
			if c.hits >= MaxHits {
				c.mut.Unlock()
				log.Println("Reffing...")
				close(c.done)
				return
			}
			c.mut.Unlock()
		}
	}
}
