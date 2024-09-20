// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

const Workers = 3
const SendMax = 100

type Comms struct {
	ch   chan string
	done chan struct{}
	wg   sync.WaitGroup
	mtx  sync.Mutex
	acc  int
}

func main() {

	fmt.Println("Worker Main")

	com := &Comms{
		ch:   make(chan string, Workers),
		done: make(chan struct{}, 1),
	}

	com.wg.Add(Workers)
	for i := 1; i <= Workers; i++ {
		go func() {
			fmt.Printf("Worker%v Listening...\n", i)
			defer com.wg.Done()

			wAccum := 0
			for {
				select {
				case info, ok := <-com.ch:
					if !ok {
						fmt.Printf("Worker%v Done: %v\n", i, wAccum)
						return
					}
					wAccum++
					com.acc++
					fmt.Printf("Worker%v Info: %v\n", i, info)

				default:
					// time.Sleep(1 * time.Second)
					// fmt.Printf("%v: Working...\n", i)
					// case <-com.done:
					// 	fmt.Printf("Worker%v Done!\n", i)
					// 	com.wg.Done()
					// 	return
				}
			}
		}()
	}

	//com.wg.Add(1)
	go sender(com)

	com.wg.Wait()
	fmt.Println("Then end of Main", com.acc)

}

func sender(c *Comms) {
	//defer c.wg.Done()
	fmt.Println("Sender operating...")
	for i := 0; i < SendMax; i++ {
		fmt.Println("Sending Action", i+1)
		c.ch <- fmt.Sprintf("Sending%v", i+1)
	}
	close(c.ch)
}

func processChannel(c *Comms, num int, info string) {
	c.acc++
	fmt.Printf("Worker%v Info: %v\n", num, info)
}
