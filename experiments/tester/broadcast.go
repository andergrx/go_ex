package main

import (
	"log"
	"sync"
	"time"
)

type Comms struct {
	start chan struct{}
	next  chan struct{}
	end   chan struct{}
	wg    sync.WaitGroup
}

func main() {

	controller := &Comms{
		start: make(chan struct{}),
		next:  make(chan struct{}),
		end:   make(chan struct{}),
	}

	for i := 1; i <= 3; i++ {
		controller.wg.Add(2)
		go drones(i, controller)
		go betterDrones(i, controller)
	}

	time.Sleep(1 * time.Second)
	close(controller.start)
	time.Sleep(1 * time.Second)
	close(controller.next)
	time.Sleep(1 * time.Second)
	close(controller.end)

	log.Println("Main waiting...")
	controller.wg.Wait()
	log.Println("Ending main.")

}

const (
	Start = iota
	Operate
	End
)

func drones(id int, c *Comms) {
	defer c.wg.Done()

	state := Start
	for {
		select {
		case <-c.start:
			if state == Start {
				log.Printf("Drone %d starting...", id)
				state = Operate
			}
		case <-c.next:
			if state == Operate {
				log.Printf("Drone %d doing next operation...", id)
				state = End
			}
		case <-c.end:
			log.Printf("Drone %d ending...", id)
			return
		}
	}

}

func betterDrones(id int, c *Comms) {
	defer c.wg.Done()

	<-c.start
	log.Printf("Better Drone %d starting...", id)
	<-c.next
	log.Printf("Better Drone %d operating...", id)
	<-c.end
	log.Printf("Better Drone %d ending...", id)
}
