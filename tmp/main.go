package main

import (
	"fmt"
)

type A struct {
	Aname string
}

type B struct {
	A
}

type C struct {
	A
}

func show(a A) {
	fmt.Println(a.Aname)
}

func main() {
	b := B{A{"b"}}

	c := C{A{"c"}}

	show(b)
	show(c)

}
