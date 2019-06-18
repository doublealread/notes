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
}

func main() {
	s1 := Student{Person{"mike", 'm', 18}, 1, "bj"}
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)
	s1.name = "yoyo"
	s1.sex = 'f'
	s1.age = 22
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)
	s1.Person = Person{"go", 'm', 18}
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)

}
