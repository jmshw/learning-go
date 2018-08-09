package main

import "fmt"

type base struct {
	n int
	v int
}

func (b *base) add() (rtn int) {
	b.n += 3
	return
}

type sub struct {
	base
	s string
}

func (s *sub) newFunc() {
	s.add()
	s.n = 5
}

func main() {
	b := base{1, 2}
	b.add()
	//b.newFunc() //Error

	s := sub{base{1, 3}, "123"}
	s.add()

	s.newFunc()
	fmt.Println("s.n:", s.n)
	fmt.Println("s.s", s.s)
}
