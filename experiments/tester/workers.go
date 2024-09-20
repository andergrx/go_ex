// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

const Workers = 3
const JobCount = 20

type Comms struct {
	ch   chan string
	done chan struct{}
}

func main() {

	fmt.Println("Worker Main")

	ch := make(chan string, Workers)
	senderDone := make(chan struct{}, 1)
	workersDone := make(chan struct{}, 1)
	var mutex sync.Mutex

	go sender(ch, senderDone, workersDone)

	procCount := 0
	for i := 0; i < Workers; i++ {
		go func() {
			//defer wg.Done()
			fmt.Printf("Listener%v...\n", i+1)
			for {
				select {
				case info := <-ch:
					mutex.Lock()
					procCount++
					if procCount == JobCount {
						close(workersDone)
					}
					mutex.Unlock()
					fmt.Printf("Worker%v Info: %v\n", i+1, info)

				case <-senderDone:
					fmt.Printf("Worker%v Done!\n", i+1)
					return
				}
			}
		}()
	}

	<-senderDone
	fmt.Println("Then end of Main:", procCount)

}

func sender(ch chan string, sendDone chan struct{}, workersDone chan struct{}) {
	defer close(sendDone)

	fmt.Println("Sender operating...")
	for i := 0; i < JobCount; i++ {
		ch <- fmt.Sprintf("Sending%v", i+1)
	}
	<-workersDone

}
