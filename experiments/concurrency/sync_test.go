package main

import (
	"fmt"
	"testing"
)

func BenchmarkSyncTest(b *testing.B) {
	// Setup code (if needed)

	for i := 0; i < b.N; i++ {
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

				//fmt.Printf("Op%v...", i)
				for iFunc := 0; iFunc < Append; iFunc++ {
					words = append(words, fmt.Sprintf("Op%v", i))
				}
			}()
		}

		syn.wg.Wait()
	}
}
