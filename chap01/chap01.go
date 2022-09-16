package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	var inha string
	var testBool bool
	fmt.Println(testBool)

	// var l, w float64
	// l = 10.5
	// w = 5.7
	// fmt.Printf("%.1f * %.1f = %.1f\n", l, w, l*w)

	// var l, w float64
	// l, w = 10.5, 5.71
	// fmt.Printf("%.1f * %.1f = %.2f\n", l, w, l*w)

	// var l, w float64 = 10.5, 5.71
	// fmt.Printf("%.1f * %.1f = %.2f\n", l, w, l*w)

	// var l, w = 10.5, 5.71
	// fmt.Println(reflect.TypeOf(l))
	// fmt.Printf("%.1f * %.1f = %.2f\n", l, w, l*w)

	l, w := 10.5, 5.71
	fmt.Println(reflect.TypeOf(l))
	fmt.Printf("%.1f * %.1f = %.2f\n", l, w, l*w)

	inha = "컴퓨터정보"
	fmt.Printf("인하공업전문대학 %s\n", inha)

	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("헤드퍼스트 고"))
	//fmt.Println("test")
}
