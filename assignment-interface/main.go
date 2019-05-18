package main

import "fmt"

type square struct{}
type triangle struct{}

func (s square) getArea() int {
	return 10
}
func (t triangle) getArea() int {
	return 17
}

type shape interface {
	getArea() int
}

func output(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	s := square{}
	t := triangle{}

	output(s)
	output(t)
}
