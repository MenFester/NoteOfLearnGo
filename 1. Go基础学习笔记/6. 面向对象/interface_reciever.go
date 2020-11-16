package main

import (
	"fmt"
)

type PersonInter interface {
	info()
}

type Person1 struct {
	name string
	age int
}

func (p Person1) info() {
	fmt.Println(p.name , p.age)
}

type Person2 struct {
	name string
	age int
}

func (p *Person2) info() {
	fmt.Println(p.name, (*p).age)
}

func main() {
	// 值接收者方法，可以用值、指针调用
	p1 := Person1{name: "linguanqiang", age: 30}
	p1.info()
	prt_p1 := &p1
	prt_p1.info()

	// 指针接收者方法，可以用值、指针调用
	p2 := Person2{name: "supperMan", age: 100}
	p2.info()
	ptr_p2 := &p2
	ptr_p2.info()

	// 实现了接口的值接收者实现方法的对象，可以用对象值、对象指针赋值给接口变量
	var interP PersonInter
	interP = p1
	interP.info()
	interP = prt_p1
	interP.info()

	// 实现了接口的指针接收者实现方法的对象，仅可以用对象指针赋值给接口变量
	// 接口中存储的具体值（Concrete Value）并不能取到地址
	// interP = p2
	interP = ptr_p2
	interP.info()
}