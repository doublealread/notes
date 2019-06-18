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
	*Person //这里对Perdion进行一个取地址操作
	id      int
	addr    string
}

func main() {

	s1 := Student{&Person{"mike", 'm', 18}, 666, "bj"}
	fmt.Println("s1=", s1) //这里的到的时候地址，输出的persion 的地址
	fmt.Println(s1.name, s1.age, s1.sex, s1.id, s1.addr)

	//先定义变量
	var s2 Student
	s2.Person = new(Person)
	s2.name = "yoyo"
	s2.sex = 'm'
	s2.age = 18
	s2.id = 666
	s2.addr = "bj"
	fmt.Println("s2 = ", s2)
	fmt.Println(s2.name, s2.age, s2.sex, s2.id, s2.addr)

}
