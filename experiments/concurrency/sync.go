package main

import (
	"fmt"
	"sync"
)

var words = []string{"One", "Two", "Dude", "What"}

const Append int = 3

type syncher struct {
	wg  sync.WaitGroup
	mtx sync.Mutex
}

func main() {

	syn := &syncher{}

	syn.wg.Add(1)
	go op1(syn)
	syn.wg.Add(1)
	go op2(syn)

	for i := 3; i < 100; i++ {
		syn.wg.Add(1)
		go func() {
			syn.mtx.Lock()
			defer syn.wg.Done()
			defer syn.mtx.Unlock()

			fmt.Printf("Op%v...", i)
			for iFunc := 0; iFunc < Append; iFunc++ {
				words = append(words, fmt.Sprintf("Op%v", i))
			}
		}()
	}

	syn.wg.Wait()

	fmt.Println("\n\nFinal Words:\n\n", words)
	fmt.Println("\n\nFinal length:", len(words))
	fmt.Println("Main Done!")
}

func op1(s *syncher) {
	s.mtx.Lock()
	defer s.wg.Done()
	defer s.mtx.Unlock()

	//fmt.Println("Op1...")
	for i := 0; i < Append; i++ {
		words = append(words, "Op1")
	}
}

func op2(s *syncher) {
	s.mtx.Lock()
	defer s.wg.Done()
	defer s.mtx.Unlock()

	//fmt.Println("Op2...")
	for i := 0; i < Append; i++ {
		words = append(words, "Op2")
	}
}
