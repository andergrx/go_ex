// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func add_ref(a *int) {
	*a += 1
}

func times(z int) int {
	return z * 10
}

type dude struct {
	age  int
	name string
}

func (d dude) dout() int {
	return d.age * 42
}

func squared() func() int {
	x := 2
	return func() int {
		x *= x
		return x
	}
}

func main() {

	var temp *int
	x := 1
	temp = &x
	*temp += 100
	fmt.Println(*temp, " ", x)

	add_ref(&x)
	add_ref(temp)
	fmt.Println(*temp, " ", x)

	var list [5]int

	for i := 0; i < 5; i++ {
		list[i] = i * 25
	}

	for i, item := range list {
		fmt.Println(i, item)
	}

	g, h, i := 1, 4.56, "poop"
	fmt.Println(g, h, i)

	z := 6
	_ = z

	for g < 5 {
		fmt.Println(g)
		g++
	}

	m := map[string]int{
		"kai":  1,
		"kara": 2,
	}

	fmt.Println(m, m["kara"])

	me := dude{}
	me.age = 43
	me.name = "Gabe"
	fmt.Println(me)

	fmt.Println(me.dout())

	you := dude{11, "kai"}
	fmt.Println(you)

	frick := squared()
	fmt.Println(frick(), frick(), frick())


	fmt.Println("testing....")
}
