package main

import "fmt"

func main() {

	var si Stack[int]
	for i := 0; i < 10; i++ {
		si.push(i)
	}
	fmt.Println(si)
	si.dumpStack()
	fmt.Println(si)

	var ss Stack[string]
	for i := 0; i < 10; i++ {
		ss.push(fmt.Sprintf("Random dumb string %v", i+1))
	}
	fmt.Println(ss)
	ss.dumpStack()
	fmt.Println(ss)

}

// type Stack[T any] struct {
// 	lifo []T
// }

type Stack[T any] []T
type StackImpl[T any] *Stack[T]

func (s *Stack[T]) push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) pop() *T {
	if len(*s) == 0 {
		return nil
	}

	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return &last
}

func (s *Stack[T]) dumpStack() {
	for {
		if l := s.pop(); l != nil {
			fmt.Println("last:", *l)
		} else {
			break
		}
	}
}
