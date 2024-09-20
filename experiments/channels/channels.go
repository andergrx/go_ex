// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var words = []string{"One", "Two", "Dude", "What"}

func main() {
	ch := make(chan string, 1)
	done := make(chan struct{}, 1)

	go sender(ch, done)

	fmt.Println("Main Time")

	//finished := false
	for {
		select {
		case info := <-ch:
			fmt.Println("Info:", info)
		case <-done:
			fmt.Println("Done!")
			return
		}
	}

	//fmt.Println("The End")

}

func sender(ch chan string, done chan struct{}) {

	defer close(done)
	fmt.Println("Sender...")
	for _, word := range words {
		// if i == 3 {
		// 	close(done)
		// 	return
		// }
		fmt.Println("Sending:", word)
		ch <- word
	}
}
