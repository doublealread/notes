package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	id   int
	addr string
	name string
}

func main() {
	var s Student

	s.name = "make"
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"

	s.Person.name = "yoyo"
	fmt.Printf("s = %+v\n", s)
}
