package main

import "fmt"

func test() {
	var a, b, c int
	fmt.Println(&a, &b, &c)
}

func main() {
	// a := 4
	// pa := &a // 변수의 주소로 단축연산자를 이용해 변수를 초기화하면 포인터로 선언됨
	//var a int = 4
	//var pa *int = &a

	var d, e int
	fmt.Println(&d, &e)
	test()

	// fmt.Println(*pa)
	// fmt.Println(pa)
	// fmt.Println(&a)
	// fmt.Println(reflect.TypeOf(pa))
}
