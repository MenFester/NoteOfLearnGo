package main

import (
	"fmt"
	"reflect"
	"unsafe"
)


func main() {
	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}

	var a = [...]int{1, 2, 3}
	var b = &a
	fmt.Println(b[0], b[1])    // 通过数组指针访问数组元素的方式和通过数组类似
	for i, v := range b {
		fmt.Println(i, "=>", v)
	}

	s := "hello, world"
	fmt.Println(s[7:])

	s1 := "hello, world"[:5]

	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len)

	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}

	const s3 = "世界abc"
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%d %x\n", i, s3[i])
	}

	fmt.Printf("%#v\n", []rune("世界"))
	fmt.Printf("%#v\n", string([]rune{'世', '界'}))    // 注意这里是单引号
	fmt.Println(len("世界"))    // 6
}