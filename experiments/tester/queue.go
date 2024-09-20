package main

import "fmt"

func main() {

	var qi Queue[int]
	for i := 0; i < 10; i++ {
		qi.push(i)
	}
	fmt.Println(qi)
	qi.dumpQueue()
	fmt.Println(qi)

	qs := new(Queue[string])
	for i := 0; i < 10; i++ {
		qs.push(fmt.Sprintf("Random dumb string %v", i+1))
	}
	fmt.Println(qs)
	qs.dumpQueue()
	fmt.Println(qs)

	qf := newQueueFloat32()
	qf.push(432.2325)
	fmt.Println("qf:", *qf.front(), qf)

	qr := Queue[rune]{t: make([]rune, 0, 20)}
	fmt.Println("empty rune q:", qr)

}

// type Stack[T any] struct {
// 	lifo []T
// }

type Queue[T any] struct {
	t []T
}

func (q *Queue[T]) push(item T) {
	q.t = append(q.t, item)
}

func (q *Queue[T]) front() *T {
	if len(q.t) == 0 {
		return nil
	}

	first := q.t[0]
	q.t = q.t[1:]

	return &first
}

func (q *Queue[T]) dumpQueue() {
	for {
		if f := q.front(); f != nil {
			fmt.Println("front:", *f)
		} else {
			break
		}
	}
}

// Go makes the decision for the developer about stack vs. heap memory
// making this function actually work where the memory of a local
// variable escapes the function
func newQueueFloat32() *Queue[float32] {
	var q Queue[float32]
	return &q
}
