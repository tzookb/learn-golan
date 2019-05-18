package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	age   int
}

func (p *person) updateFname(n string) {
	p.fname = n
}

func main() {
	tzook := person{
		fname: "Tzook",
		lname: "bar Noy",
		age:   33,
	}
	tzook.updateFname("ZUK")
	fmt.Println(tzook)
}
